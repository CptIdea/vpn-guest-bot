package entity

import "time"

type VpnConfiguration struct {
	HostInfo VpnHostInfo

	UserName     string
	UserPassword string

	ExpireAt time.Time
}

type VpnHostInfo struct {
	Host         string `json:"host,omitempty"`
	HostPassword string `json:"hostPassword,omitempty"`

	VpnType string `json:"vpnType,omitempty"`
}
