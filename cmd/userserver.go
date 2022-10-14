package cmd

import (
	"context"
	"fmt"
	"log"

	"todolist/connect"
	"todolist/middleware"
	"todolist/models"
	pb "todolist/protobuf"
	"todolist/service"

	"go.mongodb.org/mongo-driver/bson"

)

type userService struct {
	pb.UnimplementedUserServiceServer
}

func (us *userService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	hashStr, err := service.HashAndSalt(req.GetPassword())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hashStr)

	result := models.User{
		Name:     req.GetName(),
		Password: hashStr,
		Email:    req.GetEmail(),
	}
	insertOne, err := connect.UserCollection.InsertOne(ctx, result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Document: ", insertOne.InsertedID)
	return &pb.RegisterResponse{ResMessage: 1}, nil
}

func (us *userService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var resultLogin models.Login
	filter := bson.D{{Key: "email", Value: req.GetEmail()}}
	err := connect.UserCollection.FindOne(ctx, filter).Decode(&resultLogin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(req.GetEmail())

	token := middleware.GetTokenHandler(req.GetEmail())
	check := service.ComparePasswords(resultLogin.Password, req.GetPassword())
	if !check {
		return &pb.LoginResponse{ResMessage: "fail"}, nil
	} else {
		return &pb.LoginResponse{ResMessage: token}, nil
	}
}

func (us *userService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	var Redis = connect.Redis()
	n, err := Redis.Exists(ctx, req.GetEmail()).Result()
	if err != nil {
		panic(err)
	}
	if n > 0 {
		n := Redis.Del(ctx, req.GetEmail())
		fmt.Println(n.Result())
		return &pb.LogoutResponse{ResMessage: 1}, nil
	}
	return &pb.LogoutResponse{ResMessage: 2}, nil
}


