package models

type User struct {
	Name     string `json:"name" bson:"name"`
	Password string `json:"-" bson:"-"`
	Email    string `json:"email" bson:"email"`
}