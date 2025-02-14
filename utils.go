package main

import (
	"encoding/json"
	"errors"
	"time"

	"math/rand"
)

func responseSender(resp Response) []byte {
	res, _ := json.Marshal(resp)
	return res
}
func UserRegistorRequestValidator(req UserRegister) error {
	if req.Name == "" {
		return errors.New("'name' key is empty")
	}
	if req.Age == 0 {
		return errors.New("'age' key is empty")
	}
	if req.Email == "" {
		return errors.New("'email' key is empty")
	}
	if req.Contact == "" {
		return errors.New("'contact' key is empty")
	}
	if req.Role == "" {
		return errors.New("'role' key is empty")
	}
	return nil
}

func GenerateUserID() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(900000) + 100000 // Generates number between 100000 and 999999
}
