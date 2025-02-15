package main

import (
	"context"
	"strconv"
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
