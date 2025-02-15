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
		w.Write(rp)
	}
	err = UserRegistorRequestValidator(req)
	if err != nil {
		res := Response{
			ServiceName: "LMS",
			StatusCode:  http.StatusBadRequest,
			Msg:         err.Error(),
		}
		rp := responseSender(res)
		w.Write(rp)
	}
	data, err := RegistrationService(req)
	if err != nil {
		res := Response{
			ServiceName: "LMS",
			StatusCode:  http.StatusBadRequest,
			Msg:         err.Error(),
		}
		rp := responseSender(res)
		w.Write(rp)
	}
	res := Response{
		ServiceName: "LMS",
		StatusCode:  http.StatusOK,
		Msg:         "User Registered Successfully",
		Data:        data,
	}
	rp := responseSender(res)
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
		w.Write(rp)
	}
	err = BookRegistorRequestValidator(req)
	if err != nil {
		res := Response{
			ServiceName: "LMS",
			StatusCode:  http.StatusBadRequest,
			Msg:         err.Error(),
		}
		rp := responseSender(res)
		w.Write(rp)
	}
	data, err := BookRegistrationService(req)
	if err != nil {
		res := Response{
			ServiceName: "LMS",
			StatusCode:  http.StatusBadRequest,
			Msg:         err.Error(),
		}
		rp := responseSender(res)
		w.Write(rp)
	}
	res := Response{
		ServiceName: "LMS",
		StatusCode:  http.StatusOK,
		Msg:         "User Registered Successfully",
		Data:        data,
	}
	rp := responseSender(res)
	w.Write(rp)
}
