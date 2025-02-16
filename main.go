package main

import (
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

var Client *mongo.Client

func main() {

	Client = ConnectMongoDB()

	mux := http.NewServeMux()

	// API
	mux.Handle("POST /register/user", http.HandlerFunc(UserRegistration))
	mux.Handle("POST /register/book", http.HandlerFunc(BookRegistration))
	mux.Handle("POST /borrow/book", http.HandlerFunc(BookBorrow))
	mux.Handle("POST /return/book", http.HandlerFunc(BookReturn))

	// server listner
	fmt.Println("server running at localhost:2121")
	err := http.ListenAndServe(":2121", mux)
	if err != nil {
		fmt.Println(err)
	}
}
