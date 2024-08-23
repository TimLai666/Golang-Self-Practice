package main

import (
	"context"
	"fmt"

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
		uResult    *mongo.UpdateResult
	)
	collection = client.Database("test").Collection("collect_test")
	filter := bson.M{"jobName": "job1"}
	update := bson.M{"$set": model.UpdateByJobName{Command: "byModel", Content: "model"}}
	if uResult, err = collection.UpdateMany(context.TODO(), filter, update); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("更新成功", uResult.ModifiedCount)
}
