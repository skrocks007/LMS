package main

import (
	"encoding/json"
	"net/http"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req UserRegister
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		res := Response{
			ServiceName: "LMS",
			StatusCode:  http.StatusBadRequest,
			Msg:         err.Error(),
		}
		rp := responseSender(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(rp)
		return
	}
	err = UserRegistorRequestValidator(req)
	if err != nil {
		res := Response{
			ServiceName: "LMS",
			StatusCode:  http.StatusBadRequest,
			Msg:         err.Error(),
		}
		rp := responseSender(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(rp)
		return
	}
	data, err := RegistrationService(req)
	if err != nil {
		res := Response{
			ServiceName: "LMS",
			StatusCode:  http.StatusInternalServerError,
			Msg:         err.Error(),
		}
		rp := responseSender(res)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(rp)
		return
	}
	res := Response{
		ServiceName: "LMS",
		StatusCode:  http.StatusOK,
		Msg:         "User Registered Successfully",
		Data:        data,
	}
	rp := responseSender(res)
	w.WriteHeader(http.StatusOK)
	w.Write(rp)
}
func BookRegistration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req Book
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		res := Response{
			ServiceName: "LMS",
			StatusCode:  http.StatusBadRequest,
			Msg:         err.Error(),
		}
		rp := responseSender(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(rp)
		return
	}
	err = BookRegistorRequestValidator(req)
	if err != nil {
		res := Response{
			ServiceName: "LMS",
			StatusCode:  http.StatusBadRequest,
			Msg:         err.Error(),
		}
		rp := responseSender(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(rp)
		return
	}
	data, err := BookRegistrationService(req)
	if err != nil {
		res := Response{
			ServiceName: "LMS",
			StatusCode:  http.StatusInternalServerError,
			Msg:         err.Error(),
		}
		rp := responseSender(res)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(rp)
		return
	}
	res := Response{
		ServiceName: "LMS",
		StatusCode:  http.StatusOK,
		Msg:         "User Registered Successfully",
		Data:        data,
	}
	rp := responseSender(res)
	w.WriteHeader(http.StatusOK)
	w.Write(rp)
}
func BookBorrow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req Borrow
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		res := Response{
			ServiceName: "LMS",
			StatusCode:  http.StatusBadRequest,
			Msg:         err.Error(),
		}
		rp := responseSender(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(rp)
		return
	}
	err = BorrowValidator(req)
	if err != nil {
		res := Response{
			ServiceName: "LMS",
			StatusCode:  http.StatusBadRequest,
			Msg:         err.Error(),
		}
		rp := responseSender(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(rp)
		return
	}
	result, err := BorrowService(req)
	if err != nil {
		res := Response{
			ServiceName: "LMS",
			StatusCode:  http.StatusInternalServerError,
			Msg:         err.Error(),
		}
		rp := responseSender(res)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(rp)
		return
	}
	res := Response{
		ServiceName: "LMS",
		StatusCode:  http.StatusOK,
		Msg:         "Book Issued",
		Data:        result,
	}
	rp := responseSender(res)
	w.WriteHeader(http.StatusOK)
	w.Write(rp)
}
func BookReturn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req Borrow
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		res := Response{
			ServiceName: "LMS",
			StatusCode:  http.StatusBadRequest,
			Msg:         err.Error(),
		}
		rp := responseSender(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(rp)
		return
	}
	err = BorrowValidator(req)
	if err != nil {
		res := Response{
			ServiceName: "LMS",
			StatusCode:  http.StatusBadRequest,
			Msg:         err.Error(),
		}
		rp := responseSender(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(rp)
		return
	}
	result, err := BookReturnService(req)
	if err != nil {
		res := Response{
			ServiceName: "LMS",
			StatusCode:  http.StatusInternalServerError,
			Msg:         err.Error(),
		}
		rp := responseSender(res)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(rp)
		return
	}
	res := Response{
		ServiceName: "LMS",
		StatusCode:  http.StatusOK,
		Msg:         "Book Returned",
		Data:        result,
	}
	rp := responseSender(res)
	w.WriteHeader(http.StatusOK)
	w.Write(rp)
}
