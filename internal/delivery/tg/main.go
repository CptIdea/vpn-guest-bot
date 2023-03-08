package tg

import (
	"log"

	"vpn-guest-bot/internal/core/interfaces"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	telegramBot *tgbotapi.BotAPI
	vpnUsecase  interfaces.GuestVpnUsecase
	admins      []int
}

func NewBot(token string, vpnUsecase interfaces.GuestVpnUsecase, admins []int) (*Bot, error) {
	telegramBot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &Bot{telegramBot, vpnUsecase, admins}, nil
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.telegramBot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.telegramBot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if len(b.admins) > 0 {
			var isAdmin bool
			for _, admin := range b.admins {
				if update.Message.From.ID == admin {
					isAdmin = true
					break
				}
			}
			if !isAdmin {
				continue
			}
		}

		switch update.Message.Command() {
		case "start":
			b.handleStartCommand(update.Message)
		case "delete":
			b.handleDeleteCommand(update.Message)
		default:
			b.handleCreateCommand(update.Message)
		}
	}

	return nil
}
