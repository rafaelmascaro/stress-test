package cli

import (
	"flag"
)

type CLIArgs struct {
	URL         string
	Requests    int
	Concurrency int
}

func ParseArgs() *CLIArgs {
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 0, "Número total de requests")
	concurrency := flag.Int("concurrency", 0, "Número de chamadas simultâneas")
	flag.Parse()

	return &CLIArgs{
		URL:         *url,
		Requests:    *requests,
		Concurrency: *concurrency,
	}
}
