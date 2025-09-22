package domain

import "errors"

var (
	ErrInvalidZip  = errors.New("invalid zipcode")
	ErrZipNotFound = errors.New("can not find zipcode")
)

type Weather struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

type Location struct {
	City string
	UF   string
}

func NewLocation(city, uf string) *Location {
	return &Location{
		City: city,
		UF:   uf,
	}
}

func NewWeather(city string, tempC float64) *Weather {
	return &Weather{
		City:  city,
		TempC: tempC,
		TempF: tempC*1.8 + 32,
		TempK: tempC + 273,
	}
}
