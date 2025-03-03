package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type ServerResponse struct {
	CurrentBid float64 `json:"current_bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Printf("Error to create NewRequestWithContext: %v", err)
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error to execute request for server.go: %v", err)
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("Status code received from server.go is not acceptable (why not 200?): %v", res.StatusCode)
		panic("Status code received from server.go is not acceptable (has an error?)")
	}

	var sr ServerResponse
	if err := json.NewDecoder(res.Body).Decode(&sr); err != nil {
		panic("Error to unmarshall body response from server.go")
	}

	logBidToFile(sr.CurrentBid)
}

func logBidToFile(bid float64) {
	currentTime := time.Now().Format("02/01/2006 15:04:05")

	// Criando ou abrindo o arquivo .txt no modo append
	file, err := os.OpenFile("bids.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error to open bids file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Cotação do dólar (%s): %.4f\n", currentTime, bid))
	if err != nil {
		log.Fatalf("Error to write in bids file: %v", err)
	}

	fmt.Printf("Registered bid: %.2f in %s\n", bid, currentTime)
}
