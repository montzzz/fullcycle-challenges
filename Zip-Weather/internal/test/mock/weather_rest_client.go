package mock

import (
	"github.com/montzzzzz/challenges/zip-weather/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockWeatherClient struct {
	mock.Mock
}

func (m *MockWeatherClient) GetWeather(city, uf string) (*domain.Weather, error) {
	args := m.Called(city, uf)
	return args.Get(0).(*domain.Weather), args.Error(1)
}
