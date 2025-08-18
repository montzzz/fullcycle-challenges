package main

import (
	"log"

	"github.com/montzzzzz/challenges/stress-test/internal/config"
	"github.com/montzzzzz/challenges/stress-test/internal/report"
	"github.com/montzzzzz/challenges/stress-test/internal/worker"
)

func main() {
	cfg := config.Load()

	results := worker.RunTest(cfg)

	report.Print(results)
	log.Println("Stress test finished.")

}
