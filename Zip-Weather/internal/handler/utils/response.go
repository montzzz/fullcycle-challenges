package utils

import (
	"net/http"

	"github.com/montzzzzz/challenges/zip-weather/internal/domain"
)

func MapErrorToStatus(err error) (int, string) {
	switch err {
	case domain.ErrInvalidZip:
		return http.StatusUnprocessableEntity, err.Error()
	case domain.ErrZipNotFound:
		return http.StatusNotFound, err.Error()
	default:
		return http.StatusInternalServerError, "internal error"
	}
}
