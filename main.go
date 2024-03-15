package main

import (
	"flag"
	"log"
)

func main() {
	// token = flags.Get(token)

	// tgClient = telegram.New(token)

	// fetcher = fetcher.New()

	// processor = processor.Ner()

	// consumer.Start(fetcher, processor)
}

func mustToken() (string, error) {
	token := flag.String("tokent-bot-token", "", "token for access to tg bot")

	flag.Parse()

	if *token == "" {
		log.Fatal()
	}
}
