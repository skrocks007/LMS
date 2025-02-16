package main

import (
	"context"
	"errors"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegistrationService(req UserRegister) (UserRegister, error) {
	userId := GenerateID()
	req.UserId = strconv.Itoa(userId)
	Coll := GetCollection("UserInfo")
	result, err := Coll.InsertOne(context.Background(), req)
	if err != nil {
		return UserRegister{}, err
	}
	req.MongoId = result.InsertedID
	return req, nil
}

func BookRegistrationService(req Book) (Book, error) {
	bookId := GenerateID()
	req.BookId = strconv.Itoa(bookId)
	Coll := GetCollection("BookInfo")
	result, err := Coll.InsertOne(context.Background(), req)
	if err != nil {
		return Book{}, err
	}
	req.MongoId = result.InsertedID
	return req, nil
}

func BorrowService(req Borrow) error {
	err := fetchUser(req.UserId)
	if err != nil {
		return err
	}
	bookData, err := fetchBook(req.BookId)
	if err != nil {
		return err
	}
	if bookData.Status == "available" && bookData.BookCount > 0 {

	}
}

func fetchUser(userId string) error {
	Coll := GetCollection("UserInfo")
	filter := bson.D{
		primitive.E{
			Key:   "userId",
			Value: userId,
		},
	}
	var userData UserData
	err := Coll.FindOne(context.Background(), filter).Decode(&userData)
	if err != nil {
		return errors.New("Invalid User")
	}
	return nil
}

func fetchBook(bookId string) (BookData, error) {
	Coll := GetCollection("BookInfo")
	filter := bson.D{
		primitive.E{
			Key:   "bookId",
			Value: bookId,
		},
	}
	var bookData BookData
	err := Coll.FindOne(context.Background(), filter).Decode(&bookData)
	if err != nil {
		return BookData{}, errors.New("Invalid Book")
	}
	return bookData, nil
}

// Book Request

// userId
// BookId

// validating user -> ValidateUser() -> check in DB
// validating book -> ValidateBook() -> check in DB

//BookIssueRecord
