package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var insertEventCacher = []bson.M{
	bson.M{
		"$match": bson.M{
			"operationType": "insert",
		},
	},
}

var updateEventCacher = []bson.M{
	bson.M{
		"$match": bson.M{
			"operationType": "update",
		},
	},
}

var deleteEventCacher = []bson.M{
	bson.M{
		"$match": bson.M{
			"operationType": "delete",
		},
	},
}

var latestStreamOption = options.ChangeStream().SetFullDocument(options.UpdateLookup)
