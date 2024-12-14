package bot

import (
	"fmt"
	"log"
	types "telegram-bot-project/pkg/types"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleCallbackQuery(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	// Example logic for handling callback queries
	switch callback.Data {
	case "add_wallet":
		// Step 1: Ask for wallet name
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Please provide your wallet name:")
		msg.ReplyMarkup = tgbotapi.ForceReply{
			ForceReply: true,
		}

		// Send the message and store the MessageID
		if _, err := bot.Send(msg); err != nil {
			log.Printf("Failed to send message: %v", err)
			return
		}

		// Save the step (waiting for wallet name) in the map or any other way
		walletsState[callback.Message.Chat.ID] = "waiting_for_wallet_name" // Track the user's current state

	case "request_faucet":

		var msgText string
		var constructedWallets = make([]types.AccountButton, 0)

		if len(wallets) == 0 {
			msgText = fmt.Sprintf(
				"*Welcome to Faucet Bot 101!*\n\n" +
					"Click below to **get started**\n\n" +
					"👉 *No wallets connected!*\n" +
					"Please connect your wallet to continue and claim your faucet.\n\n" +
					"You can receive a faucet once every 10 minutes.\n" +
					"The maximum faucet amount is **5 SOL** per claim.\n\n" +
					"Join us now and start exploring the world of decentralized finance. Don't miss out!\n\n" +
					"_Need help?_ Contact our support team anytime.",
			)
		} else {
			msgText = fmt.Sprintf(
				"*Welcome to Faucet Bot 101!*\n\n" +
					"Click below to **get started**\n\n" +
					"👉 *Faucet Info:*\n" +
					"You can receive a faucet once every 10 minutes.\n" +
					"The maximum faucet amount is **5 SOL** per claim.\n\n" +
					"👉 *Select wallets to use for the faucet:*\n" +
					"Choose the wallets you want to use to claim your faucet.\n\n" +
					"Join us now and start exploring the world of decentralized finance. Don't miss out!\n\n" +
					"_Need help?_ Contact our support team anytime.",
			)
		}

		for _, value := range wallets {
			constructedWallets = append(constructedWallets, types.AccountButton{Name: value.Name, Data: value.Address})
		}

		inlineKeyboard := CreateAccountsButtons(constructedWallets)
		replyMarkup := tgbotapi.NewInlineKeyboardMarkup(inlineKeyboard...)

		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, msgText)
		if len(wallets) > 0 {
			msg.ReplyMarkup = replyMarkup
		}
		msg.ParseMode = "Markdown"
		// Send the message and store the MessageID
		if _, err := bot.Send(msg); err != nil {
			log.Printf("Failed to send message: %v", err)
			return
		}

	default:
		shouldReply := false
		for _, value := range wallets {
			if value.Address == callback.Data {
				shouldReply = true
				break
			}
		}
		if shouldReply {

			return
		}
		log.Printf("Unhandled callback data: %s", callback.Data)
	}
}
