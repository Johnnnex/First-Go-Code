package bot

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// SendTemporaryMessage sends a message that deletes itself after a delay
func SendTemporaryMessage(bot *tgbotapi.BotAPI, chatID int64, text string, delay time.Duration) {
	// Send the notification message
	msg := tgbotapi.NewMessage(chatID, text)
	sentMsg, err := bot.Send(msg)
	if err != nil {
		log.Printf("Failed to send notification: %v", err)
		return
	}

	// Schedule deletion
	time.AfterFunc(delay, func() {
		deleteMsg := tgbotapi.NewDeleteMessage(chatID, sentMsg.MessageID)
		if _, err := bot.Send(deleteMsg); err != nil {
			log.Printf("Failed to delete notification: %v", err)
		}
	})
}
