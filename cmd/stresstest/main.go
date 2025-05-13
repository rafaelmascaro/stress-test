package main

import (
	"fmt"
	"log"

	"github.com/rafaelmascaro/stress-test/internal/adapters/cli"
	"github.com/rafaelmascaro/stress-test/internal/usecase"
)

func main() {
	args := cli.ParseArgs()
	uc, err := usecase.NewStressTestUseCase(args.URL, args.Requests, args.Concurrency)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Executando teste de carga...\n")
	report := uc.Execute()

	fmt.Printf("\nRelatório do teste de carga:\n")
	fmt.Printf("----------------------------------------------------------------\n")
	fmt.Printf("Tempo total gasto na execução: %v\n", report.TotalTime)
	fmt.Printf("Quantidade total de requests realizados: %d\n", report.TotalRequests)
	fmt.Printf("Quantidade de requests com status HTTP 200: %d\n", report.SuccessRequests)
	fmt.Printf("Distribuição de outros códigos de status HTTP:\n")

	for statusCode, count := range report.ErrorRequests {
		fmt.Printf("- %d: %d\n", statusCode, count)
	}
}
