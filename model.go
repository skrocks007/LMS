package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRegister struct {
	MongoId interface{} `json:"mongoId,omitempty" bson:"mongoId,omitempty"`
	UserId  string      `json:"userId,omitempty" bson:"userId,omitempty"`
	Name    string      `json:"name" bson:"name"`
	Age     int         `json:"age" bson:"age"`
	Email   string      `json:"email" bson:"email"`
	Contact string      `json:"contact" bson:"contact"`
	Role    string      `json:"role" bson:"role"`
}

type Response struct {
	ServiceName string      `json:"serviceName"`
	StatusCode  int         `json:"statusCode"`
	Msg         string      `json:"msg"`
	Data        interface{} `json:"data,omitempty"`
}

type Book struct {
	MongoId   interface{} `json:"mongoId,omitempty" bson:"mongoId,omitempty"`
	BookId    string      `json:"bookId,omitempty" bson:"bookId,omitempty"`
	Title     string      `json:"title" bson:"title"`
	Author    string      `json:"author" bson:"author"`
	Genre     string      `json:"genre" bson:"genre"`
	Status    string      `json:"status" bson:"status"`
	BookCount int         `json:"bookCount" bson:"bookCount"`
}

type Borrow struct {
	BookId string `json:"bookId,omitempty" bson:"bookId,omitempty"`
	UserId string `json:"userId,omitempty" bson:"userId,omitempty"`
}

type UserData struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId  string             `json:"userId,omitempty" bson:"userId,omitempty"`
	Name    string             `json:"name" bson:"name"`
	Age     int                `json:"age" bson:"age"`
	Email   string             `json:"email" bson:"email"`
	Contact string             `json:"contact" bson:"contact"`
	Role    string             `json:"role" bson:"role"`
}

type BookData struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	BookId    string             `json:"bookId,omitempty" bson:"bookId,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Author    string             `json:"author" bson:"author"`
	Genre     string             `json:"genre" bson:"genre"`
	Status    string             `json:"status" bson:"status"`
	BookCount int                `json:"bookCount" bson:"bookCount"`
}
type BorrowDB struct {
	Id         interface{} `bson:"_id,omitempty" json:"_id,omitempty"`
	UserId     string      `bson:"userId" json:"userId"`
	BookId     string      `bson:"bookId" json:"bookId"`
	IsReturned bool        `bson:"isReturned" json:"isReturned"`
	IssueDate  time.Time   `bson:"issueDate" json:"issueDate"`
	DueDate    time.Time   `bson:"dueDate" json:"dueDate"`
	ReturnDate *time.Time  `bson:"returnDate,omitempty" json:"returnDate,omitempty"`
}
