package app

import (
	"errors"
	"fmt"

	"github.com/brickstudy/blockchain-module/src/app/global"
	. "github.com/brickstudy/blockchain-module/src/dto"
	"go.mongodb.org/mongo-driver/mongo"
)

func useCase() {
	fmt.Println()

	fmt.Println("[Blockchain command]")
	fmt.Println("Use Case")

	fmt.Println("1. ", CreateWallet)
	fmt.Println("2. ", ConnectWallet, " <pk>")
	fmt.Println("3. ", ChangeWallet, " <pk>")
	fmt.Println("4. ", TransferCoin, "<to> <amount>")
	fmt.Println("5. ", MintCoin, "<to> <amount>")
	fmt.Println()
}

func (a App) command(input []string) error {
	msg := errors.New("Check command")

	if len(input) == 0 {
		return msg
	} else {
		from := global.FROM()

		switch input[0] {
		case "1", CreateWallet:
			confirmChat(CreateWallet)
			if wallet := a.serivce.MakeWallet(); wallet == nil {
				panic("Failed to create wallet")
			} else {
				fmt.Println()
				a.log.Info("Success to create wallet", "privateKey", wallet.PrivateKey, "publicKey", wallet.PublicKey)
				fmt.Println()
			}
		case "2", ConnectWallet:
			confirmChat(ConnectWallet)
			if from != "" {
				a.log.Debug("Alread connected wallet", "from", from)
			} else {
				if wallet, err := a.serivce.GetWallet(input[1]); err != nil {
					if err == mongo.ErrNoDocuments {
						a.log.Debug("Failed to find wallet pk", "pk", input[1])
					} else {
						a.log.Crit("Failed to find wallet pk", "pk", input[1], "err", err)
					}
				} else {
					fmt.Println(wallet)
					global.SetFrom(wallet.PublicKey)
					a.log.Info("Success to Connect Wallet", "from", wallet.PublicKey)
					fmt.Println()
				}
			}

		case "3", ChangeWallet:
			confirmChat(ChangeWallet)
			if from == "" {
				a.log.Debug("Connected wallet first.")
				fmt.Println()
			} else {
				if wallet, err := a.serivce.GetWallet(input[1]); err != nil {
					if err == mongo.ErrNoDocuments {
						a.log.Debug("Failed to find wallet pk", "pk", input[1])
					} else {
						a.log.Crit("Failed to find wallet pk", "pk", input[1], "err", err)
					}
				} else {
					fmt.Println(wallet)
					global.SetFrom(wallet.PublicKey)
					a.log.Info("Success to Change Wallet", "from", wallet.PublicKey)
					fmt.Println()
				}
			}

		case "4", TransferCoin:
			confirmChat(TransferCoin)
		case "5", MintCoin:
			confirmChat(MintCoin)
		default:
			return msg
		}
		fmt.Println()
	}
	return nil
}

func confirmChat(cmd string) {
	fmt.Printf("You select %s.", cmd)
	fmt.Println()
}
