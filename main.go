package main

import (
	"flag"

	"log"

	tgClient "github.com/Skulllalka/bot_on_go/clients/telegram"
	event_consumer "github.com/Skulllalka/bot_on_go/consumer/event-consumer"
	"github.com/Skulllalka/bot_on_go/events/telegram"
	"github.com/Skulllalka/bot_on_go/storage/files"
)

const (
	tgBotHOST   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(tgClient.New(tgBotHOST, mustToken()), files.New(storagePath))

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String("tg-bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is empty")
	}

	return *token

}

//.\bot_on_go.exe -tg-bot-token '8378863236:AAHo9QlsKwGmA0NIvL1WRrZ8tv92IDWz4Gw'
