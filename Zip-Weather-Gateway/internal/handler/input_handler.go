package handler

import (
	"encoding/json"
	"net/http"

	"github.com/montzzzzz/challenges/zip-weather-gateway/internal/dto"
	"github.com/montzzzzz/challenges/zip-weather-gateway/internal/usecase"
)

type InputHandler struct {
	usecase *usecase.ProcessInput
}

func NewInputHandler(u *usecase.ProcessInput) *InputHandler {
	return &InputHandler{usecase: u}
}

func (h *InputHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	var req dto.WeatherInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusUnprocessableEntity)
		return
	}

	resp, err := h.usecase.Execute(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
