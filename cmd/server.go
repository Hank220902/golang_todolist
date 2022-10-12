package cmd

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"
	"todolist/connect"
	"todolist/models"
	pb "todolist/protobuf"

	"google.golang.org/grpc"
)

// type EchoServiceServer interface {
// 	// Unary
// 	SayHello(context.Context, *EchoRequest) (*EchoResponse, error)
// }

type todoListService struct {
	pb.UnimplementedTodoListServiceServer
}

func (ts *todoListService) CreateTodolist(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	// res := "recovide:" + req.GetMatter() + req.GetEndTime() + req.GetFinishedCondition()
	layout:="2022-09-30 03:31:31.72"
	endTime, error := time.Parse(layout, req.GetEndTime())
	createTime, error := time.Parse(layout, req.GetCreateTime())
	fmt.Println(req.GetEndTime())
	fmt.Println(endTime)

	if error != nil {
		fmt.Println(error)
		panic(error)
	}
	result := models.TodoList{
		Matter:            req.GetMatter(),
		EndTime:           endTime,
		FinishedCondition: req.GetFinishedCondition(),
		Status:            req.GetStatus(),
		Email:             req.GetEmail(),
		CreateTime:        createTime,
	}
	insertOne, err := connect.TodolistCollection.InsertOne(ctx, result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Document: ", insertOne.InsertedID)
	fmt.Println(result)
	return &pb.CreateResponse{ResMessage: "success"}, nil
}

func Server() {
	rpcs := grpc.NewServer()
	pb.RegisterTodoListServiceServer(rpcs, new(todoListService))
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	rpcs.Serve(lis)

}
