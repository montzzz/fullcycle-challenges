package dto

type WeatherInput struct {
	CEP string `json:"cep"`
}

type WeatherOutput struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
	TempK float64 `json:"temp_k"`
}
