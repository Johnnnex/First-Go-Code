package tests

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"telegram-bot-project/pkg/bot"
	"testing"
)

func TestHandleMessage(t *testing.T) {
	// Create a mock message
	msg := &tgbotapi.Message{
		Text: "/start",
		Chat: &tgbotapi.Chat{
			ID: 12345,
		},
	}

	// Create a mock bot instance
	bot := &bot.Bot{}

	// Run the handler function and check for any errors or expected outcomes
	bot.HandleMessage(bot.API, msg)
}
