package mock

import (
	"context"

	"github.com/montzzzzz/challenges/zip-weather/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockWeatherClient struct {
	mock.Mock
}

func (m *MockWeatherClient) GetWeather(ctx context.Context, city, uf string) (*domain.Weather, error) {
	args := m.Called(ctx, city, uf)
	return args.Get(0).(*domain.Weather), args.Error(1)
}
