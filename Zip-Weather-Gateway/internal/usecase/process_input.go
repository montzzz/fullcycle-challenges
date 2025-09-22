package usecase

import (
	"context"

	"github.com/montzzzzz/challenges/zip-weather-gateway/internal/dto"
	restclient "github.com/montzzzzz/challenges/zip-weather-gateway/internal/rest_client"
)

type ProcessInput struct {
	zipWeatherClient *restclient.ZipWeatherClient
}

func NewProcessInput(c *restclient.ZipWeatherClient) *ProcessInput {
	return &ProcessInput{zipWeatherClient: c}
}

func (u *ProcessInput) Execute(ctx context.Context, req dto.WeatherInput) (*dto.WeatherOutput, error) {
	return u.zipWeatherClient.GetWeather(ctx, req.CEP)
}
