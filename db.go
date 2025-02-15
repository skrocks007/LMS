package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//mongo connection -> database -> collection -> document
//SQL -> database -> table -> feild

func ConnectMongoDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // Change if your DB is hosted elsewhere

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Ping to check connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("MongoDB connection failed:", err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

func GetCollection(collectionName string) *mongo.Collection {
	// if Client == nil {
	// 	Client = ConnectMongoDB()
	// }
	collection := Client.Database("LMS").Collection(collectionName)

	return collection
}
