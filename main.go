package main

import (
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

var Client *mongo.Client

func main() {

	ConnectMongoDB()

	mux := http.NewServeMux()

	// API
	mux.Handle("POST /register", http.HandlerFunc(Registration))

	// server listner
	fmt.Println("server running at localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
	}
}
