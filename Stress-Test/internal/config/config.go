package config

import (
	"flag"
	"log"
)

type Config struct {
	URL         string
	Requests    int
	Concurrency int
}

func Load() Config {
	url := flag.String("url", "", "Target service URL")
	requests := flag.Int("requests", 1, "Total number of requests")
	concurrency := flag.Int("concurrency", 1, "Number of concurrent requests")

	flag.Parse()

	if *url == "" {
		log.Fatal("Parameter --url is required")
	}

	return Config{
		URL:         *url,
		Requests:    *requests,
		Concurrency: *concurrency,
	}
}
