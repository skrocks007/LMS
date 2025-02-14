package main

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
