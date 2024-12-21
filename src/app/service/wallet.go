package service

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/brickstudy/blockchain-module/src/dto"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func (s *Service) newWallet() (string, string, error) {
	p256 := elliptic.P256()

	if private, err := ecdsa.GenerateKey(p256, rand.Reader); err != nil {
		return "", "", err
	} else if private == nil {
		return "", "", errors.New("private key is nill")
	} else {
		privateKeyBytes := crypto.FromECDSA(private)
		privateKey := hexutil.Encode(privateKeyBytes)
		fmt.Println(privateKey)

		importedPrivateKey, err := crypto.HexToECDSA(privateKey[2:])
		if err != nil {
			return "", "", err
		}

		PublicKey := importedPrivateKey.Public()
		publicKeyECDSA, ok := PublicKey.(*ecdsa.PublicKey)
		if !ok {
			return "", "", errors.New("Error casting public key type.")
		}

		address := crypto.PubkeyToAddress(*publicKeyECDSA)

		return privateKey, hexutil.Encode(address[:]), nil
	}
}

func (s *Service) MakeWallet() *dto.Wallet {
	var wallet dto.Wallet
	var err error

	if wallet.PrivateKey, wallet.PublicKey, err = s.newWallet(); err != nil {
		panic(err)
	} else if err = s.repository.CreateNewWallet(&wallet); err != nil {
		return nil
	} else {
		return &wallet
	}
}

func (s *Service) GetWallet(pk string) (*dto.Wallet, error) {
	if wallet, err := s.repository.GetWallet(pk); err != nil {
		return nil, err
	} else {
		return wallet, nil
	}
}
