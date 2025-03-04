package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"server/database"
	"server/model"
	"strconv"
	"time"
)

type ClientResponse struct {
	CurrentBid float64 `json:"current_bid"`
}

type GetRateDollarResponse struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", getCurrentDollarRateHandler)
	log.Println("Server started on port :8080")
	http.ListenAndServe(":8080", mux)
}

func getCurrentDollarRateHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := handleRequest(w)
	if err != nil {
		// error to process request
		return
	}

	err = saveToDatabase(w, resp.CurrentBid)
	if err != nil {
		// error to save bid at database
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*resp)
}

func handleRequest(w http.ResponseWriter) (*ClientResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		log.Printf("Error to create NewRequestWithContext: %v", err)
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error to execute request for API: %v", err)
		http.Error(w, "Error to get data from API", http.StatusInternalServerError)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error to read body response: %v", err)
		http.Error(w, "Error to process response", http.StatusInternalServerError)
		return nil, err
	}

	var apiResp GetRateDollarResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		log.Printf("Error to unmarshal API response: %v", err)
		http.Error(w, "Error to unmarshal API response", http.StatusInternalServerError)
		return nil, err
	}

	bidFloat, _ := strconv.ParseFloat(apiResp.USDBRL.Bid, 64)
	return &ClientResponse{
		CurrentBid: bidFloat,
	}, nil
}

func saveToDatabase(w http.ResponseWriter, currentBid float64) error {
	db, err := database.Connect("app.db")
	if err != nil {
		log.Printf("%v", err)
		http.Error(w, "Error to connect at database", http.StatusInternalServerError)
		return err
	}
	defer db.Close()

	if err := database.Migrate(db); err != nil {
		log.Printf("%v", err)
		http.Error(w, "Error to migrate initial database schema", http.StatusInternalServerError)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	dolarQuote := model.NewDolarQuote(currentBid)
	err = dolarQuote.CreateDolarQuote(ctx, db)
	if err != nil {
		log.Printf("%v", err)
		http.Error(w, "Error to create dolar quote in database", http.StatusInternalServerError)
		return err
	}

	return nil
}
