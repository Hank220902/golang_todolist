package models

type User struct {
	Name     string `json:"name" bson:"name"`
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
}

type Login struct{
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
}