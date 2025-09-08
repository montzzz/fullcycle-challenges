package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/montzzzzz/challenges/zip-weather/internal/domain"
	"github.com/montzzzzz/challenges/zip-weather/internal/dto"
	"github.com/montzzzzz/challenges/zip-weather/internal/handler"
	"github.com/montzzzzz/challenges/zip-weather/internal/test/mock"
	"github.com/stretchr/testify/suite"
)

type WeatherHandlerSuite struct {
	suite.Suite
	handler *handler.WeatherHandler
	mockUC  *mock.MockGetWeatherByCEP
}

func (s *WeatherHandlerSuite) SetupTest() {
	s.mockUC = new(mock.MockGetWeatherByCEP)
	s.handler = &handler.WeatherHandler{GetWeatherByCEP: s.mockUC}
}

func (s *WeatherHandlerSuite) TestGetWeather_Success() {
	expectedWeather := domain.NewWeather(10.0)

	s.mockUC.On("Execute", "01001000").Return(expectedWeather, nil)

	req := httptest.NewRequest("GET", "/weather?cep=01001000", nil)
	w := httptest.NewRecorder()

	s.handler.GetWeather(w, req)

	s.mockUC.AssertCalled(s.T(), "Execute", "01001000")

	s.Equal(http.StatusOK, w.Result().StatusCode)

	var resp domain.Weather
	json.NewDecoder(w.Body).Decode(&resp)
	s.Equal(expectedWeather.TempC, resp.TempC)
	s.Equal(expectedWeather.TempF, resp.TempF)
	s.Equal(expectedWeather.TempK, resp.TempK)
}

func (s *WeatherHandlerSuite) TestGetWeather_Error() {
	s.mockUC.On("Execute", "00000000").Return((*domain.Weather)(nil), domain.ErrInvalidZip)

	req := httptest.NewRequest("GET", "/weather?cep=00000000", nil)
	w := httptest.NewRecorder()

	s.handler.GetWeather(w, req)

	s.mockUC.AssertCalled(s.T(), "Execute", "00000000")

	var resp dto.ErrorResponse
	json.NewDecoder(w.Body).Decode(&resp)
	s.Equal("invalid zipcode", resp.Message)
	s.Equal(http.StatusUnprocessableEntity, w.Result().StatusCode)
}

// Executa a suite
func TestWeatherHandlerSuite(t *testing.T) {
	suite.Run(t, new(WeatherHandlerSuite))
}
