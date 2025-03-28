package documents

import "go.mongodb.org/mongo-driver/bson"

type Document interface {
	*struct{}
}

type UpdateDocument[M Document] interface {
	RemoveUnchangedFields(entity M) UpdateDocument[M]
	ToFilterBson() bson.M
	ToSetBson() bson.M
}
