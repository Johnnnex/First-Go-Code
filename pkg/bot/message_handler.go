package bot

import (
	"fmt"
	"log"
	"telegram-bot-project/pkg/contract"
	"telegram-bot-project/pkg/types"
	"telegram-bot-project/pkg/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// A global or scoped map to store MessageIDs for wallet requests
var walletsState = make(map[int64]string) // Map of ChatID -> MessageState

var wallets = make([]types.Wallet, 0) // Initialize an empty slice of Wallet

var messageId int
var chatId int64
var walletName string

// HandleMessage handles text messages and sends a welcome message with buttons.
func HandleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {

	// Check if the message is a reply and references the wallet request
	if message.ReplyToMessage != nil {
		chatID := message.Chat.ID
		// Check if the user is in the process of adding a wallet
		if state, exists := walletsState[chatID]; exists {
			switch state {
			case "waiting_for_wallet_name":
				// Step 2: Save the wallet name
				walletName = message.Text

				// Send message asking for wallet address
				msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Wallet name '%s' saved! Now, please provide your wallet address:", walletName))
				msg.ReplyMarkup = tgbotapi.ForceReply{
					ForceReply: true,
				}
				if _, err := bot.Send(msg); err != nil {
					log.Printf("Failed to send message: %v", err)
					return
				}
				walletsState[chatID] = "waiting_for_wallet_address"

			case "waiting_for_wallet_address":
				// Step 3: Validate and save the wallet address and balance
				walletAddress := message.Text
				if utils.IsValidWallet(walletAddress) && utils.IsRepeatedWallet(wallets, walletAddress) {
					// Step 4: Fetch balance
					balance := contract.FetchBalance(walletAddress)
					log.Printf("Balance for %s: %.4f SOL", walletAddress, balance)
					// Save wallet info to the slice
					wallets = append(wallets, types.Wallet{
						Name:    walletName,
						Address: walletAddress,
						Balance: balance,
					})

					// Send confirmation message
					responseMessage := fmt.Sprintf("Wallet address '%s' received! Your current balance is %.4f SOL. Thank you.", walletAddress, balance)
					msg := tgbotapi.NewMessage(chatID, responseMessage)
					if _, err := bot.Send(msg); err != nil {
						log.Printf("Failed to send confirmation message: %v", err)
					}

					// Reset state after processing
					delete(walletsState, chatID)

					updateStartMessage(bot)
				} else {
					var errorMessage string
					fmt.Println(utils.IsRepeatedWallet(wallets, walletAddress))
					if !utils.IsRepeatedWallet(wallets, walletAddress) {
						errorMessage = "Wallet address has been inputted before, try another!"
					} else {
						errorMessage = "Invalid wallet address. Please try again."
					}

					msg := tgbotapi.NewMessage(chatID, errorMessage)

					msg.ReplyToMessageID = message.MessageID
					if _, err := bot.Send(msg); err != nil {
						log.Printf("Failed to send invalid wallet message: %v", err)
					}

					// Re-ask for wallet address
					msg = tgbotapi.NewMessage(chatID, "Please provide a valid wallet address:")
					msg.ReplyMarkup = tgbotapi.ForceReply{
						ForceReply: true,
					}
					if _, err := bot.Send(msg); err != nil {
						log.Printf("Failed to send message: %v", err)
					}
				}
			}
		}
		return
	}

	switch message.Text {
	case "/start":
		// Create inline buttons
		inlineKeyboard := CreateFaucetButtons()
		replyMarkup := tgbotapi.NewInlineKeyboardMarkup(inlineKeyboard...)

		// Create and send the welcome message with buttons
		// msg := tgbotapi.NewMessage(message.Chat.ID, "Welcome to the bot! Click below to get started:")
		msg := tgbotapi.NewMessage(
			message.Chat.ID,
			"*Welcome to Faucet Bot 101!*\n\n"+
				"Click below to **get started**\n\n"+
				"👉 *Wallets:*\n"+
				"No Accounts yet\n\n"+
				"Join us now and start exploring the world of decentralized finance. Don't miss out!\n\n"+
				"_Need help?_ Contact our support team anytime.",
		)
		msg.ParseMode = "Markdown" // to parse?

		msg.ReplyMarkup = replyMarkup
		// Send the message and capture the sent message
		sentMsg, err := bot.Send(msg)
		if err != nil {
			log.Printf("Failed to send message: %v", err)
			return
		}

		// Save the ChatID and MessageID
		chatId = sentMsg.Chat.ID
		messageId = sentMsg.MessageID

	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Invalid selection, but hey, I love you!")

		if _, err := bot.Send(msg); err != nil {
			log.Printf("Failed to send message: %v", err)
		}
	}
}

func updateStartMessage(bot *tgbotapi.BotAPI) {
	// Prepare the message text with the updated wallet count
	walletCount := len(wallets)
	var walletStatusStr string

	if walletCount > 0 {
		for index, value := range wallets {
			walletStatusStr += fmt.Sprintf(
				"🔒 Account %d:\nAccount Name: %s\nAccount Balance: %.2f\nAddress: %s\n\n",
				index+1,
				value.Name,
				value.Balance,
				value.Address,
			)
		}
	} else {
		walletStatusStr = "No Accounts yet"
	}

	// Create inline buttons
	inlineKeyboard := CreateFaucetButtons()
	replyMarkup := tgbotapi.NewInlineKeyboardMarkup(inlineKeyboard...)

	// Prepare the message text with the current wallet status
	msgText := fmt.Sprintf(
		"*Welcome to Faucet Bot 101!*\n\n"+
			"Click below to **get started**\n\n"+
			"👉 *Wallets:*\n%s"+
			"Join us now and start exploring the world of decentralized finance. Don't miss out!\n\n"+
			"_Need help?_ Contact our support team anytime.",
		walletStatusStr,
	)

	// Edit the existing message with the updated wallet count and retain the inline keyboard
	editMsg := tgbotapi.NewEditMessageText(chatId, messageId, msgText)
	editMsg.ParseMode = "Markdown"
	editMsg.ReplyMarkup = &replyMarkup

	// Send the edited message
	if _, err := bot.Send(editMsg); err != nil {
		log.Printf("Failed to edit start message: %v", err)
	}
}
