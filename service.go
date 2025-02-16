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

func BookReturnService(req Borrow) (any, error) {
	_, err := fetchBook(req.BookId)
	if err != nil {
		return interface{}(nil), errors.New("Invalid BookId")
	}

	err = fetchUser(req.UserId)
	if err != nil {
		return interface{}(nil), errors.New("Invalid UserId")
	}

	bookIssueData, err := fetchBookIssueRecord(req)
	if err != nil {
		return interface{}(nil), err
	}
	err = updateBookCount(req.BookId)
	if err != nil {
		return interface{}(nil), err
	}
	return bookIssueData, nil
}

func updateBookCount(bookId string) error {
	coll := GetCollection("BookInfo")
	filter := bson.M{"bookId": bookId}
	update := bson.M{
		"$set": bson.M{
			"status": "available",
		},
		"$inc": bson.M{
			"bookCount": 1,
		},
	}
	result := coll.FindOneAndUpdate(context.Background(), filter, update)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func fetchBookIssueRecord(req Borrow) (BorrowDB, error) {
	coll := GetCollection("BookIssueRecord")
	filter := bson.M{
		"userId":     req.UserId,
		"bookId":     req.BookId,
		"isReturned": false,
	}
	returnDate := time.Now().Truncate(24 * time.Hour)
	update := bson.M{
		"$set": bson.M{
			"isReturned": true,
			"returnDate": returnDate,
		},
	}
	var issueRecord BorrowDB
	err := coll.FindOneAndUpdate(context.Background(), filter, update).Decode(&issueRecord)
	if err != nil {
		return BorrowDB{}, err
	}
	return issueRecord, nil
}
