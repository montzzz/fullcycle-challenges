package mock

import (
	"context"

	"github.com/montzzzzz/challenges/zip-weather/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockGetWeatherByCEP struct {
	mock.Mock
}

func (m *MockGetWeatherByCEP) Execute(ctx context.Context, cep string) (*domain.Weather, error) {
	args := m.Called(ctx, cep)
	return args.Get(0).(*domain.Weather), args.Error(1)
}
