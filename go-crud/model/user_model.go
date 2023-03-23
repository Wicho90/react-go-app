package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name"`
}

func NewUser() *User {
	return &User{}
}

type Users []User
