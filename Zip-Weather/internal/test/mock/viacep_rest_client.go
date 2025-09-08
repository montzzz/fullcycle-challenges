package mock

import (
	"github.com/montzzzzz/challenges/zip-weather/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockViaCEPClient struct {
	mock.Mock
}

func (m *MockViaCEPClient) GetLocation(cep string) (*domain.Location, error) {
	args := m.Called(cep)
	return args.Get(0).(*domain.Location), args.Error(1)
}
