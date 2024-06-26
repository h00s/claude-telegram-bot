package main

import (
	"log"

	"github.com/h00s/claude-telegram-bot/bot"
	"github.com/h00s/claude-telegram-bot/config"
	"github.com/h00s/claude-telegram-bot/services"
)

func main() {
	config := config.NewConfig()
	services := services.NewServices(config)
	bot, err := bot.NewBot(&config.Telegram, services)
	if err != nil {
		log.Panicln(err)
	}

	bot.Start()
}
