package repository

import "github.com/brickstudy/blockchain-module/src/config"

type Repository struct {
	config *config.Config
}

func NewRepository(config *config.Config) (*Repository, error) {
	r := &Repository{
		config: config,
	}
	return r, nil
}
