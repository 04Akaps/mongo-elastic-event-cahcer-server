package mongo

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"mongo-event-cacher/types"
	"sync"
	"time"
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

func (c *collectionOne) CatchInsertEvent(elasticClient *elastic.Client) {
	bulk := elasticClient.Bulk()
	ctx := context.TODO()

	ticker := time.NewTicker(SYNC_TIME_TTL)
	s := sync.Mutex{}

	go func() {
		for {
			select {
			case <-ticker.C:
				s.Lock()

				if bulk.NumberOfActions() > 0 {
					_, err := bulk.Do(ctx)

					if err != nil {
						log.Fatalf("bulk.Do(ctx) ERROR:", err)
					}

					bulk = elasticClient.Bulk()
					log.Println("Insert Bulk Event")
				}

				s.Unlock()
			}
		}
	}()

	for c.insert.Next(ctx) {
		s.Lock()
		var changeEvent types.CollectionOneChangeEvent
		if err := c.insert.Decode(&changeEvent); err != nil {
			return
		}

		req := elastic.NewBulkUpdateRequest()
		fmt.Println(changeEvent.FullDocument.ID.Hex())

		req.UseEasyJSON(true)
		req.Id(changeEvent.FullDocument.ID.Hex())
		req.Index(c.collection.Name())
		req.Doc(&changeEvent.FullDocument)
		req.DocAsUpsert(true)

		bulk.Add(req)

		s.Unlock()
	}

	defer c.insert.Close(ctx)
}

func (c *collectionOne) CatchUpdateEvent(elasticClient *elastic.Client) {
	//bulk := elasticClient.Bulk()
}

func (c *collectionOne) CatchDeleteEvent(elasticClient *elastic.Client) {
	//bulk := elasticClient.Bulk()
}
