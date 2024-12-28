package service

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"

	"github.com/brickstudy/blockchain-module/src/dto"
)

type PowWork struct {
	Block      *dto.Block `json: "block"`
	Target     *big.Int   `json: "target"`
	Difficulty int64      `json: "difficulty"`
}

func (s *Service) NewPow(b *dto.Block) *PowWork {
	target := new(big.Int).SetInt64(1)

	target.Lsh(target, uint(256-s.difficulty))
	return &PowWork{Block: b, Target: target, Difficulty: s.difficulty}
}

func (p *PowWork) RunMinning() (int64, []byte) {
	var iHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		// fmt.Println("nonce : ", nonce)
		d := p.makeHash(int64(nonce))
		hash = sha256.Sum256(d)

		fmt.Printf("\r%x", hash)
		iHash.SetBytes(hash[:])

		if iHash.Cmp(p.Target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Println()
	return int64(nonce), hash[:]
}

func (p *PowWork) makeHash(nonce int64) []byte {
	return bytes.Join(
		[][]byte{
			p.Block.PrevHash,
			intToHex(p.Difficulty),
			intToHex(int64(nonce)),
		},
		[]byte{},
	)
}

func intToHex(number int64) []byte {
	b := new(bytes.Buffer)

	if err := binary.Write(b, binary.BigEndian, number); err != nil {
		panic(err)
	} else {
		return b.Bytes()
	}
}
