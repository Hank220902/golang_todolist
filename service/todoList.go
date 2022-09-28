package service

import (
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	"log"
	"todolist/database"
	"todolist/model"

	"time"
	"strconv"
	"github.com/davecgh/go-spew/spew"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateToDoList(ctx iris.Context) string {

	//需要用一个结构体来｀存放数据 并且结构体当中要有tag标签
	var TodoList model.TodoList

	if err := ctx.ReadJSON(&TodoList); err != nil {
		panic(err.Error())
	}

	fmt.Println(TodoList)
	TodoList.CreateTime = time.Now()

	result := model.TodoList{TodoList.Matter, TodoList.CreateTime, TodoList.EndTime, TodoList.FinishedCondition, TodoList.Status, TodoList.Email}
	insertOne, err := database.TodolistCollection.InsertOne(ctx, result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Document: ", insertOne.InsertedID)
	return "新增成功"

}

func GetOneToDoList(ctx iris.Context) model.TodoList {
	var result model.TodoList
	matter := ctx.Request().URL.Query().Get("matter")

	filter := bson.D{{"matter", matter}}
	err := database.TodolistCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)
	return result
}

func GetManyToDoList(ctx iris.Context) []*model.TodoList {
	findOptions := options.Find()
	findOptions.SetLimit(10)
	var results []*model.TodoList
	cur, err := database.TodolistCollection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {
		//定義一個文件，將單個文件解碼為result
		var result model.TodoList
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &result)
		fmt.Println(result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	// 遍歷結束後關閉遊標
	cur.Close(ctx)
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	//fmt.Println(results)
	spew.Dump(results)
	return results
}

func UpdateToDoList(ctx iris.Context) string {
	var TodoList model.TodoList //需要用一个结构体来存放数据 并且结构体当中要有tag标签
	//context.ReadJson() 里面传入的是结构体的指针类型 内存地址
	if err := ctx.ReadJSON(&TodoList); err != nil {
		panic(err.Error())
	}
	filter := bson.D{{"matter", TodoList.Matter}}
	// 如果過濾的文件不存在，則插入新的文件
	opts := options.Update().SetUpsert(true)
	update := bson.D{
		{"$set", bson.D{
			{"finishedCondition", TodoList.FinishedCondition}},
		}}
	result, err := database.TodolistCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Fatal(err)
		return "修改失敗"
	}
	if result.MatchedCount != 0 {
		fmt.Printf("Matched %v documents and updated %v documents.\n", result.MatchedCount, result.ModifiedCount)
		return "修改成功"		
	}
	if result.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", result.UpsertedID)
		return "修改成功"
	}
	fmt.Println(result)
	return "未知"
}

func DeleteToDoList(ctx iris.Context) (string) {
	var TodoList model.TodoList //需要用一个结构体来存放数据 并且结构体当中要有tag标签
	//context.ReadJson() 里面传入的是结构体的指针类型 内存地址
	if err := ctx.ReadJSON(&TodoList); err != nil {
		panic(err.Error())
	}
	fmt.Println(TodoList)

	fmt.Println(TodoList.FinishedCondition)
	deleteResult, err := database.TodolistCollection.DeleteMany(ctx, bson.M{"finishedCondition": TodoList.FinishedCondition})
	if err != nil {
		log.Fatal(err)
	}
	result := strconv.FormatInt(deleteResult.DeletedCount,10)

	
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	result = "刪除"+result+"筆資料"
	return result


}
