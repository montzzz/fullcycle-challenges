package limiter

import (
	"net/http"
	"strings"
	"time"

	"github.com/montzzzzz/challenges/rate-limiter/internal/config"
	"github.com/montzzzzz/challenges/rate-limiter/internal/limiter/strategy"
)

type RateLimiter struct {
	cfg     *config.Config
	storage strategy.Strategy
}

func NewRateLimiter(cfg *config.Config, store strategy.Strategy) *RateLimiter {
	return &RateLimiter{
		cfg:     cfg,
		storage: store,
	}
}

func (rl *RateLimiter) Allow(req *http.Request) (bool, error) {
	token := req.Header.Get("API_KEY")
	var key string
	var limit int

	if token != "" {
		limitVal, found, err := rl.storage.GetTokenLimit(token)
		if err != nil {
			return false, err
		}
		if found {
			key = "token:" + token
			limit = limitVal
		}
	}

	if key == "" {
		ip := strings.Split(req.RemoteAddr, ":")[0]
		key = "ip:" + ip
		limit = rl.cfg.RateLimitDefault
	}

	blocked, err := rl.storage.IsBlocked(key)
	if err != nil || blocked {
		return false, err
	}

	allowed, err := rl.storage.IncrementRequestCount(key, limit, time.Second)
	if err != nil {
		return false, err
	}

	if !allowed {
		err = rl.storage.BlockKey(key, rl.cfg.BlockDuration)
		return false, err
	}
	return true, nil
}
