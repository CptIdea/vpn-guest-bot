package guestVpn

import (
	"time"

	"vpn-guest-bot/internal/core/entity"
	"vpn-guest-bot/internal/core/interfaces"

	"github.com/pkg/errors"
)

type usecase struct {
	scheduler  interfaces.Scheduler
	vpnManager interfaces.VpnManager
	hostInfo   entity.VpnHostInfo
}

func New(
	scheduler interfaces.Scheduler, vpnManager interfaces.VpnManager, hostInfo entity.VpnHostInfo,
) interfaces.GuestVpnUsecase {
	return &usecase{scheduler: scheduler, vpnManager: vpnManager, hostInfo: hostInfo}
}

func (u *usecase) CreateVpn(name string, expireWith time.Duration) (entity.VpnConfiguration, error) {
	err := u.scheduler.AddTask(
		name, func() error {
			return u.vpnManager.Delete(name)
		}, time.Now().Add(expireWith),
	)
	if err != nil {
		return entity.VpnConfiguration{}, errors.Wrap(err, "create deletion task")
	}

	password, err := u.vpnManager.Create(name)
	if err != nil {
		return entity.VpnConfiguration{}, errors.Wrap(err, "create vpn")
	}

	return entity.VpnConfiguration{
		HostInfo:     u.hostInfo,
		UserName:     name,
		UserPassword: password,
		ExpireAt:     time.Now().Add(expireWith),
	}, nil
}

func (u *usecase) DeleteVpn(name string) error {
	u.scheduler.CancelTask(name)

	err := u.vpnManager.Delete(name)
	if err != nil {
		return errors.Wrap(err, "delete vpn")
	}

	return nil
}
