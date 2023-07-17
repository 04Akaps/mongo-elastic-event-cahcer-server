package types

import (
	"errors"
	"fmt"
)

const (
	COLLECTION_ONE = "collectionOne"
	COLLECTION_TWO = "collectionTwo"
)

var collectionStructMap = map[string]interface{}{
	COLLECTION_ONE: &CollectionOne{},
	COLLECTION_TWO: &CollectionTwo{},
}

func VerifyStructMap(name string) error {
	if _, ok := collectionStructMap[name]; !ok {
		message := fmt.Sprintf("Not Valid Struct Map : %s", name)
		return errors.New(message)
	}

	return nil
}
