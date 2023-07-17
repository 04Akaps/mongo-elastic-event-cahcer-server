package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var insertEventCacher = []bson.M{
	{
		"$match": bson.M{
			"operationType": "insert",
		},
	},
}

var updateEventCacher = []bson.M{
	{
		"$match": bson.M{
			"operationType": "update",
		},
	},
}

var deleteEventCacher = []bson.M{
	{
		"$match": bson.M{
			"operationType": "delete",
		},
	},
}

var latestStreamOption = options.ChangeStream().SetFullDocument(options.UpdateLookup)

const SYNC_TIME_TTL = 3 * time.Second
