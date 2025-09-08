package usecase

import (
	"github.com/montzzzzz/challenges/zip-weather/internal/domain"
	"github.com/montzzzzz/challenges/zip-weather/internal/test/mock"
	"github.com/stretchr/testify/suite"
)

type GetWeatherByCEPSuite struct {
	suite.Suite
	usecase     GetWeatherByCep
	mockViaCEP  *mock.MockViaCEPClient
	mockWeather *mock.MockWeatherClient
}

func (s *GetWeatherByCEPSuite) SetupTest() {
	s.mockViaCEP = new(mock.MockViaCEPClient)
	s.mockWeather = new(mock.MockWeatherClient)
	s.usecase = NewGetWeatherByCEP(s.mockViaCEP, s.mockWeather)
}

func (s *GetWeatherByCEPSuite) TestExecute_Success() {
	cep := "01001000"
	location := domain.NewLocation("São Paulo", "SP")
	weather := domain.NewWeather(10.50)

	s.mockViaCEP.On("GetLocation", cep).Return(location, nil)
	s.mockWeather.On("GetWeather", "São Paulo", "SP").Return(weather, nil)

	result, err := s.usecase.Execute(cep)

	s.NoError(err)
	s.Equal(weather, result)
	s.mockViaCEP.AssertCalled(s.T(), "GetLocation", cep)
	s.mockWeather.AssertCalled(s.T(), "GetWeather", "São Paulo", "SP")
}

func (s *GetWeatherByCEPSuite) TestExecute_InvalidCEP() {
	_, err := s.usecase.Execute("123")
	s.ErrorIs(err, domain.ErrInvalidZip)
}
