package dto

import "time"

type Result struct {
	StatusCode int
	Duration   time.Duration
	Error      error
}
