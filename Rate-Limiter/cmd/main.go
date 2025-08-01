package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/montzzzzz/challenges/rate-limiter/internal/config"
	"github.com/montzzzzz/challenges/rate-limiter/internal/router"
)

func main() {
	_ = godotenv.Load()
	cfg := config.LoadConfig()

	mux, err := router.NewRouter(cfg)
	if err != nil {
		log.Fatal("erro ao iniciar router:", err)
	}

	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
