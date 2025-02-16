package main

import (
	"context"
	"errors"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegistrationService(req UserRegister) (UserRegister, error) {
	userId := GenerateID()
	req.UserId = strconv.Itoa(userId)
	coll := GetCollection("UserInfo")
	result, err := coll.InsertOne(context.Background(), req)
	if err != nil {
		return UserRegister{}, err
	}
	req.MongoId = result.InsertedID
	return req, nil
}

func BookRegistrationService(req Book) (Book, error) {
	bookId := GenerateID()
	req.BookId = strconv.Itoa(bookId)
	coll := GetCollection("BookInfo")
	result, err := coll.InsertOne(context.Background(), req)
	if err != nil {
		return Book{}, err
	}
	req.MongoId = result.InsertedID
	return req, nil
}

func BorrowService(req Borrow) (any, error) {
	err := fetchUser(req.UserId)
	if err != nil {
		return interface{}(nil), err
	}
	bookData, err := fetchBook(req.BookId)
	if err != nil {
		return interface{}(nil), err
	}

	if bookData.Status != "available" && bookData.BookCount == 0 {
		return interface{}(nil), errors.New("requested book not available")
	}

	err = updateBookInfo(req.BookId)
	if err != nil {
		return interface{}(nil), err
	}

	data, err := issueBook(req)
	if err != nil {
		return interface{}(nil), err
	}

	return data, nil
}

func updateBookInfo(bookId string) error {
	coll := GetCollection("BookInfo")
	filter := bson.M{
		"bookId": bookId,
	}
	update := bson.A{
		bson.M{"$set": bson.M{
			"bookCount": bson.M{"$subtract": bson.A{"$bookCount", 1}}, // bookCount - 1
		}},
		bson.M{"$set": bson.M{
			"status": bson.M{
				"$cond": bson.M{
					"if":   bson.M{"$lte": bson.A{"$bookCount", 0}},
					"then": "unavailable",
					"else": "available",
				},
			},
		}},
	}
	_, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func issueBook(req Borrow) (BorrowDB, error) {
	issueDate := time.Now().Truncate(24 * time.Hour)
	dueData := issueDate.AddDate(0, 0, 14)
	issueData := BorrowDB{
		UserId:     req.UserId,
		BookId:     req.BookId,
		IsReturned: false,
		IssueDate:  issueDate,
		DueDate:    dueData,
	}
	coll := GetCollection("BookIssueRecord")
	result, err := coll.InsertOne(context.Background(), issueData)
	if err != nil {
		return BorrowDB{}, err
	}
	issueData.Id = result.InsertedID
	return issueData, nil

}

func fetchUser(userId string) error {
	coll := GetCollection("UserInfo")
	filter := bson.D{
		primitive.E{
			Key:   "userId",
			Value: userId,
		},
	}
	var userData UserData
	err := coll.FindOne(context.Background(), filter).Decode(&userData)
	if err != nil {
		return errors.New("Invalid User")
	}
	return nil
}

func fetchBook(bookId string) (BookData, error) {
	coll := GetCollection("BookInfo")
	filter := bson.D{
		primitive.E{
			Key:   "bookId",
			Value: bookId,
		},
	}
	var bookData BookData
	err := coll.FindOne(context.Background(), filter).Decode(&bookData)
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

// already have user registered
// already have book registered

// book borrow request ->
// 	verify if user is valid ->
// 		verify if book is valid ->
// 			check if book is available ->
// 				Issue book and update booInfo (reduce book count by 1)
// 				if book count becomes 0 then mark it not available
