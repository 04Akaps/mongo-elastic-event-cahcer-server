package app

import (
	"fmt"
	"mongo-event-cacher/config"
	"mongo-event-cacher/repository/elastic"
	"mongo-event-cacher/repository/mongo"
)

type Listener struct {
	config  *config.Config
	mongo   *mongo.MongoDB
	elastic *elastic.Elastic
}

func NewListener(cfg *config.Config) {
	listener := &Listener{
		config: cfg,
	}
	var err error

	if listener.mongo, err = mongo.NewMongoDB(cfg); err != nil {
		panic(err)
	}

	if listener.elastic, err = elastic.NewElastic(cfg); err != nil {
		panic(err)
	}

	listener.waitUntilBug()
}

func (l *Listener) waitUntilBug() {
	fmt.Println("Event Listner Server Started")
	for {
		select {
		case <-l.config.CancelContext.Done():
			return
		}
	}
}
