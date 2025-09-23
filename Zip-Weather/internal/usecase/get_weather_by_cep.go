package usecase

import (
	"context"
	"regexp"
	"strings"

	"github.com/montzzzzz/challenges/zip-weather/internal/domain"
	restclient "github.com/montzzzzz/challenges/zip-weather/internal/rest_client"
	"go.opentelemetry.io/otel"
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
	tr := otel.Tracer("zip-weather/usecase")

	cep = strings.ReplaceAll(cep, "-", "")
	matched, _ := regexp.MatchString(`^\d{8}$`, cep)
	if !matched {
		return nil, domain.ErrInvalidZip
	}

	ctx, spanCEP := tr.Start(ctx, "ViaCEP Lookup")
	defer spanCEP.End()

	location, err := uc.ViaCEPClient.GetLocation(ctx, cep)
	if err != nil {
		return nil, err
	}

	ctx, spanWeather := tr.Start(ctx, "WeatherAPI Lookup")
	defer spanWeather.End()

	weather, err := uc.WeatherClient.GetWeather(ctx, location.City, location.UF)
	if err != nil {
		return nil, err
	}

	return weather, nil
}
