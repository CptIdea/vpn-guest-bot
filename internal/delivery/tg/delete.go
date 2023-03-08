package tg

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleDeleteCommand(message *tgbotapi.Message) {
	name := message.CommandArguments()
	err := b.vpnUsecase.DeleteVpn(name)
	if err != nil {
		log.Printf("Error deleting VPN configuration: %v", err)
		msg := tgbotapi.NewMessage(
			message.Chat.ID, "Sorry, we could not delete your VPN configuration. Please try again later.",
		)
		b.telegramBot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Your VPN configuration has been deleted.")
	b.telegramBot.Send(msg)
}
