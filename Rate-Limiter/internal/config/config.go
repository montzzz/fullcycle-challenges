package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	RateLimitDefault int
	BlockDuration    time.Duration
	RedisAddr        string
	RedisPassword    string
	RedisDB          int
	Tokens           map[string]int
}

func LoadConfig() *Config {
	rate, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_DEFAULT"))
	blockSec, _ := strconv.Atoi(os.Getenv("BLOCK_DURATION"))
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	tokens := parseTokens(os.Getenv("TOKENS"))

	return &Config{
		RateLimitDefault: rate,
		BlockDuration:    time.Duration(blockSec) * time.Second,
		RedisAddr:        os.Getenv("REDIS_ADDR"),
		RedisPassword:    os.Getenv("REDIS_PASSWORD"),
		RedisDB:          db,
		Tokens:           tokens,
	}
}

func parseTokens(env string) map[string]int {
	result := make(map[string]int)
	if env == "" {
		return result
	}

	pairs := strings.Split(env, ";")
	for _, pair := range pairs {
		parts := strings.Split(pair, ":")
		if len(parts) != 2 {
			continue
		}
		limit, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		result[parts[0]] = limit
	}
	return result
}
