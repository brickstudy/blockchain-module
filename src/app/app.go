package app

import (
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
	}
}
