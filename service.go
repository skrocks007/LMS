package main

import (
	"context"
	"strconv"
)

func RegistrationService(req UserRegister) (UserRegister, error) {
	userId := GenerateUserID()
	req.UserId = strconv.Itoa(userId)
	Coll := GetCollection("UserInfo")
	result, err := Coll.InsertOne(context.Background(), req)
	if err != nil {
		return UserRegister{}, err
	}
	req.MongoId = result.InsertedID
	return req, nil
}
