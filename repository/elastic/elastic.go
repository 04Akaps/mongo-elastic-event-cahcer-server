package elastic

import (
	"context"
	"errors"
	"fmt"
	"mongo-event-cacher/config"

	"github.com/olivere/elastic/v7"
)

type Elastic struct {
	Es *elastic.Client
}

func NewElastic(cfg *config.Config) (*Elastic, error) {
	result := &Elastic{}
	var err error

	result.Es, err = elastic.NewClient(
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

func (els *Elastic) CheckIndexExisted(index string) error {
	ctx := context.TODO()
	indices := []string{index}

	existService := elastic.NewIndicesExistsService(els.Es)
	existService.Index(indices)

	exist, err := existService.Do(ctx)

	if err != nil {
		message := fmt.Sprintf("NewIndicesExistsService.Do() %s", err.Error())
		return errors.New(message)
	} else if !exist {
		fmt.Println("nOh no! The index", index, "doesn't exist.")
		fmt.Println("Create the index, and then run the Go script again")
		if _, err = els.Es.CreateIndex(index).Do(ctx); err != nil {
			return err
		} else {
			return nil
		}
	} else if exist {
		fmt.Println("Index name:", index, " exists!")
		return nil
	} else {
		return nil
	}
}
