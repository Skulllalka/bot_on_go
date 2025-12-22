package main

import (
	"context"
	"flag"

	"log"

	tgClient "github.com/Skulllalka/bot_on_go/clients/telegram"
	event_consumer "github.com/Skulllalka/bot_on_go/consumer/event-consumer"
	"github.com/Skulllalka/bot_on_go/events/telegram"

	//"github.com/Skulllalka/bot_on_go/storage/files"
	"github.com/Skulllalka/bot_on_go/storage/sqlite"
)

const (
	tgBotHOST         = "api.telegram.org"
	sqliteStoragePath = "data/sqlite/storage.db"
	batchSize         = 100
)

func main() {
	//s:= files.New(storagePath)
	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatalf("[ERR] can't reach database %v", err)
	}
	err = s.Init(context.TODO())
	if err != nil {
		log.Fatalf("can't initialize database %v", err)
	}
	eventsProcessor := telegram.New(tgClient.New(tgBotHOST, mustToken()), s)

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
