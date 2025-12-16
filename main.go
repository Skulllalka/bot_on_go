package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	token := mustToken()

	// tgClient = telegram.New(token)
	
	//fetcher = fetcher.New()

	//processor = processor.New()

	//fmt.Println("hello world")
}

func mustToken() string{
	token := flag.String("token-bot-token", "", "token for access to telegram bot")

	flag.Parse()
	
	if *token == "" {
		log.Fatal("token is empty")
	}

	return *token
}