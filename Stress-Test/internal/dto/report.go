package dto

import "time"

type Report struct {
	TotalTime time.Duration
	Total     int
	Success   int
	ByStatus  map[int]int
}
