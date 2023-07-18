package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"mongo-event-cacher/types"
	"sync"
	"time"

	"github.com/olivere/elastic/v7"
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

		req.UseEasyJSON(true)
		req.Id(changeEvent.FullDocument.ID.String())
		req.Index(c.collection.Name())
		req.Doc(&changeEvent.FullDocument)
		req.DocAsUpsert(true)

		bulk.Add(req)

		s.Unlock()
	}

	defer c.insert.Close(ctx)
}

func (c *collectionOne) CatchUpdateEvent(elasticClient *elastic.Client) {
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
					log.Println("Update Bulk Event")
				}

				s.Unlock()
			}
		}
	}()

	for c.update.Next(ctx) {
		s.Lock()
		var changeEvent types.CollectionOneChangeEvent
		if err := c.update.Decode(&changeEvent); err != nil {
			return
		}

		if exists, err := elasticClient.Exists().Index(c.collection.Name()).Id(changeEvent.FullDocument.ID.String()).Do(ctx); err != nil {
			log.Println("elastic Exists Error : ", err)
		} else if exists {
			// ID에 해당하는 값이 존재하는 경우
			if _, err = elasticClient.Update().
				Index(c.collection.Name()).
				Id(changeEvent.FullDocument.ID.String()).
				Doc(&changeEvent.FullDocument).
				Do(ctx); err != nil {
				log.Println("elastic Update Error : ", err)
			} else {
				message := fmt.Sprintf("Update Success %s, id : %s", c.collection.Name(), changeEvent.FullDocument.ID.String())
				log.Println(message)
			}
		} else {
			// TODO elasticSearch에 존재 하지 않는경우
			// 에러 케이스 처리 해야함
			// -> 이런 경우는 없기는 하지만 트러블 슈팅 가능함
			//req := elastic.NewBulkUpdateRequest()
			//
			//req.UseEasyJSON(true)
			//req.Id(changeEvent.FullDocument.ID.String())
			//req.Index(c.collection.Name())
			//req.Doc(&changeEvent.FullDocument)
			//req.DocAsUpsert(true)
			//
			//bulk.Add(req)
		}

		s.Unlock()
	}

	defer c.update.Close(ctx)
}

func (c *collectionOne) CatchDeleteEvent(elasticClient *elastic.Client) {
	ctx := context.TODO()
	s := sync.Mutex{}

	var deleteArray []*elastic.DeleteService
	ticker := time.NewTicker(SYNC_TIME_TTL)

	go func() {
		for {
			select {
			case <-ticker.C:
				s.Lock()

				if len(deleteArray) != 0 {
					for _, query := range deleteArray {
						if _, err := query.Do(ctx); err != nil {
							log.Println("Delete Failed!!!")
						}
					}

					deleteArray = []*elastic.DeleteService{}
				}

				s.Unlock()
			}
		}
	}()

	for c.delete.Next(ctx) {
		s.Lock()

		var changeEvent struct {
			DocumentKey struct {
				ID primitive.ObjectID `bson:"_id"`
			} `bson:"documentKey"`
		}

		if err := c.delete.Decode(&changeEvent); err != nil {
			return
		}

		id := changeEvent.DocumentKey.ID.String()
		deleteQuery := elasticClient.Delete().
			Index(c.collection.Name()).
			Id(id)

		if exists, err := elasticClient.Exists().Index(c.collection.Name()).Id(id).Do(ctx); err != nil {
			log.Println("ERROR : ", err)
		} else if !exists {
			deleteArray = append(deleteArray, deleteQuery)
			return
		} else {
			if _, err = deleteQuery.Do(ctx); err != nil {
				log.Println(err)
				return
			} else {
				log.Println("Delete Success!!")
			}
		}

		s.Unlock()
	}

	defer c.delete.Close(ctx)
}
