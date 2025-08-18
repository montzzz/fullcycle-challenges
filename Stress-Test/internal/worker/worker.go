package worker

import (
	"crypto/tls"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/montzzzzz/challenges/stress-test/internal/config"
	"github.com/montzzzzz/challenges/stress-test/internal/dto"
)

func RunTest(cfg config.Config) dto.Report {
	results := make(chan dto.Result, cfg.Requests)
	var wg sync.WaitGroup
	start := time.Now()

	semaphore := make(chan struct{}, cfg.Concurrency)
	for i := 0; i < cfg.Requests; i++ {
		wg.Add(1)
		semaphore <- struct{}{} // limit concurrency

		go func() {
			defer wg.Done()
			defer func() { <-semaphore }()

			startReq := time.Now()

			resp, err := getHttpClient().Get(cfg.URL)
			duration := time.Since(startReq)

			if err != nil {
				log.Printf("error for execute request to url:%s -> error: %v", cfg.URL, err)
				results <- dto.Result{StatusCode: 0, Duration: duration, Error: err}
				return
			}
			defer resp.Body.Close()

			results <- dto.Result{StatusCode: resp.StatusCode, Duration: duration}
		}()
	}

	wg.Wait()
	close(results)

	return aggregate(results, time.Since(start))
}

func getHttpClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}

func aggregate(results chan dto.Result, totalTime time.Duration) dto.Report {
	report := dto.Report{
		TotalTime: totalTime,
		ByStatus:  make(map[int]int),
	}

	for r := range results {
		report.Total++
		if r.StatusCode == http.StatusOK {
			report.Success++
		}
		report.ByStatus[r.StatusCode]++
	}

	return report
}
