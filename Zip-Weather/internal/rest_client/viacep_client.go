package restclient

import (
	"context"
	"fmt"

	"github.com/montzzzzz/challenges/zip-weather/internal/domain"
	"github.com/montzzzzz/challenges/zip-weather/internal/dto"
)

type ViaCEPClient interface {
	GetLocation(ctx context.Context, cep string) (*domain.Location, error)
}

type ViaCEPClientImpl struct {
	RestClient RestClient
}

func NewViaCEPClient(restClient RestClient) *ViaCEPClientImpl {
	return &ViaCEPClientImpl{RestClient: restClient}
}

func (c *ViaCEPClientImpl) GetLocation(ctx context.Context, cep string) (*domain.Location, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	data, err := DoRequest[dto.ViaCEPResponse](ctx, c.RestClient, url)
	if err != nil {
		return nil, err
	}

	if data.HasError() {
		return nil, domain.ErrZipNotFound
	}

	return &domain.Location{
		City: data.Localidade,
		UF:   data.UF,
	}, nil
}
