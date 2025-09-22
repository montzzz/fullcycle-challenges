package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/montzzzzz/challenges/zip-weather/internal/handler"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func NewRouter(weatherHandler *handler.WeatherHandler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)    // log das requests
	r.Use(middleware.Recoverer) // recover de panics

	r.Get("/weather", func(w http.ResponseWriter, r *http.Request) {
		otelhttp.NewHandler(http.HandlerFunc(weatherHandler.GetWeather), "GetWeather").ServeHTTP(w, r)
	})
	return r
}
