package tg

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleCreateCommand(message *tgbotapi.Message) {
	name := message.Text
	vpnConfig, err := b.vpnUsecase.CreateVpn(name, time.Hour*24)
	if err != nil {
		log.Printf("Error creating VPN configuration: %v", err)
		msg := tgbotapi.NewMessage(
			message.Chat.ID, "Sorry, we could not create your VPN configuration. Please try again later.",
		)
		b.telegramBot.Send(msg)
		return
	}

	// Construct the response message
	responseMsg := "Your VPN configuration has been created:\n\n"
	responseMsg += "Host: `" + vpnConfig.HostInfo.Host + "`\n"
	responseMsg += "HostPassword: `" + vpnConfig.HostInfo.HostPassword + "`\n\n"
	responseMsg += "Username: `" + vpnConfig.UserName + "`\n"
	responseMsg += "Password: `" + vpnConfig.UserPassword + "`\n\n"
	responseMsg += "Expires at: " + vpnConfig.ExpireAt.Format(time.RFC822) + "\n"

	msg := tgbotapi.NewMessage(message.Chat.ID, responseMsg)
	msg.ParseMode = "markdown"
	b.telegramBot.Send(msg)
}
