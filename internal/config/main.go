package config

import (
	"encoding/json"
	"os"

	"vpn-guest-bot/internal/core/entity"

	"github.com/pkg/errors"
)

type Config struct {
	HostInfo      entity.VpnHostInfo `json:"hostInfo"`
	BotKey        string             `json:"botKey"`
	BotAdmins     []int              `json:"botAdmins"`
	VpnConfigFile string             `json:"vpnConfigFile"`
}

func ParseConfigFile(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, errors.Wrap(err, "cannot read file")
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, errors.Wrap(err, "cannot parse file")
	}

	return cfg, nil
}
