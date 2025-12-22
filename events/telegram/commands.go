package telegram

import (
	"context"
	"errors"
	"log"
	"net/url"

	"strings"

	e "github.com/Skulllalka/bot_on_go/lib"
	"github.com/Skulllalka/bot_on_go/storage"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(text string, chatId int, username string) error {
	text = strings.TrimSpace(text)
	log.Printf("got new command %s from %s", text, username)

	if isAddCmd(text) {
		return p.savePage(chatId, text, username)
	}

	switch text {
	case RndCmd:
		return p.sendRandom(chatId, username)
	case HelpCmd:
		return p.sendHelp(chatId)
	case StartCmd:
		return p.sendHello(chatId)
	default:
		p.tg.SendMessage(chatId, msgUnknownCmd)
	}
	return nil
}

func (p *Processor) savePage(chatID int, pageURL string, username string) error {
	page := &storage.Page{
		URL:      pageURL,
		UserName: username,
	}

	isExists, err := p.storage.IsExists(context.Background(), page)
	if err != nil {
		return e.Wrap("page is exists", err)
	}
	if isExists {
		return p.tg.SendMessage(chatID, msgAlreadyExists)
	}

	if err := p.storage.Save(context.Background(), page); err != nil {
		return e.Wrap("can't save page with", err)
	}

	if err := p.tg.SendMessage(chatID, msgSaved); err != nil {
		return e.Wrap("can't send message with ", err)
	}

	return nil
}

func (p *Processor) sendHelp(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelp)
}

func (p *Processor) sendHello(chatID int) error {
	return p.tg.SendMessage(chatID, msgHello)
}
func (p *Processor) sendRandom(chatId int, username string) error {
	page, err := p.storage.PickRandom(context.Background(), username)
	if err != nil && !errors.Is(err, storage.ErrNoSavedPages) {
		return e.Wrap("can't pick random url", err)
	}
	if errors.Is(err, storage.ErrNoSavedPages) {
		return p.tg.SendMessage(chatId, msgNoSavedPages)
	}
	if err := p.tg.SendMessage(chatId, page.URL); err != nil {
		return e.Wrap("can't send message", err)
	}

	if err := p.storage.Remove(context.Background(), page); err != nil {
		return e.Wrap("can't remove page after sending", err)
	}
	return nil
}

func isAddCmd(text string) bool {
	return isURL(text)
}

func isURL(text string) bool {
	u, err := url.Parse(text)

	return err == nil && u.Host != ""
}
