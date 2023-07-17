package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type NftChangeEvent struct {
	FullDocument Nft `bson:"fullDocument"`
}

type Nft struct {
	ID           primitive.ObjectID `bson:"_id"`
	Tid          int64              `bson:"tid"`
	Address      string             `bson:"address"`
	Owner        string             `bson:"owner"`
	Name         string             `bson:"name"`
	Image        string             `bson:"image"`
	BgColor      string             `bson:"bgColor"`
	Tag          []string           `bson:"tag"`
	Attributes   interface{}        `bson:"attributes"`
	OnSale       bool               `bson:"onSale"`
	IsWrapped    bool               `bson:"isWrapped"`
	OnChain      bool               `bson:"onChain"`
	UserOffer    bool               `bson:"userOffer"`
	IsBurnt      bool               `bson:"isBurnt"`
	Description  string             `bson:"description"`
	AnimationUrl string             `bson:"animationUrl"`
	ExternalUrl  string             `bson:"externalUrl"`
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
