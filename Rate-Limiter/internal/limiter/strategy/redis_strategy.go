package strategy

import (
	"context"
	"fmt"
	"time"

	"github.com/montzzzzz/challenges/rate-limiter/internal/config"
	"github.com/redis/go-redis/v9"
)

type RedisStrategy struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisStrategy(cfg *config.Config) (*RedisStrategy, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	strategy := &RedisStrategy{
		client: rdb,
		ctx:    ctx,
	}

	// Load tokens from .env
	for token, limit := range cfg.Tokens {
		key := "API_KEY_LIMIT:" + token
		exists, err := rdb.Exists(ctx, key).Result()
		if err != nil {
			return nil, err
		}
		if exists == 0 {
			err := rdb.Set(ctx, key, limit, 0).Err()
			if err != nil {
				return nil, err
			}
		}
	}

	return strategy, nil
}

func (r *RedisStrategy) IsBlocked(key string) (bool, error) {
	exists, err := r.client.Exists(r.ctx, "block:"+key).Result()
	return exists == 1, err
}

func (r *RedisStrategy) IncrementRequestCount(key string, limit int, window time.Duration) (bool, error) {
	fullKey := "rate:" + key

	count, err := r.client.Incr(r.ctx, fullKey).Result()
	if err != nil {
		return false, err
	}

	if count == 1 {
		_, err := r.client.Expire(r.ctx, fullKey, window).Result()
		if err != nil {
			return false, err
		}
	}

	return int(count) <= limit, nil
}

func (r *RedisStrategy) BlockKey(key string, duration time.Duration) error {
	return r.client.Set(r.ctx, "block:"+key, "1", duration).Err()
}

func (r *RedisStrategy) GetTokenLimit(token string) (int, bool, error) {
	val, err := r.client.Get(r.ctx, "API_KEY_LIMIT:"+token).Result()
	if err == redis.Nil {
		return 0, false, nil
	} else if err != nil {
		return 0, false, err
	}
	var limit int
	fmt.Sscanf(val, "%d", &limit)
	return limit, true, nil
}
