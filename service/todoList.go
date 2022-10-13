package service

import (
	"fmt"
	"log"
	"time"
	"todolist/connect"
	"todolist/middleware"
	"todolist/models"
	"github.com/davecgh/go-spew/spew"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const(
	success int = 1
	fail int =2
	tokenError int = 3
	emailExists int = 4
)

func CreateToDoList(ctx iris.Context) int {
	email := middleware.MyAuthenticatedHandler(ctx)
	if email == "token not found" {
		return 3
	}

	var TodoList models.TodoList

	if err := ctx.ReadJSON(&TodoList); err != nil {
		panic(err.Error())
	}

	status := ""
	createTime := time.Now()
	fmt.Println(TodoList.EndTime)
	result := models.TodoList{
		Matter:            TodoList.Matter,
		EndTime:           TodoList.EndTime,
		FinishedCondition: TodoList.FinishedCondition,
		Status:            status,
		Email:             email,
		CreateTime:        createTime,
	}
	insertOne, err := connect.TodolistCollection.InsertOne(ctx, result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Document: ", insertOne.InsertedID)
	return 1

}

func GetOneToDoList(ctx iris.Context) ([]*models.HaveIDTodoList, int) {

	var results []*models.HaveIDTodoList
	email := middleware.MyAuthenticatedHandler(ctx)
	if email == "token not found" {
		return results, 3
	}
	paramsStatus := ctx.URLParam("status")
	paramsFinishedCondition := ctx.URLParam("finishedCondition")
	filter := bson.D{{Key: "email", Value: email}}
	if paramsStatus != "" && paramsFinishedCondition != ""{
		filter =bson.D{{Key: "email", Value: email},{Key: "status", Value: paramsStatus},{Key: "finishedCondition", Value: paramsFinishedCondition}}
	}else if paramsStatus != ""{
		filter =bson.D{{Key: "email", Value: email},{Key: "status", Value: paramsStatus}}
	}else if paramsFinishedCondition != ""{
		filter =bson.D{{Key: "email", Value: email},{Key: "finishedCondition", Value: paramsFinishedCondition}}
	}
	

	updateStatus(ctx,email)
	

	cur, err := connect.TodolistCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {

		var result models.HaveIDTodoList
		err := cur.Decode(&result)

		if err != nil {
			log.Fatal(err)
		}
		zone := time.FixedZone("", +8*60*60)
		result.CreateTime = result.CreateTime.In(zone)
		

		results = append(results, &result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(ctx)
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	fmt.Println(results)
	spew.Dump(results)
	return results, 1
}

func GetAllToDoList(ctx iris.Context) ([]*models.HaveIDTodoList, int) {

	var results []*models.HaveIDTodoList
	email := middleware.MyAuthenticatedHandler(ctx)
	if email == "token not found" {
		return results, 3
	}
	updateStatus(ctx,email)
	filter := bson.D{{Key: "email", Value: email}}

	cur, err := connect.TodolistCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {

		var result models.HaveIDTodoList
		err := cur.Decode(&result)

		if err != nil {
			log.Fatal(err)
		}
		zone := time.FixedZone("", +8*60*60)
		result.CreateTime = result.CreateTime.In(zone)
		

		results = append(results, &result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(ctx)
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	fmt.Println(results)
	spew.Dump(results)
	return results, 1
}

func UpdateToDoList(ctx iris.Context) int {
	var TodoList models.HaveIDTodoList

	email := middleware.MyAuthenticatedHandler(ctx)
	if email == "token not found" {
		return 3
	}
	if err := ctx.ReadJSON(&TodoList); err != nil {
		panic(err.Error())
	}
	id, _ := primitive.ObjectIDFromHex(TodoList.ID)

	filter := bson.D{{Key: "_id", Value: id}, {Key: "email", Value: email}}

	opts := options.Update().SetUpsert(false)
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "finishedCondition", Value: TodoList.FinishedCondition},
			{Key: "note", Value: TodoList.Note},
			{Key: "updateTime", Value: time.Now()},
		},
		}}
	result, err := connect.TodolistCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Fatal(err)
		return 2
	}
	if result.MatchedCount != 0 {
		fmt.Printf("Matched %v documents and updated %v documents.\n", result.MatchedCount, result.ModifiedCount)
		return 1
	}

	fmt.Println(result)
	return 2
}

func DeleteToDoList(ctx iris.Context) int {
	// var TodoList models.HaveIDTodoList
	email := middleware.MyAuthenticatedHandler(ctx)
	if email == "token not found" {
		return 3
	}


    paramsId := ctx.URLParam("id")
	fmt.Println(paramsId)

	id, _ := primitive.ObjectIDFromHex(paramsId)

	filter := bson.D{{Key: "_id", Value: id}, {Key: "email", Value: email}}

	deleteResult, err := connect.TodolistCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}


	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	if deleteResult.DeletedCount == 0 {
		return 2
	}
	return 1

}

func updateStatus(ctx iris.Context, email string) {
	var results []*models.HaveIDTodoList

	filter := bson.D{{Key: "email", Value: email}}

	cur, err := connect.TodolistCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {

		var result models.HaveIDTodoList
		err := cur.Decode(&result)

		if err != nil {
			log.Fatal(err)
		}
		zone := time.FixedZone("", +8*60*60)
		result.CreateTime = result.CreateTime.In(zone)
		id, _ := primitive.ObjectIDFromHex(result.ID)
		if result.EndTime.After(time.Now()) {
			filter := bson.D{{Key: "_id", Value: id}, {Key: "email", Value: result.Email}}
			opts := options.Update().SetUpsert(true)
			update := bson.D{
				{Key: "$set", Value: bson.D{
					{Key: "status", Value: "未逾期"},
				},
				}}
			connect.TodolistCollection.UpdateOne(ctx, filter, update, opts)

		} else {
			filter := bson.D{{Key: "_id", Value: id}, {Key: "email", Value: result.Email}}
			opts := options.Update().SetUpsert(true)
			update := bson.D{
				{Key: "$set", Value: bson.D{
					{Key: "status", Value: "逾期"},
				},
				}}
			connect.TodolistCollection.UpdateOne(ctx, filter, update, opts)

		}

		results = append(results, &result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(ctx)
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	// spew.Dump(results)
	// return results, "success"

}
