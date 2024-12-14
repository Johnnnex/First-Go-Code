package bot

import (
	"telegram-bot-project/pkg/types"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// CreateInlineButtons returns a slice of InlineKeyboardButton for the start command.
func CreateInlineButtons() [][]tgbotapi.InlineKeyboardButton {
	return [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonURL("Visit Our Awesome Community", "https://t.me/your_community_link"),
		},
		{
			tgbotapi.NewInlineKeyboardButtonURL("Follow Us on Twitter Now!", "https://twitter.com/your_twitter_handle"),
		},
		{
			tgbotapi.NewInlineKeyboardButtonData("🚀 Get Started", "get_started"),
			tgbotapi.NewInlineKeyboardButtonData("Input your secret phrase", "input_wallet"),
		},
		{
			tgbotapi.NewInlineKeyboardButtonData("Copy demo", "copy"),
		},
	}
}

func CreateFaucetButtons() [][]tgbotapi.InlineKeyboardButton {
	return [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("🚀 Add Wallet", "add_wallet"),
		},
		{
			tgbotapi.NewInlineKeyboardButtonData("Request Faucet", "request_faucet"),
		},
	}
}

func CreateAccountsButtons(data []types.AccountButton) [][]tgbotapi.InlineKeyboardButton {
	var inlineButtons [][]tgbotapi.InlineKeyboardButton

	for _, value := range data {

		button := tgbotapi.NewInlineKeyboardButtonData(value.Name, value.Data)

		inlineButtons = append(inlineButtons, []tgbotapi.InlineKeyboardButton{button})
	}

	return inlineButtons
}
