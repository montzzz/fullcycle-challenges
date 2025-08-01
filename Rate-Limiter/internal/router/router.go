package router

import (
	"net/http"

	"github.com/montzzzzz/challenges/rate-limiter/internal/config"
	"github.com/montzzzzz/challenges/rate-limiter/internal/limiter"
	"github.com/montzzzzz/challenges/rate-limiter/internal/limiter/strategy"
	"github.com/montzzzzz/challenges/rate-limiter/internal/middleware"
)

func NewRouter(cfg *config.Config) (http.Handler, error) {
	strategy, err := strategy.NewRedisStrategy(cfg)
	if err != nil {
		return nil, err
	}

	rateLimiter := limiter.NewRateLimiter(cfg, strategy)

	mux := http.NewServeMux()
	mux.Handle("/", middleware.NewRateLimiterMiddleware(rateLimiter)(
		http.HandlerFunc(indexHandler),
	))

	return mux, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Requisição aceita!"))
}
