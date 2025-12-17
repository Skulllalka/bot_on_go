package main

import (
	"flag"
	"fmt"
	//"fmt"
	"log"

	"github.com/Skulllalka/bot_on_go/clients/telegram"
)

const (
	tgBotHOST = "api.telegram.org"
)

func main() {

	tgClient := telegram.New(tgBotHOST, mustToken())

	//fetcher = fetcher.New()

	//processor = processor.New()

	//fmt.Println("hello world")
}

func mustToken() string {
	token := flag.String("token-bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is empty")
	}

	return *token

	fmt.Println("some ol++")
}
