package middleware

import (
	"net/http"

	"github.com/montzzzzz/challenges/rate-limiter/internal/limiter"
)

func NewRateLimiterMiddleware(rl *limiter.RateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ok, err := rl.Allow(r)
			if err != nil {
				http.Error(w, "Internal Error", http.StatusInternalServerError)
				return
			}
			if !ok {
				http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
