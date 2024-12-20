package service

import (
	"fmt"

	"github.com/brickstudy/blockchain-module/src/constants"
)

func (s *Service) newWallet() (string, string, error) {
	return "", "", nil
}

func (s *Service) MakeWallet() *constants.Wallet {
	fmt.Println("들어옴")
	var wallet constants.Wallet
	var err error

	if wallet.PrivateKey, wallet.PrivateKey, err = s.newWallet(); err != nil {
		panic(err)
	} else {
		return &wallet
	}
}
