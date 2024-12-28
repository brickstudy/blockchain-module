package service

import (
	"time"

	"github.com/brickstudy/blockchain-module/src/dto"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) CreateBlock(txs []*dto.Transaction, prevHash []byte, height int64) *dto.Block {
	var pHash []byte

	if latestBlock, err := s.repository.GetLatestBlock(); err != nil {
		if err == mongo.ErrNoDocuments {
			s.log.Info("Genesis Block will be created")
			newBlock := createBlockInner(txs, pHash, height)

			return newBlock
		} else {
			s.log.Crit("Failed to get Latest block", "err", err)
			panic(err)
		}
	} else {
		pHash = latestBlock.Hash

		newBlock := createBlockInner(txs, pHash, height)
		return newBlock
	}
}

func createBlockInner(txs []*dto.Transaction, prevHash []byte, height int64) *dto.Block {
	return &dto.Block{
		Time:         time.Now().Unix(),
		Hash:         []byte{},
		Transactions: txs,
		PrevHash:     prevHash,
		Nonce:        0,
		Height:       height,
	}
}
