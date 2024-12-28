package repository

import (
	"context"

	"github.com/brickstudy/blockchain-module/src/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) GetLatestBlock() (*dto.Block, error) {
	ctx := context.Background()

	var block dto.Block

	opt := options.FindOne().SetSort(bson.M{"time": -1})

	if err := r.block.FindOne(ctx, bson.M{}, opt).Decode(&block); err != nil {
		return nil, err
	} else {
		return &block, nil
	}
}
