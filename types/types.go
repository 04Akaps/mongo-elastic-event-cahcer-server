package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type CollectionOneChangeEvent struct {
	FullDocument CollectionOne `bson:"fullDocument"`
}

type CollectionTwoChangeEvent struct {
	FullDocument CollectionTwo `bson:"fullDocument"`
}

type CollectionOne struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
	Age  int64              `bson:"age"`
}

type CollectionTwo struct {
	ID      primitive.ObjectID `bson:"_id"`
	Address string             `bson:"address"`
	Price   int64              `bson:"price"`
	Inner   Inner              `bson:"inner"`
	Array   Array              `bson:"array"`
}

type Inner struct {
	Owner string `bson:"owner"`
}

type Array struct {
	Array []string `bson:"array"`
}
