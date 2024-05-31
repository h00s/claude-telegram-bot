package services

import "github.com/h00s/claude-telegram-bot/config"

type Services struct {
}

func NewServices(c *config.Config) *Services {
	return &Services{}
}
