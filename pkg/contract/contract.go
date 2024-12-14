package contract

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"telegram-bot-project/pkg/config"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/system"
	"github.com/blocto/solana-go-sdk/types"
)

var c *client.Client = client.NewClient("https://api.devnet.solana.com")
var ctx context.Context = context.Background()

const programID = "2LnQChqZcx8NKpNn3L5SmPmtz4tt6B2Wo31bY25sgRGs"

func FetchBalance(addr string) float64 {

	// Fetch balance
	balance, err := c.GetBalance(ctx, addr)
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}

	// Optionally, convert to SOL
	const lamportsPerSOL = 1_000_000_000
	return float64(balance) / lamportsPerSOL
}

func FaucetToWallet(addr string) {
	faucetAccountPrivateKey := config.FaucetSecret

	// Decode the faucet account private key
	var privateKey []byte
	if err := json.Unmarshal([]byte(faucetAccountPrivateKey), &privateKey); err != nil {
		log.Fatalf("failed to decode private key: %v", err)
	}

	// Create faucet account
	faucetAccount, _ := types.AccountFromBytes(privateKey)
	recipientAccount, _ := types.AccountFromBase58(addr)
	fmt.Printf("Recipient Public Key: %s\n", recipientAccount.PublicKey.ToBase58())

	// Compute Faucet PDA
	faucetSeed := []byte("faucet_pda")
	faucetPDA, _, err := common.FindProgramAddress([][]byte{faucetSeed}, programID)
	if err != nil {
		log.Fatalf("failed to find program address: %v", err)
	}

	// Check if the Faucet PDA exists
	faucetAccountInfo, err := c.GetAccountInfo(ctx, faucetPDA.ToBase58())
	if err != nil || faucetAccountInfo == nil {
		// Faucet PDA doesn't exist, initialize it
		fmt.Println("Faucet PDA not found, creating...")

		tx, err := c.SendTransaction(ctx, types.Transaction{
			Signers: []types.Account{faucetAccount},
			Message: system.CreateAccount(
				faucetAccount.PublicKey, // Signer
				faucetPDA,               // New account
				1000000,                 // Lamports
				0,                       // Space
				programID,               // Program ID
			),
		})
		if err != nil {
			log.Fatalf("failed to initialize faucet PDA: %v", err)
		}
		fmt.Printf("Faucet PDA created: %s\n", tx)
	} else {
		fmt.Printf("Faucet PDA exists: %s\n", faucetPDA.ToBase58())
	}

	// Compute Recipient PDA
	recipientSeed := []byte("recipient_pda")
	recipientPDA, _, err := common.FindProgramAddress(
		[][]byte{recipientSeed, recipientAccount.PublicKey.Bytes()},
		programID,
	)
	if err != nil {
		log.Fatalf("failed to find recipient program address: %v", err)
	}

	// Check if the Recipient PDA exists
	recipientAccountInfo, err := c.GetAccountInfo(ctx, recipientPDA.ToBase58())
	if err != nil || recipientAccountInfo == nil {
		// Recipient PDA doesn't exist, initialize it
		fmt.Println("Recipient PDA not found, creating...")

		// Construct and send transaction
		tx, err := c.SendTransaction(ctx, types.Transaction{
			Signers: []types.Account{faucetAccount},
			Message: system.CreateAccount(
				faucetAccount.PublicKey,
				recipientPDA,
				1000000,
				0,
				programID,
			),
		})
		if err != nil {
			log.Fatalf("failed to initialize recipient PDA: %v", err)
		}
		fmt.Printf("Recipient PDA created: %s\n", tx)
	} else {
		fmt.Printf("Recipient PDA exists: %s\n", recipientPDA.ToBase58())
	}

	// Amount to transfer (in lamports)
	amount := uint64(2_000_000)

	// Create transfer instruction
	transferInstruction := system.Transfer(
		faucetAccount.PublicKey,
		recipientAccount.PublicKey,
		amount,
	)

	// Send the transaction
	tx, err := c.SendTransaction(ctx, types.Transaction{
		Signers: []types.Account{faucetAccount},
		Message: []types.Instruction{transferInstruction},
	})
	if err != nil {
		log.Fatalf("failed to transfer SOL: %v", err)
	}
	fmt.Printf("Token Transfer Successful: %s\n", tx)

}
