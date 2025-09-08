package dto

import "strings"

type ViaCEPResponse struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
	Erro       string `json:"erro,omitempty"`
}

func (v *ViaCEPResponse) HasError() bool {
	return strings.EqualFold(v.Erro, "true")
}
