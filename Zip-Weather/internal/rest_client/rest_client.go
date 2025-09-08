package restclient

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type RestClient interface {
	Get(url string) ([]byte, error)
}

type HttpRestClient struct {
	HTTP *http.Client
}

func NewHttpRestClient() *HttpRestClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &HttpRestClient{
		HTTP: &http.Client{Transport: tr},
	}
}

func (c *HttpRestClient) Get(url string) ([]byte, error) {
	resp, err := c.HTTP.Get(url)
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

func DoRequest[T any](client RestClient, url string) (*T, error) {
	body, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	var data T
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("error unmarshalling: %w", err)
	}

	return &data, nil
}
