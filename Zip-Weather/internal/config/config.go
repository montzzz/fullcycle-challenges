package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	WeatherAPIKey string
	ZipkinURL     string
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		Port:          getEnv("PORT", "8080"),
		WeatherAPIKey: getEnv("WEATHER_API_KEY", ""),
		ZipkinURL:     getEnv("ZIPKIN_URL", "http://localhost:9411"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
