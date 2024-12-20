package service

import (
	"github.com/brickstudy/blockchain-module/src/app/repository"
	"github.com/brickstudy/blockchain-module/src/config"
)

type Service struct {
	config *config.Config

	repository *repository.Repository
}

func NewSerivce(config *config.Config, repository *repository.Repository) *Service {
	s := &Service{
		config: config,
	}
	return s
}
