package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/montzzzzz/challenges/zip-weather/internal/handler"
)

func NewRouter(weatherHandler *handler.WeatherHandler) http.Handler {
	r := chi.NewRouter()

	r.Get("/weather", weatherHandler.GetWeather)
	return r
}
