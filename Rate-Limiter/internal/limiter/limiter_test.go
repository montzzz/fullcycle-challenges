package limiter

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/montzzzzz/challenges/rate-limiter/internal/config"
	"github.com/montzzzzz/challenges/rate-limiter/internal/limiter/strategy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RateLimiterSuite struct {
	suite.Suite
	mockStore *strategy.MockStrategy
	rl        *RateLimiter
	cfg       *config.Config
}

func (s *RateLimiterSuite) SetupTest() {
	s.mockStore = new(strategy.MockStrategy)
	s.cfg = &config.Config{
		RateLimitDefault: 5,
		BlockDuration:    time.Minute,
	}
	s.rl = NewRateLimiter(s.cfg, s.mockStore)
}

func (s *RateLimiterSuite) Test_Allow_RequestBelowLimit() {
	req := httptest.NewRequest("GET", "/", nil)
	req.RemoteAddr = "127.0.0.1:1234"

	s.mockStore.On("IsBlocked", "ip:127.0.0.1").Return(false, nil)
	s.mockStore.On("IncrementRequestCount", "ip:127.0.0.1", s.cfg.RateLimitDefault, s.cfg.BlockDuration).Return(true, nil)

	allowed, err := s.rl.Allow(req)
	s.NoError(err)
	s.True(allowed)
	s.mockStore.AssertExpectations(s.T())
}

func (s *RateLimiterSuite) Test_Allow_RequestHitsLimit() {
	req := httptest.NewRequest("GET", "/", nil)
	req.RemoteAddr = "127.0.0.1:1234"

	s.mockStore.On("IsBlocked", "ip:127.0.0.1").Return(false, nil)
	s.mockStore.On("IncrementRequestCount", "ip:127.0.0.1", s.cfg.RateLimitDefault, s.cfg.BlockDuration).Return(false, nil)
	s.mockStore.On("BlockKey", "ip:127.0.0.1", s.cfg.BlockDuration).Return(nil)

	allowed, err := s.rl.Allow(req)
	s.NoError(err)
	s.False(allowed)
	s.mockStore.AssertExpectations(s.T())
}

func (s *RateLimiterSuite) Test_Allow_AlreadyBlocked() {
	req := httptest.NewRequest("GET", "/", nil)
	req.RemoteAddr = "127.0.0.1:1234"

	s.mockStore.On("IsBlocked", "ip:127.0.0.1").Return(true, nil)

	allowed, err := s.rl.Allow(req)
	s.NoError(err)
	s.False(allowed)
	s.mockStore.AssertExpectations(s.T())
}

func (s *RateLimiterSuite) Test_Allow_WithTokenLimit() {
	req := httptest.NewRequest("GET", "/", nil)
	req.RemoteAddr = "127.0.0.1:1234"
	req.Header.Set("API_KEY", "abc123")

	s.mockStore.On("GetTokenLimit", "abc123").Return(10, true, nil)
	s.mockStore.On("IsBlocked", "token:abc123").Return(false, nil)
	s.mockStore.On("IncrementRequestCount", "token:abc123", 10, s.cfg.BlockDuration).Return(true, nil)

	allowed, err := s.rl.Allow(req)
	s.NoError(err)
	s.True(allowed)
	s.mockStore.AssertExpectations(s.T())
}

func (s *RateLimiterSuite) Test_Allow_ErrorOnStorage() {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("API_KEY", "123")
	req.RemoteAddr = "127.0.0.1:1234"

	s.mockStore.On("GetTokenLimit", "123").Return(0, false, assert.AnError)

	allowed, err := s.rl.Allow(req)
	s.Error(err)
	s.False(allowed)
	s.mockStore.AssertExpectations(s.T())
}

func TestRateLimiterSuite(t *testing.T) {
	suite.Run(t, new(RateLimiterSuite))
}
