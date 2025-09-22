package usecase

import (
	"context"
	"regexp"
	"strings"

	"github.com/montzzzzz/challenges/zip-weather/internal/domain"
	restclient "github.com/montzzzzz/challenges/zip-weather/internal/rest_client"
)

type GetWeatherByCep interface {
	Execute(ctx context.Context, cep string) (*domain.Weather, error)
}

type GetWeatherByCEPImpl struct {
	ViaCEPClient  restclient.ViaCEPClient
	WeatherClient restclient.WeatherClient
}

func NewGetWeatherByCEP(viaCEP restclient.ViaCEPClient, weather restclient.WeatherClient) GetWeatherByCep {
	return &GetWeatherByCEPImpl{
		ViaCEPClient:  viaCEP,
		WeatherClient: weather,
	}
}

func (uc *GetWeatherByCEPImpl) Execute(ctx context.Context, cep string) (*domain.Weather, error) {
	cep = strings.ReplaceAll(cep, "-", "")

	matched, _ := regexp.MatchString(`^\d{8}$`, cep)
	if !matched {
		return nil, domain.ErrInvalidZip
	}

	location, err := uc.ViaCEPClient.GetLocation(ctx, cep)
	if err != nil {
		return nil, err
	}

	weather, err := uc.WeatherClient.GetWeather(ctx, location.City, location.UF)
	if err != nil {
		return nil, err
	}

	return weather, nil
}
