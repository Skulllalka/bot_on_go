package telegram

import "github.com/Skulllalka/bot_on_go/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	// storage
}

func New(client *telegram.Client)
