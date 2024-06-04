package services

import "github.com/h00s/claude-telegram-bot/config"

type Services struct {
	Claude *Claude
}

func NewServices(c *config.Config) *Services {
	return &Services{
		Claude: NewClaude(c.Anthropic.APIKey),
	}
}
