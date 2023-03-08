package interfaces

import (
	"time"

	"vpn-guest-bot/internal/core/entity"
)

type GuestVpnUsecase interface {
	CreateVpn(name string, expireWith time.Duration) (entity.VpnConfiguration, error)
	DeleteVpn(name string) error
}
