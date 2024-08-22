package main

import (
	"context"
	"fmt"
	"log"
	"mongodb/model"
	"mongodb/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	var (
		client     = mongodb.MgoCli()
		err        error
		collection *mongo.Collection
		result     *mongo.InsertManyResult
		id         primitive.ObjectID
	)
	collection = client.Database("test").Collection("collect_test")

	result, err = collection.InsertMany(context.TODO(), []interface{}{
		model.LogRecord{
			JobName: "job1",
			Command: "echo hello",
			Err:     "",
			Content: "1",
			Tp: model.ExecTime{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
		model.LogRecord{
			JobName: "job2",
			Command: "echo hello",
			Err:     "",
			Content: "2",
			Tp: model.ExecTime{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if result == nil {
		log.Fatal("result is nil")
	}
	for _, v := range result.InsertedIDs {
		id = v.(primitive.ObjectID)
		fmt.Println("自增ID:", id.Hex())
	}
}
