package uuidPasswordGenerator

import (
	"vpn-guest-bot/internal/core/interfaces"

	"github.com/google/uuid"
)

type generator struct {
}

func New() interfaces.PasswordGenerator {
	return &generator{}
}

func (g *generator) Generate() string {
	return uuid.New().String()
}
