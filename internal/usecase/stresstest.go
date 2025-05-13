package usecase

import (
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/rafaelmascaro/stress-test/internal/adapters/client"
	"github.com/rafaelmascaro/stress-test/internal/entity"
)

type StressTestOutputDTO struct {
	TotalTime       time.Duration
	TotalRequests   int
	SuccessRequests int
	ErrorRequests   map[int]int
}

type StressTestUseCase struct {
	URL         string
	Requests    int
	Concurrency int
}

func NewStressTestUseCase(
	url string,
	requests int,
	concurrency int,
) (*StressTestUseCase, error) {
	uc := StressTestUseCase{
		URL:         url,
		Requests:    requests,
		Concurrency: concurrency,
	}

	err := uc.ValidateArgs()
	if err != nil {
		return nil, err
	}

	return &uc, nil
}

func (uc *StressTestUseCase) Execute() StressTestOutputDTO {
	var wg sync.WaitGroup
	chReq := make(chan struct{}, uc.Concurrency)

	report := entity.NewReport()

	for i := 0; i < uc.Requests; i++ {
		chReq <- struct{}{}
		wg.Add(1)

		go func() {
			defer wg.Done()
			statusCode, err := client.MakeRequest(uc.URL)
			if err == nil {
				report.AddRequest(statusCode)
			}

			<-chReq
		}()
	}

	wg.Wait()
	close(chReq)

	report.Finish()

	dto := StressTestOutputDTO{
		TotalTime:       report.TotalTime,
		TotalRequests:   report.TotalRequests,
		SuccessRequests: report.SuccessRequests,
		ErrorRequests:   report.ErrorRequests,
	}

	return dto
}

func (uc *StressTestUseCase) ValidateArgs() error {
	var args []string

	if uc.URL == "" {
		args = append(args, "--url")
	}

	if !(uc.Requests > 0) {
		args = append(args, "--requests")
	}

	if !(uc.Concurrency > 0) {
		args = append(args, "--concurrency")
	}

	if len(args) > 0 {
		return errors.New("parâmetros inválidos: " + strings.Join(args, ", "))
	}

	return nil
}
