package telegram

import (
	"github.com/Skulllalka/bot_on_go/clients/telegram"
	"github.com/Skulllalka/bot_on_go/storage"
)

type Processor struct {
	tg      *telegram.Client
	offset  int
	storage storage.Storage
}

func New(client *telegram.Client)
