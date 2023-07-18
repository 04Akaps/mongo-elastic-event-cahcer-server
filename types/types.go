package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type CollectionOneChangeEvent struct {
	FullDocument CollectionOne `bson:"fullDocument"`
}

type CollectionOne struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
	Age  int64              `bson:"age"`
}
