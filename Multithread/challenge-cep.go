package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/montzzzzz/challenges/multithread/dto"
)

const (
	URLBrasilAPI = "https://brasilapi.com.br/api/cep/v1/%s"
	URLViaCEP    = "http://viacep.com.br/ws/%s/json/"
)

func fetch[T dto.ToResult](ctx context.Context, url string, ch chan<- dto.Result) {
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return
	}
	defer resp.Body.Close()

	var e T
	if err = json.NewDecoder(resp.Body).Decode(&e); err == nil {
		ch <- e.ToResult()
	}
}

func validateArgs() (string, error) {
	if len(os.Args) != 2 {
		return "", fmt.Errorf("expected 1 argument for zip code, but %d were supplied", len(os.Args)-1)
	}
	return os.Args[1], nil
}

func main() {
	zipCode, err := validateArgs()
	if err != nil {
		panic(err)
	}

	ch := make(chan dto.Result, 2)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go fetch[dto.BrasilAPIResponse](ctx, fmt.Sprintf(URLBrasilAPI, zipCode), ch)
	go fetch[dto.ViaCepResponse](ctx, fmt.Sprintf(URLViaCEP, zipCode), ch)

	select {
	case result := <-ch:
		fmt.Printf("Answer from %s:\n%+v\n", result.OriginRequest, result)
	case <-ctx.Done():
		fmt.Println("Timeout: None of the APIs responded in time.")
	}
}
