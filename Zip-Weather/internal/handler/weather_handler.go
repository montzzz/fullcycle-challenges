package handler

import (
	"encoding/json"
	"net/http"

	"github.com/montzzzzz/challenges/zip-weather/internal/dto"
	"github.com/montzzzzz/challenges/zip-weather/internal/handler/utils"
	"github.com/montzzzzz/challenges/zip-weather/internal/usecase"
)

type WeatherHandler struct {
	GetWeatherByCEP usecase.GetWeatherByCep
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")

	result, err := h.GetWeatherByCEP.Execute(cep)
	if err != nil {
		status, msg := utils.MapErrorToStatus(err)
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(dto.ErrorResponse{Message: msg})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
