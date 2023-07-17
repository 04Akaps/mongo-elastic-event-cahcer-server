package mongo

import (
	"context"
	"mongo-event-cacher/config"
	"mongo-event-cacher/types"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
	db     *mongo.Database

	CollectionOne collectionOne
	CollectionTwo collectionTwo

	cancelFunc context.CancelFunc
}

func NewMongoDB(cfg *config.Config) (*MongoDB, error) {
	m := &MongoDB{
		cancelFunc: cfg.CancelContextFunc,
	}

	var err error

	ctx := context.Background()

	if m.client, err = mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoDB.DataSource)); err != nil {
		return nil, err
	} else if err = m.client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	m.db = m.client.Database(cfg.MongoDB.DB)

	m.CollectionOne = collectionOne{
		cancelFunc: m.cancelFunc,
		collection: m.db.Collection(types.COLLECTION_ONE),
	}
	m.CollectionTwo = collectionTwo{
		cancelFunc: m.cancelFunc,
		collection: m.db.Collection(types.COLLECTION_TWO),
	}

	m.collectionInit()

	return m, nil
}

func (m *MongoDB) collectionInit() {
	m.CollectionOne.init()
	m.CollectionTwo.init()
}
