package models

import "time"

type TodoList struct {
	Matter            string    `json:"matter" bson:"matter"`
	CreateTime        time.Time `json:"createTime" bson:"createTime"`
	EndTime           time.Time `json:"endtTime" bson:"endtTime"`
	FinishedCondition string    `json:"finishedCondition" bson:"finishedCondition"`
	Status            string    `json:"status" bson:"status"`
	Email             string    `json:"email" bson:"email"`
}

type User struct {
	Name     string `json:"name" bson:"name"`
	Password string `json:"-" bson:"-"`
	Email    string `json:"email" bson:"email"`
}
