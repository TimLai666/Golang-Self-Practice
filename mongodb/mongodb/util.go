package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgoCli *mongo.Client

func initDb() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	mgoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func MgoCli() *mongo.Client {
	if mgoCli == nil {
		initDb()
	}
	return mgoCli
}
