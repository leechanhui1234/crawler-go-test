package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	name    string `bson:"name"`
	age   int `bson:"age"`
	description   string `bson:"description"`
}

func insertData(ctx context.Context, collection *mongo.Collection) {
	dt, _ := collection.InsertOne(ctx, bson.D{{Key: "name", Value: "Ash"}, {Key: "age", Value: 10}, {Key: "description", Value: "Pallet Town"}})
	fmt.Println(dt)
}

func getData(ctx context.Context, collection *mongo.Collection) {
	var datas []bson.M
	// 데이터 읽기
    res, err := collection.Find(ctx, bson.D{{}})
	
    // 결과를 변수에 담기
    if err = res.All(ctx, &datas); err != nil {
        fmt.Println(err)
    }
    
    // []byte를 String타입으로 변환
    for _, result := range datas {
		output, err := json.MarshalIndent(result, "", "   ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}
}

// func main () {
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	collection := client.Database("test").Collection("trainers")
// 	fmt.Println(collection)

// 	err = client.Ping(ctx,readpref.Primary()) // Primary DB에 대한 연결 체크

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	insertData(ctx, collection)
// 	dt := getData(ctx, collection)
// 	fmt.Println(dt)
// }

func main () {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("<<input local db uri>>").SetServerAPIOptions(serverAPI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
	  panic(err)
	}
	defer func() {
	  if err = client.Disconnect(ctx); err != nil {
		panic(err)
	  }
	}()
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(ctx, bson.D{{"ping", 1}}).Err(); err != nil {
	  panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	collection := client.Database("test").Collection("trainers")
	insertData(ctx, collection)
	getData(ctx, collection)
}

 