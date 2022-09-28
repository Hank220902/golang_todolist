package service

import(
	"fmt"
	"log"
	"todolist/database"
	"todolist/models"
	// "github.com/davecgh/go-spew/spew"
	"github.com/kataras/iris/v12"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"crypto/md5"
)

func Register(ctx iris.Context) string {
	var user models.User 
	if err := ctx.ReadJSON(&user); err != nil {
		panic(err.Error())
	}

	fmt.Println(user)

	result := models.User{user.Name,user.Password,user.Email}
	insertOne, err := database.UserCollection.InsertOne(ctx, result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Document: ", insertOne.InsertedID)
	return "註冊成功"
}

func Login(){

}