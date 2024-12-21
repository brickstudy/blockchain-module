package service

import (
	"github.com/brickstudy/blockchain-module/src/app/repository"
	"github.com/brickstudy/blockchain-module/src/config"
	"github.com/inconshreveable/log15"
)

type Service struct {
	config *config.Config

	repository *repository.Repository

	log log15.Logger
}

func NewSerivce(config *config.Config, repository *repository.Repository) *Service {
	s := &Service{
		config:     config,
		repository: repository,
		log:        log15.New("module", "app"),
	}
	return s
}
