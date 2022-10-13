package service

import (
	"fmt"
	"log"
	"todolist/connect"
	"todolist/models"
	"todolist/middleware"
	// "github.com/davecgh/go-spew/spew"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/options"
)


func Register(ctx iris.Context) int {

	var user models.User
	if err := ctx.ReadJSON(&user); err != nil {
		panic(err.Error())
	}

	fmt.Println(user)

	if haveEmail(ctx,user.Email)==4{
		return 4
	}
	hashStr, err := HashAndSalt(user.Password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hashStr)

	result := models.User{
		Name:     user.Name,
		Password: hashStr,
		Email:    user.Email,
	}
	insertOne, err := connect.UserCollection.InsertOne(ctx, result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Document: ", insertOne.InsertedID)
	return 1
}

func haveEmail(ctx iris.Context, email string) int{
	var reqLogin models.Login
	fmt.Println(reqLogin.Password)

	var resultLogin models.Login
	filter := bson.D{{Key: "email", Value: email}}
	err := connect.UserCollection.FindOne(ctx, filter).Decode(&resultLogin)
	if err != nil {
		return 1
	}else{
		return 4
	}
}

func Login(ctx iris.Context) string {
	var reqLogin models.Login
	if err := ctx.ReadJSON(&reqLogin); err != nil {
		panic(err.Error())
	}
	fmt.Println(reqLogin.Password)
	email := reqLogin.Email


	password := reqLogin.Password
	var resultLogin models.Login
	filter := bson.D{{Key: "email", Value: email}}
	err := connect.UserCollection.FindOne(ctx, filter).Decode(&resultLogin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(email)

	token := middleware.GetTokenHandler(email)
	check := ComparePasswords(resultLogin.Password, password)
	if !check {
		return "fail"
	} else {
		return token
	}

}

func Logout(ctx iris.Context)int{
	var Redis = connect.Redis()
	email := middleware.MyAuthenticatedHandler(ctx)
	n, err := Redis.Exists(ctx, email).Result()
	if err != nil {
		panic(err)
	}
	if n > 0 {
		n := Redis.Del(ctx, email)
		fmt.Println(n.Result())
		return 1
	}
	return 2
}