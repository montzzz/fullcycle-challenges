package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/montzzzzz/challenges/zip-weather-gateway/internal/handler"
)

func NewRouter(h *handler.InputHandler) http.Handler {
	r := chi.NewRouter()
	r.Post("/weather", h.GetWeather)
	return r
}
