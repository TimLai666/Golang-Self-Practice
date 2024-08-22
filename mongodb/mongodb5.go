package main

import (
	"context"
	"fmt"
	"log"
	"mongodb/model"
	"mongodb/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var (
		client     = mongodb.MgoCli()
		err        error
		collection *mongo.Collection
		cursor     *mongo.Cursor
	)

	collection = client.Database("test").Collection("collect_test")
	cond := model.FindByJobName{JobName: "job1"}
	if cursor, err = collection.Find(
		context.TODO(),
		cond,
		options.Find().SetSkip(0),
		options.Find().SetLimit(2)); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	for cursor.Next(context.TODO()) {
		var lr model.LogRecord
		if err = cursor.Decode(&lr); err != nil {
			log.Fatal(err)
		}
		fmt.Println(lr)
	}

	var results []model.LogRecord
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	for _, v := range results {
		fmt.Println(v)
	}
}
