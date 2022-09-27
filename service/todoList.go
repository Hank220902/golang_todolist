package service

import (
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	"log"
	"todolist/database"
	"todolist/model"

	"github.com/davecgh/go-spew/spew"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateToDoList(ctx iris.Context) model.Sunshareboy {
	//path := ctx.Path()
	//app.Logger().Info("请求的路径是:"+path)
	//json数据解析
	var user model.Sunshareboy //需要用一个结构体来存放数据 并且结构体当中要有tag标签
	//context.ReadJson() 里面传入的是结构体的指针类型 内存地址
	if err := ctx.ReadJSON(&user); err != nil {
		panic(err.Error())
	}

	wanger := model.Sunshareboy{user.Name, user.Age, user.City}
	insertOne, err := database.Collection.InsertOne(ctx, wanger)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Document: ", insertOne.InsertedID)
	return wanger

}

func GetOneToDoList(ctx iris.Context) model.Sunshareboy {
	var result model.Sunshareboy
	name := ctx.Request().URL.Query().Get("Name")

	filter := bson.D{{"name", name}}
	err := database.Collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)
	return result
}

func GetManyToDoList(ctx iris.Context) []*model.Sunshareboy {
	findOptions := options.Find()
	findOptions.SetLimit(10)
	var results []*model.Sunshareboy
	cur, err := database.Collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {
		//定義一個文件，將單個文件解碼為result
		var result model.Sunshareboy
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

func UpdateToDoList(ctx iris.Context) {
	var user model.Sunshareboy //需要用一个结构体来存放数据 并且结构体当中要有tag标签
	//context.ReadJson() 里面传入的是结构体的指针类型 内存地址
	if err := ctx.ReadJSON(&user); err != nil {
		panic(err.Error())
	}
	filter := bson.D{{"name", user.Name}}
	// 如果過濾的文件不存在，則插入新的文件
	opts := options.Update().SetUpsert(true)
	update := bson.D{
		{"$set", bson.D{
			{"city", user.City}},
		}}
	result, err := database.Collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Fatal(err)
	}
	if result.MatchedCount != 0 {
		fmt.Printf("Matched %v documents and updated %v documents.\n", result.MatchedCount, result.ModifiedCount)
	}
	if result.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", result.UpsertedID)
	}
	fmt.Println(result)

}

func DeleteToDoList(ctx iris.Context) {
	var user model.Sunshareboy //需要用一个结构体来存放数据 并且结构体当中要有tag标签
	//context.ReadJson() 里面传入的是结构体的指针类型 内存地址
	if err := ctx.ReadJSON(&user); err != nil {
		panic(err.Error())
	}
	fmt.Println(user)
	filter := bson.D{{"city", user.City}}
	deleteResult, err := database.Collection.DeleteMany(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

}
