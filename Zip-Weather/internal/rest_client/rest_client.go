package restclient

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type RestClient interface {
	Get(ctx context.Context, url string) ([]byte, error)
}

type HttpRestClient struct {
	HTTP *http.Client
}

func NewHttpRestClient() *HttpRestClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &HttpRestClient{
		HTTP: &http.Client{Transport: otelhttp.NewTransport(tr)},
	}
}

func (c *HttpRestClient) Get(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Printf("Error to execute request, status: %d, status_text: %s, body: %s",
			resp.StatusCode, resp.Status, string(bodyBytes))
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %w", err)
	}

	return body, nil
}

func DoRequest[T any](ctx context.Context, client RestClient, url string) (*T, error) {
	body, err := client.Get(ctx, url)
	if err != nil {
		return nil, err
	}

	var data T
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("error unmarshalling: %w", err)
	}

	return &data, nil
}
