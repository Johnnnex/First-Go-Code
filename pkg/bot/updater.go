package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	API *tgbotapi.BotAPI
}

// NewBot creates and returns a new Bot instance.
func NewBot(token string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	api.Debug = true // Enable debugging (optional)
	return &Bot{API: api}, nil
}

// Start listens for incoming updates and routes them to the appropriate handler.
func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := b.API.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("Error getting updates: %v", err)
	}

	for update := range updates {
		if update.Message != nil {
			b.handleMessage(update.Message)
		}

		if update.CallbackQuery != nil {
			b.handleCallbackQuery(update.CallbackQuery)
		}
	}
}

// handleMessage processes incoming messages.
func (b *Bot) handleMessage(message *tgbotapi.Message) {

	HandleMessage(b.API, message)

	log.Printf("Received message: %s", message.Text)

}

// handleCallbackQuery processes callback queries (inline button clicks).
func (b *Bot) handleCallbackQuery(callback *tgbotapi.CallbackQuery) {

	HandleCallbackQuery(b.API, callback)

	log.Printf("Received callback query from user %d", callback.From.ID)
}
