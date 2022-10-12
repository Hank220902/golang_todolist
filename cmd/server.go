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
	// layout:="2006-01-02T15:04:05.000Z"
	// result, err = time.Parse(time.RFC3339, in)
	endTime, error := time.Parse(time.RFC3339, req.GetEndTime())
	createTime, error := time.Parse(time.RFC3339, req.GetCreateTime())
	fmt.Println(req.GetEndTime())
	fmt.Println(endTime)
	fmt.Println(req.GetCreateTime())
	fmt.Println(createTime)

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
