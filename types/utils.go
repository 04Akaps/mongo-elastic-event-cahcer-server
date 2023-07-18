package types

import (
	"errors"
	"fmt"
)

const (
	COLLECTION_ONE = "collection-one"
)

var collectionStructMap = map[string]interface{}{
	COLLECTION_ONE: &CollectionOne{},
}

func VerifyStructMap(name string) error {
	if _, ok := collectionStructMap[name]; !ok {
		message := fmt.Sprintf("Not Valid Struct Map : %s", name)
		return errors.New(message)
	}

	return nil
}
