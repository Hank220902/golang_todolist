package service

import (
	"fmt"
	"log"
	"todolist/database"
	"todolist/model"

	"github.com/kataras/iris/v12"
)

func CreateToDoList(ctx iris.Context) model.Sunshareboy {

	//var ctx = context.TODO()
	// var sunshareboy model.Sunshareboy
	//tasks := new(model.Sunshareboy) 
	wanger := model.Sunshareboy{"wanger", 24, "北京"}
	insertOne, err := database.Collection.InsertOne(ctx, wanger)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Document: ", insertOne.InsertedID)
	return wanger

}
