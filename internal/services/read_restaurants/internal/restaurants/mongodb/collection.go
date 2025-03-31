package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Collections struct {
	Book          *mongo.Collection
	Restaurant    *mongo.Collection
	Category      *mongo.Collection
	FailedMessage *mongo.Collection
}

func ProvideCollections(client *mongo.Client) (*Collections, error) {
	db := client.Database("restaurants")
	return &Collections{
		Book:          db.Collection("book"),
		Restaurant:    db.Collection("restaurant"),
		Category:      db.Collection("category"),
		FailedMessage: db.Collection("failed_message"),
	}, nil
}
