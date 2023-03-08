package main

import (
	uuidPasswordGenerator "vpn-guest-bot/internal/adapter/passwordGenerator/uuid"
	memoryScheduler "vpn-guest-bot/internal/adapter/scheduler/memory"
	"vpn-guest-bot/internal/adapter/vpnManager/l2tpFileManager"
	"vpn-guest-bot/internal/config"
	"vpn-guest-bot/internal/core/usecases/guestVpn"
	"vpn-guest-bot/internal/delivery/tg"
)

func main() {
	cfg, err := config.ParseConfigFile()
	if err != nil {
		panic(err)
	}

	passwordGenerator := uuidPasswordGenerator.New()
	scheduler := memoryScheduler.New()
	vpnManager := l2tpFileManager.New(cfg.VpnConfigFile, passwordGenerator)

	usecase := guestVpn.New(scheduler, vpnManager, cfg.HostInfo)

	bot, err := tg.NewBot(cfg.BotKey, usecase, cfg.BotAdmins)
	if err != nil {
		panic(err)
	}

	panic(bot.Start())
}
