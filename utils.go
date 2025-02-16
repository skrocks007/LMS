package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"math/rand"
)

func responseSender(resp Response) []byte {
	res, _ := json.Marshal(resp)
	return res
}
func UserRegistorRequestValidator(req UserRegister) error {
	err := validateName(req.Name) // validating name based on rules
	if err != nil {
		return err
	}
	_, err = validateDOB(req.DOB)
	if err != nil {
		return err
	}
	err = validateEmail(req.Email)
	if err != nil {
		return err
	}
	err = validateContact(req.Contact)
	if err != nil {
		return err
	}
	err = validateRole(req.Role)
	if err != nil {
		return err
	}
	return nil
}

func validateRole(reqRole string) error {
	if reqRole == "" {
		return errors.New("role cannot be empty")
	}
	role := strings.ToLower(reqRole)
	if role != "admin" && role != "librarian" && role != "student" && role != "nonstudent" {
		return errors.New("role not valid (excepted roles are 'Admin', 'Librarian', 'Student' or 'NonStudent')")
	}
	return nil
}

func validateContact(contact string) error {
	if contact == "" {
		return errors.New("contact cannot empty")
	}
	exp, err := regexp.Compile(`^[0-9]{10}$`)
	if err != nil {
		return err
	}
	matched := exp.MatchString(contact)
	if !matched {
		return errors.New("contact should be of 10 digits only")
	}
	return nil
}

func validateEmail(email string) error {
	if email == "" {
		return errors.New("email cannot be empty")
	}
	matched, err := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("Not a valid email")
	}
	return nil
}

func validateName(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	cleamName := strings.TrimSpace(name)
	matched, err := regexp.MatchString(`^[A-Za-z ]+$`, cleamName)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("name can only contain alphabets and whitespaces")
	}
	cleamName = strings.ReplaceAll(cleamName, " ", "")
	if len(cleamName) < 3 {
		return errors.New("name should be atleast of 3 letter (exluding whitespace)")
	}
	return nil
}
func validateDOB(dob string) (int, error) {
	if dob == "" {
		return 0, errors.New("dob cannot be empty")
	}
	// Parse the input date string (DD-MM-YYYY format)
	parsedDOB, err := time.Parse("02-01-2006", dob)
	if err != nil {
		return 0, fmt.Errorf("invalid date format, expected DD-MM-YYYY")
	}
	// Check if the date is in the future
	if parsedDOB.After(time.Now()) {
		return 0, fmt.Errorf("date of birth cannot be in the future")
	}
	today := time.Now()
	years := today.Year() - parsedDOB.Year()
	months := int(today.Month() - parsedDOB.Month())
	days := today.Day() - parsedDOB.Day()

	// Adjust if the month/day is before DOB
	if days < 0 {
		months--
		days += daysInMonth(parsedDOB.Year(), int(parsedDOB.Month()))
	}
	if months < 0 {
		years--
		months += 12
	}
	// Check if age is at least 5 years
	if years < 5 {
		return 0, fmt.Errorf("age must be at least 5 years")
	}
	return years, nil
}
func daysInMonth(year, month int) int {
	if month == 2 {
		if isLeapYear(year) {
			return 29
		}
		return 28
	}
	if month == 4 || month == 6 || month == 9 || month == 11 {
		return 30
	}
	return 31
}
func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func GenerateID() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(900000) + 100000 // Generates number between 100000 and 999999
}

func BookRegistorRequestValidator(req Book) error {
	if req.Title == "" {
		return errors.New("'title' key is empty")
	}
	if req.Author == "" {
		return errors.New("'author' key is empty")
	}
	if req.Genre == "" {
		return errors.New("'genre' key is empty")
	}
	if req.Status == "" {
		return errors.New("'status' key is empty")
	}
	if req.BookCount == 0 {
		return errors.New("'bookCount' key is empty")
	}
	return nil
}

func BorrowValidator(req Borrow) error {
	if req.BookId == "" {
		return errors.New("'bookId' key is empty")
	}
	if req.UserId == "" {
		return errors.New("'userId' key is empty")
	}
	return nil
}
