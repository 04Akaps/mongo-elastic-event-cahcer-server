package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type collectionOne struct {
	cancelFunc context.CancelFunc
	collection *mongo.Collection
	update     *mongo.ChangeStream
	insert     *mongo.ChangeStream
	delete     *mongo.ChangeStream
}

func (c *collectionOne) init() {
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
