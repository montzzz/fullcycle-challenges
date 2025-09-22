package main

import (
	"context"
	"log"
	"net/http"

	"github.com/montzzzzz/challenges/zip-weather-gateway/internal/config"
	"github.com/montzzzzz/challenges/zip-weather-gateway/internal/handler"
	restclient "github.com/montzzzzz/challenges/zip-weather-gateway/internal/rest_client"
	"github.com/montzzzzz/challenges/zip-weather-gateway/internal/router"
	"github.com/montzzzzz/challenges/zip-weather-gateway/internal/tracing"
	"github.com/montzzzzz/challenges/zip-weather-gateway/internal/usecase"
)

func main() {
	ctx := context.Background()
	cfg := config.Load()

	// init tracing
	cleanup, err := tracing.Init("zip-gateway", cfg.ZipkinURL)
	if err != nil {
		log.Fatalf("failed to init tracing: %v", err)
	}
	defer cleanup(ctx)

	// setup dependencies
	restClient := restclient.NewZipWeatherClient(cfg.ZipWeatherApiUrl)
	uc := usecase.NewProcessInput(restClient)
	h := handler.NewInputHandler(uc)

	r := router.NewRouter(h)

	log.Printf("Server running on :%s", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, r)
}
