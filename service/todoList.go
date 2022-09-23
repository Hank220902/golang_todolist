package service

import (
    "context"
	"todolist/model"
    "todolist/database"
    "fmt"
    "log"
)


func CreateToDoList() string {
    
    var ctx = context.TODO()
    // var sunshareboy model.Sunshareboy
	wanger := model.Sunshareboy{"wanger", 24, "北京"}
	insertOne, err := database.Collection.InsertOne(ctx, wanger)
    
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Document: ", insertOne.InsertedID)
    return "Inserted a Single Document: "

}