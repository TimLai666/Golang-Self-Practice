package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"gitee.com/shirdonl/goWebActualCombat/chapter4/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeleteCond struct {
	BeforeCond TimeBeforeCond `bson:"tp.startTime"`
}

type TimeBeforeCond struct {
	BeforeTime int64 `bson:"$lt"`
}

func main() {
	var (
		client     = mongodb.MgoCli()
		collection *mongo.Collection
		err        error
		uResult    *mongo.DeleteResult
		delCond    *DeleteCond
	)
	collection = client.Database("test").Collection("collect_test")

	delCond = &DeleteCond{
		BeforeCond: TimeBeforeCond{
			BeforeTime: time.Now().Unix() - 60,
		},
	}

	if uResult, err = collection.DeleteMany(context.TODO(), delCond); err != nil {
		log.Fatal(err)
	}
	fmt.Println(uResult.DeletedCount)
}
