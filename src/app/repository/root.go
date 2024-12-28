package repository

import (
	"context"

	"github.com/brickstudy/blockchain-module/src/config"
	"github.com/inconshreveable/log15"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client *mongo.Client
	wallet *mongo.Collection
	tx     *mongo.Collection
	block  *mongo.Collection

	log log15.Logger
}

func NewRepository(config *config.Config) (*Repository, error) {
	r := &Repository{
		log: log15.New("module", "mongoDB/repository"),
	}

	var err error
	ctx := context.Background()

	mConfig := config.Mongo

	if r.client, err = mongo.Connect(ctx, options.Client().ApplyURI(mConfig.Uri)); err != nil {
		r.log.Error("Falied to connect to mongo", "uri", mConfig.Uri)
		return nil, err
	} else if err = r.client.Ping(ctx, nil); err != nil {
		r.log.Error("Falied to ping to mongo", "uri", mConfig.Uri)
		return nil, err
	} else {
		db := r.client.Database(config.Mongo.DB, nil)

		r.wallet = db.Collection("wallet")
		r.tx = db.Collection("tx")
		r.block = db.Collection("block")

		r.log.Info("Succes to connet mongo", "uri", mConfig.Uri, "db", mConfig.DB)
		return r, nil
	}
}
