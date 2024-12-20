package app

import (
	"github.com/brickstudy/blockchain-module/src/app/repository"
	"github.com/brickstudy/blockchain-module/src/app/service"
	"github.com/brickstudy/blockchain-module/src/config"
)

type App struct {
	config *config.Config

	serivce    *service.Service
	repository *repository.Repository
}

func NewApp(config *config.Config) {
	a := &App{
		config: config,
	}

	var err error
	if a.repository, err = repository.NewRepository(config); err != nil {
		panic(err)
	} else {
		a.serivce = service.NewSerivce(config, a.repository)
	}
}
