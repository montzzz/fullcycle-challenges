package strategy

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type MockStrategy struct {
	mock.Mock
}

func (m *MockStrategy) IsBlocked(key string) (bool, error) {
	args := m.Called(key)
	return args.Bool(0), args.Error(1)
}

func (m *MockStrategy) IncrementRequestCount(key string, limit int, window time.Duration) (bool, error) {
	args := m.Called(key, limit, window)
	return args.Bool(0), args.Error(1)
}

func (m *MockStrategy) BlockKey(key string, duration time.Duration) error {
	args := m.Called(key, duration)
	return args.Error(0)
}

func (m *MockStrategy) GetTokenLimit(token string) (int, bool, error) {
	args := m.Called(token)
	return args.Int(0), args.Bool(1), args.Error(2)
}
