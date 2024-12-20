package app

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/brickstudy/blockchain-module/src/app/repository"
	"github.com/brickstudy/blockchain-module/src/app/service"
	"github.com/brickstudy/blockchain-module/src/config"
	. "github.com/brickstudy/blockchain-module/src/constants"

	"github.com/inconshreveable/log15"
)

type App struct {
	config *config.Config

	serivce    *service.Service
	repository *repository.Repository

	log log15.Logger
}

func NewApp(config *config.Config) {
	a := &App{
		config: config,
		log:    log15.New("module", "app"),
	}

	var err error
	if a.repository, err = repository.NewRepository(config); err != nil {
		panic(err)
	} else {
		a.serivce = service.NewSerivce(config, a.repository)
		a.log.Info("Module started", "time", time.Now().Unix())

		sc := bufio.NewScanner(os.Stdin)

		useCase()

		for {
			sc.Scan()

			input := strings.Split(sc.Text(), " ")
			if err = a.command(input); err != nil {
				a.log.Error("Falied to call cli", "err", err, "input", input)
				fmt.Println()
				panic(err)
			}
		}
	}
}

func (a App) command(input []string) error {
	msg := errors.New("Check use case.")

	if len(input) == 0 {
		return msg
	} else {
		switch input[0] {
		case CreateWallet:
			fmt.Println("CreateWallet!!")
			a.serivce.MakeWallet()
		case TransferCoin:
			fmt.Println("TransferCoin!!")
		case MintCoin:
			fmt.Println("MintCoin!!")
		default:
			return msg
		}
		fmt.Println()
	}
	return nil
}

func useCase() {
	fmt.Println()

	fmt.Println("Blockchain command")
	fmt.Println()
	fmt.Println("Use Case")

	fmt.Println("1. ", CreateWallet)
	fmt.Println("2. ", TransferCoin, "<to> <amount>")
	fmt.Println("3. ", MintCoin, "<to> <amount>")
	fmt.Println()
}
