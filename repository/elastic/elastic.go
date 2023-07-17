package elastic

import (
	"mongo-event-cacher/config"

	"github.com/olivere/elastic"
)

type Elastic struct {
	es *elastic.Client
}

func NewElastic(cfg *config.Config) (*Elastic, error) {
	result := &Elastic{}
	var err error

	result.es, err = elastic.NewClient(
		elastic.SetBasicAuth(
			cfg.Elastic.User,
			cfg.Elastic.Password,
		),
		elastic.SetURL(cfg.Elastic.Uri),
		elastic.SetSniff(false),
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}
