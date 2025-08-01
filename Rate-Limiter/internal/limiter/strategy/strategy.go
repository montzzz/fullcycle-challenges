package strategy

import "time"

type Strategy interface {
	IsBlocked(key string) (bool, error)
	IncrementRequestCount(key string, limit int, window time.Duration) (bool, error)
	BlockKey(key string, duration time.Duration) error
	GetTokenLimit(token string) (int, bool, error)
}
