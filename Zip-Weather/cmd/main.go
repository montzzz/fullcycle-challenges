package main

import (
	"log"
	"net/http"

	"github.com/montzzzzz/challenges/zip-weather/internal/config"
	"github.com/montzzzzz/challenges/zip-weather/internal/handler"
	restclient "github.com/montzzzzz/challenges/zip-weather/internal/rest_client"
	"github.com/montzzzzz/challenges/zip-weather/internal/router"
	"github.com/montzzzzz/challenges/zip-weather/internal/usecase"
)

func main() {
	cfg := config.Load()

	restClient := initRestClient()
	viaCEPClient := initViaCEPClient(restClient)
	weatherClient := initWeatherClient(restClient, cfg)

	getWeatherUC := initUseCases(viaCEPClient, weatherClient)
	weatherHandler := initHandlers(getWeatherUC)
	r := initRouter(weatherHandler)

	startServer(cfg.Port, r)
}

func initRestClient() *restclient.HttpRestClient {
	return restclient.NewHttpRestClient()
}

func initViaCEPClient(restClient *restclient.HttpRestClient) restclient.ViaCEPClient {
	return restclient.NewViaCEPClient(restClient)
}

func initWeatherClient(restClient *restclient.HttpRestClient, cfg *config.Config) restclient.WeatherClient {
	log.Printf("Chave definida %s", cfg.WeatherAPIKey)
	if cfg.WeatherAPIKey == "" {
		log.Println("WARNING: WEATHER_API_KEY não definido")
	}
	return restclient.NewWeatherClient(restClient, cfg.WeatherAPIKey)
}

func initUseCases(viaCEPClient restclient.ViaCEPClient, weatherClient restclient.WeatherClient) usecase.GetWeatherByCep {
	return usecase.NewGetWeatherByCEP(viaCEPClient, weatherClient)
}

func initHandlers(getWeatherUC usecase.GetWeatherByCep) *handler.WeatherHandler {
	return &handler.WeatherHandler{
		GetWeatherByCEP: getWeatherUC,
	}
}

func initRouter(weatherHandler *handler.WeatherHandler) http.Handler {
	return router.NewRouter(weatherHandler)
}

func startServer(port string, handler http.Handler) {
	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
