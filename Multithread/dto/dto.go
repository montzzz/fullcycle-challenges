package dto

const (
	OriginBrasilAPI = "BrasilAPI"
	OriginViaCEP    = "ViaCEP"
)

type ToResult interface {
	ToResult() Result
}

type BrasilAPIResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func (b BrasilAPIResponse) ToResult() Result {
	return Result{
		CEP:           b.Cep,
		Address:       b.Street,
		Neighborhood:  b.Neighborhood,
		City:          b.City,
		State:         b.State,
		OriginRequest: OriginBrasilAPI,
	}
}

type ViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (v ViaCepResponse) ToResult() Result {
	return Result{
		CEP:           v.Cep,
		Address:       v.Logradouro,
		Neighborhood:  v.Bairro,
		City:          v.Localidade,
		State:         v.Uf,
		OriginRequest: OriginViaCEP,
	}
}

type Result struct {
	CEP           string `json:"cep"`
	Address       string `json:"address"`
	Neighborhood  string `json:"neighborhood"`
	City          string `json:"city"`
	State         string `json:"uf"`
	OriginRequest string `json:"origin_request"`
}
