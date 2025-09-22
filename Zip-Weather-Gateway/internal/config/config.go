package config

import (
	"log"
	"os"
)

type Config struct {
	ZipWeatherApiUrl string
	ZipkinURL        string
	Port             string
}

func Load() *Config {
	cfg := &Config{
		ZipWeatherApiUrl: getEnv("ZIP_WEATHER_API_URL", ""),
		ZipkinURL:        getEnv("ZIPKIN_URL", ""),
		Port:             getEnv("PORT", "8080"),
	}

	if cfg.ZipWeatherApiUrl == "" {
		log.Fatal("ZIP_WEATHER_API_URL not set")
	}
	if cfg.ZipkinURL == "" {
		log.Fatal("ZIPKIN_URL not set")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
