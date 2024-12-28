package service

import (
	"math/big"

	"github.com/brickstudy/blockchain-module/src/dto"
)

type PowWork struct {
	Block      *dto.Block `json: "block"`
	Target     *big.Int   `json: "target`
	Difficulty int64      `json: "difficulty"`
}

func (s *Service) NewPow(b *dto.Block) *PowWork {
	target := new(big.Int).SetInt64(1)

	target.Lsh(target, uint(256-s.difficulty))
	return &PowWork{Block: b, Target: target, Difficulty: s.difficulty}
}
