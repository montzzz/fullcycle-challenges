package restclient

import (
	"context"
	"fmt"

	"net/url"

	"github.com/montzzzzz/challenges/zip-weather/internal/domain"
	"github.com/montzzzzz/challenges/zip-weather/internal/dto"
)

type WeatherClient interface {
	GetWeather(ctx context.Context, city, uf string) (*domain.Weather, error)
}

type WeatherClientImpl struct {
	RestClient RestClient
	APIKey     string
}

func NewWeatherClient(restClient RestClient, apiKey string) *WeatherClientImpl {
	return &WeatherClientImpl{
		RestClient: restClient,
		APIKey:     apiKey,
	}
}

func (c *WeatherClientImpl) GetWeather(ctx context.Context, city, uf string) (*domain.Weather, error) {
	url := buildWeatherURL(c.APIKey, city, uf)

	data, err := DoRequest[dto.WeatherAPIResponse](ctx, c.RestClient, url)
	if err != nil {
		return nil, err
	}

	return domain.NewWeather(city, data.Current.TempC), nil
}

func buildWeatherURL(apiKey, city, uf string) string {
	location := fmt.Sprintf("%s,%s", city, uf)
	encoded := url.QueryEscape(location)
	return fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, encoded)
}
