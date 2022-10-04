package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	//"time"
)

var TodolistCollection *mongo.Collection
var UserCollection *mongo.Collection


func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://Hank220902:asd8836259@cluster0.9mt5wf9.mongodb.net/?retryWrites=true&w=majority")
	var ctx = context.TODO()
	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	//defer client.Disconnect(ctx)

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)



	TodolistCollection = client.Database("todolist").Collection("todolists")
	UserCollection = client.Database("todolist").Collection("users")
	fmt.Println("Collection instance created!")

}


