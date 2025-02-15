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
	mux.Handle("POST /register/book", http.HandlerFunc(BookRegistration))

	// server listner
	fmt.Println("server running at localhost:2121")
	err := http.ListenAndServe(":2121", mux)
	if err != nil {
		fmt.Println(err)
	}
}
