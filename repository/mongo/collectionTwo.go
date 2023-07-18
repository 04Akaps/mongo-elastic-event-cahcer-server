package mongo

import (
	"context"
	"sync"

	"github.com/olivere/elastic/v7"

	"go.mongodb.org/mongo-driver/mongo"
)

type collectionTwo struct {
	cancelFunc context.CancelFunc
	collection *mongo.Collection
	update     *mongo.ChangeStream
	insert     *mongo.ChangeStream
	delete     *mongo.ChangeStream

	sync sync.Mutex
}

func (c *collectionTwo) init() {
	ctx := context.TODO()
	var err error

	if c.update, err = c.collection.Watch(ctx, updateEventCacher, latestStreamOption); err != nil {
		panic(err)
	}

	if c.insert, err = c.collection.Watch(ctx, insertEventCacher, latestStreamOption); err != nil {
		panic(err)
	}

	if c.delete, err = c.collection.Watch(ctx, deleteEventCacher, latestStreamOption); err != nil {
		panic(err)
	}

}

func (c *collectionTwo) CatchInsertEvent(elasticClient *elastic.Client) {
	//bulk := elasticClient.Bulk()
}

func (c *collectionTwo) CatchUpdateEvent(elasticClient *elastic.Client) {
	//bulk := elasticClient.Bulk()
}

func (c *collectionTwo) CatchDeleteEvent(elasticClient *elastic.Client) {
	//bulk := elasticClient.Bulk()
}
