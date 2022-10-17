package models

import "time"

type HaveIDTodoListForMongo struct {
	ID                string    `json:"_id" bson:"_id"`
	Matter            string    `json:"matter" bson:"matter"`
	EndTime           time.Time `json:"endTime" bson:"endTime"`
	FinishedCondition string    `json:"finishedCondition" bson:"finishedCondition"`
	Status            string    `json:"status" bson:"status"`
	Email             string    `json:"email" bson:"email"`
	Note              string    `json:"note" bson:"note"`
	CreateTime        time.Time `json:"createTime" bson:"createTime"`
	UpeateTime        time.Time `json:"updateTime" bson:"updateTime"`
}
