package restclient

import (
	"context"
	"crypto/tls"
	"fmt"

	"encoding/json"
	"net/http"

	"github.com/montzzzzz/challenges/zip-weather-gateway/internal/dto"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
)

type ZipWeatherClient struct {
	baseURL string
	client  *http.Client
}

func NewZipWeatherClient(baseURL string) *ZipWeatherClient {
	baseTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	otelTransport := otelhttp.NewTransport(baseTransport)

	return &ZipWeatherClient{
		baseURL: baseURL,
		client:  &http.Client{Transport: otelTransport},
	}
}

func (c *ZipWeatherClient) GetWeather(ctx context.Context, cep string) (*dto.WeatherOutput, error) {
	tracer := otel.Tracer("zip-gateway")
	ctx, span := tracer.Start(ctx, "ZipWeatherClient.GetWeather")
	defer span.End()

	url := fmt.Sprintf("%s/weather?cep=%s", c.baseURL, cep)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("service Zip-Weather returned status %d", resp.StatusCode)
	}

	var data dto.WeatherOutput
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		span.RecordError(err)
		return nil, err
	}

	return &data, nil
}
