package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/brickstudy/blockchain-module/src/app/global"
	"github.com/brickstudy/blockchain-module/src/app/repository"
	"github.com/brickstudy/blockchain-module/src/app/service"
	"github.com/brickstudy/blockchain-module/src/config"

	"github.com/inconshreveable/log15"
)

type App struct {
	config *config.Config

	serivce    *service.Service
	repository *repository.Repository

	log log15.Logger
}

func NewApp(config *config.Config, difficulty int64) {
	a := &App{
		config: config,
		log:    log15.New("module", "app"),
	}

	var err error
	if a.repository, err = repository.NewRepository(config); err != nil {
		panic(err)
	} else {
		a.serivce = service.NewSerivce(config, a.repository, difficulty)
		a.log.Info("Module started", "time", time.Now().Unix())

		sc := bufio.NewScanner(os.Stdin)

		for {
			useCase()

			from := global.FROM()
			if from != "" {
				a.log.Info("Current Connected Wallet", "from", from)
				fmt.Println()
			}

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
