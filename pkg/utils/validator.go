package utils

import "telegram-bot-project/pkg/types"

// Example wallet validation function
func IsValidWallet(wallet string) bool {
	// Simple example: Check length or regex (adjust as per wallet type)
	return len(wallet) > 20 && len(wallet) < 50
}

func IsRepeatedWallet(wallets []types.Wallet, address string) bool {
	resolve := true
	for _, value := range wallets {
		if value.Address == address {
			resolve = false
		}
	}

	return resolve
}
