package tg

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (b *Bot) handleStartCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Welcome to the VPN bot!")
	b.telegramBot.Send(msg)
}
