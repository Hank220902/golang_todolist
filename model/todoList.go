package model

import "time"

type Sunshareboy struct {
	Name string `json:name`
	Age  int    `json:age`
	City string `json:city`
}

type TodoList struct {
	Matter            string    `json:matter`
	StartTime         time.Time `json:startTime`
	EndTime           time.Time `json:endtTime`
	FinishedCondition string    `json:finishedCondition`
	Status            string    `json:status`
	EveryDaymatter    string    `json:everydayMatter`
}

type User struct {
	Name     string `json:name`
	Password string `json:"-"`
	Email    string `json:email`
}
