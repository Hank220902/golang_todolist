package service

import (
	"todolist/middleware"
	"fmt"
	"log"
	"todolist/database"
	"todolist/models"

	// "github.com/davecgh/go-spew/spew"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func Register(ctx iris.Context) string {
	var user models.User
	if err := ctx.ReadJSON(&user); err != nil {
		panic(err.Error())
	}

	fmt.Println(user)

	hashStr, err := HashAndSalt(user.Password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hashStr)

	result := models.User{
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
	}
	insertOne, err := database.UserCollection.InsertOne(ctx, result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Document: ", insertOne.InsertedID)
	return "註冊成功"
}

func Login(ctx iris.Context) string {
	var reqLogin models.Login
	if err := ctx.ReadJSON(&reqLogin); err != nil {
		panic(err.Error())
	}
	fmt.Println(reqLogin.Password)
	email := reqLogin.Email
	password := reqLogin.Password
	var resultLongin models.Login
	filter := bson.D{{Key: "email", Value: email}}
	err := database.UserCollection.FindOne(ctx, filter).Decode(&resultLongin)
	if err != nil {
		log.Fatal(err)
	}

	token := middleware.GetTokenHandler(email)

	fmt.Println(resultLongin)
	if password==resultLongin.Password{
		return token
	}else{
		return "密碼錯誤"
	}
	//fmt.Printf("Found a single document: %+v\n", login.Password)
	// hashStr, err := HashAndSalt(password)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(hashStr)

	// check := ComparePasswords(resultLongin.Password, password)
	// if !check {
	// 	fmt.Println("pw wrong")
	// 	return "帳號或密碼錯誤"
	// } else {
	// 	fmt.Println("pw ok")
	// 	return "登入成功"
	// }

}
