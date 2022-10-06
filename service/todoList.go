package service

import (
	"fmt"
	"log"
	"strconv"
	"time"
	"todolist/database"
	"todolist/middleware"
	"todolist/models"

	"github.com/davecgh/go-spew/spew"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateToDoList(ctx iris.Context) string {
	email := middleware.MyAuthenticatedHandler(ctx)
	if email == "token not found" {
		return "token not found"
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
	insertOne, err := database.TodolistCollection.InsertOne(ctx, result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Document: ", insertOne.InsertedID)
	return "success"

}

func GetOneToDoList(ctx iris.Context) (models.HaveIDTodoList, string) {
	var result models.HaveIDTodoList

	email := middleware.MyAuthenticatedHandler(ctx)
	if email == "token not found" {
		return result, "token not found"
	}
	fmt.Println(email)
	filter := bson.D{{Key: "email", Value: email}}
	err := database.TodolistCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)
	return result, "success"
}

func GetManyToDoList(ctx iris.Context) ([]*models.HaveIDTodoList, string) {

	var results []*models.HaveIDTodoList
	email := middleware.MyAuthenticatedHandler(ctx)
	if email == "token not found" {
		return results, "token not found"
	}
	updateStatus(ctx,email)
	filter := bson.D{{Key: "email", Value: email}}

	cur, err := database.TodolistCollection.Find(ctx, filter)
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

	spew.Dump(results)
	return results, "success"
}

func UpdateToDoList(ctx iris.Context) string {
	var TodoList models.HaveIDTodoList

	email := middleware.MyAuthenticatedHandler(ctx)
	if email == "token not found" {
		return "token not found"
	}
	if err := ctx.ReadJSON(&TodoList); err != nil {
		panic(err.Error())
	}
	id, _ := primitive.ObjectIDFromHex(TodoList.ID)

	filter := bson.D{{Key: "_id", Value: id}, {Key: "email", Value: email}}

	opts := options.Update().SetUpsert(true)
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "finishedCondition", Value: TodoList.FinishedCondition},
			{Key: "note", Value: TodoList.Note},
			{Key: "updateAt", Value: time.Now()},
		},
		}}
	result, err := database.TodolistCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Fatal(err)
		return "update failed"
	}
	if result.MatchedCount != 0 {
		fmt.Printf("Matched %v documents and updated %v documents.\n", result.MatchedCount, result.ModifiedCount)
		return "update success"
	}
	if result.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", result.UpsertedID)
		return "insert success"
	}
	fmt.Println(result)
	return "update failed"
}

func DeleteToDoList(ctx iris.Context) string {
	var TodoList models.HaveIDTodoList
	email := middleware.MyAuthenticatedHandler(ctx)
	if email == "token not found" {
		return "token not found"
	}
	if err := ctx.ReadJSON(&TodoList); err != nil {
		panic(err.Error())
	}
	fmt.Println(TodoList)
	id, _ := primitive.ObjectIDFromHex(TodoList.ID)

	filter := bson.D{{Key: "_id", Value: id}, {Key: "email", Value: email}}

	deleteResult, err := database.TodolistCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	result := strconv.FormatInt(deleteResult.DeletedCount, 10)

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	result = "刪除" + result + "筆資料"
	return result

}

func updateStatus(ctx iris.Context, email string) {
	var results []*models.HaveIDTodoList

	filter := bson.D{{Key: "email", Value: email}}

	cur, err := database.TodolistCollection.Find(ctx, filter)
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
		if result.EndTime.Before(time.Now()) {
			filter := bson.D{{Key: "_id", Value: id}, {Key: "email", Value: result.Email}}
			opts := options.Update().SetUpsert(true)
			update := bson.D{
				{Key: "$set", Value: bson.D{
					{Key: "status", Value: "未逾期"},
				},
				}}
			database.TodolistCollection.UpdateOne(ctx, filter, update, opts)

		} else {
			filter := bson.D{{Key: "_id", Value: id}, {Key: "email", Value: result.Email}}
			opts := options.Update().SetUpsert(true)
			update := bson.D{
				{Key: "$set", Value: bson.D{
					{Key: "status", Value: "逾期"},
				},
				}}
			database.TodolistCollection.UpdateOne(ctx, filter, update, opts)

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
