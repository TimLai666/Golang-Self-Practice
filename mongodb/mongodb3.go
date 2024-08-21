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
		iResult    *mongo.InsertOneResult
		id         primitive.ObjectID
	)
	collection = client.Database("test").Collection("collect_test")

	logRecord := model.LogRecord{
		JobName: "job1",
		Command: "echo hello",
		Err:     "",
		Content: "hello",
		Tp: model.ExecTime{
			StartTime: time.Now().Unix(),
			EndTime:   time.Now().Unix() + 10,
		},
	}
	if iResult, err = collection.InsertOne(context.TODO(), logRecord); err != nil {
		log.Fatal(err)
	}
	id = iResult.InsertedID.(primitive.ObjectID)
	fmt.Println("自增ID:", id.Hex())
}
