package main

import (
	"log"
	"telegram-bot-project/pkg/bot"
	"telegram-bot-project/pkg/config"
)

func main() {
	// Load configuration
	if err := config.Load(); err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Initialize the bot
	b, err := bot.NewBot(config.BotToken)
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	// Start listening for updates
	b.Start()
}
