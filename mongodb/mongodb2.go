package main

import (
	"mongodb/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	var (
		client     = mongodb.MgoCli()
		db         *mongo.Database
		collection *mongo.Collection
	)

	db = client.Database("test")

	collection = db.Collection("collect_test")
	collection = collection
}
