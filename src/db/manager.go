package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/connections/

func InitDatabase() {
	uri := "mongodb://localhost:27017"
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}
	DB = client.Database("grasscutter")
}
