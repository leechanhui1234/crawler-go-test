package main

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	name    string `bson:"name"`
	age   int `bson:"age"`
	description   string `bson:"description"`
}

var collection *mongo.Collection

func init() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://chanhuilee:rnFQYCYxJ56w6PJ3@cluster0.g79oeik.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
	  panic(err)
	}

	collection = client.Database("test").Collection("trainers")
}

func insertData() {
	dt, _ := collection.InsertOne(context.TODO(), bson.D{{Key: "name", Value: "Ash"}, {Key: "age", Value: 10}, {Key: "description", Value: "Pallet Town"}})
	fmt.Println(dt)
}

func getData() {
	var datas []bson.M
	// 데이터 읽기
    res, err := collection.Find(context.TODO(), bson.D{{}})
	
    // 결과를 변수에 담기
    if err = res.All(context.TODO(), &datas); err != nil {
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
	insertData()
	getData()
}

 