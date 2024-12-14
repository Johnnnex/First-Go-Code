package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var BotToken string
var FaucetSecret string

func Load() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	BotToken = os.Getenv("TELEGRAM_BOT_TOKEN")
	BotToken = os.Getenv("FAUCET_SECRET")
	if BotToken == "" {
		return fmt.Errorf("TELEGRAM_BOT_TOKEN is not set")
	}

	return nil
}
