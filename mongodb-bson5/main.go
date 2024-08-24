package main

import (
	"context"
	"fmt"
	"log"

	"gitee.com/shirdonl/goWebActualCombat/chapter4/model"
	"gitee.com/shirdonl/goWebActualCombat/chapter4/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	var (
		client     = mongodb.MgoCli()
		collection *mongo.Collection
		err        error
		cursor     *mongo.Cursor
	)
	collection = client.Database("test").Collection("collect_test")

	groupStage := []model.Group{}
	groupStage = append(groupStage, model.Group{
		Group: bson.D{
			{"_id", "$jobName"},
			{"countJob", model.Sum{Sum: 1}},
		},
	})

	if cursor, err = collection.Aggregate(context.TODO(), groupStage); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	var result []bson.M
	if err = cursor.All(context.TODO(), &result); err != nil {
		log.Fatal(err)
	}
	for _, res := range result {
		fmt.Println(res)
	}
}
