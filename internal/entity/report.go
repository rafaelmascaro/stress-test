package entity

import (
	"sync"
	"time"
)

type Report struct {
	TotalTime       time.Duration
	TotalRequests   int
	SuccessRequests int
	ErrorRequests   map[int]int
	start           time.Time
	mu              sync.Mutex
}

func NewReport() *Report {
	return &Report{
		TotalTime:       0,
		TotalRequests:   0,
		SuccessRequests: 0,
		ErrorRequests:   make(map[int]int),
		start:           time.Now(),
		mu:              sync.Mutex{},
	}
}

func (r *Report) AddRequest(statusCode int) {
	r.mu.Lock()

	r.TotalRequests += 1

	if statusCode == 200 {
		r.SuccessRequests += 1
	} else {
		r.ErrorRequests[statusCode] += 1
	}

	r.mu.Unlock()
}

func (r *Report) Finish() {
	r.TotalTime = time.Since(r.start)
}
