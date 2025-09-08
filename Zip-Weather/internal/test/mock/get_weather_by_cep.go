package mock

import (
	"github.com/montzzzzz/challenges/zip-weather/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockGetWeatherByCEP struct {
	mock.Mock
}

func (m *MockGetWeatherByCEP) Execute(cep string) (*domain.Weather, error) {
	args := m.Called(cep)
	return args.Get(0).(*domain.Weather), args.Error(1)
}
