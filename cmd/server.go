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

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	result := models.Todo{
		Matter:            req.GetMatter(),
		EndTime:           req.GetEndTime(),
		FinishedCondition: req.GetFinishedCondition(),
		Status:            req.GetStatus(),
		Email:             req.GetEmail(),
		CreateTime:        time.Now().Local(),
	}
	insertOne, err := connect.TodolistCollection.InsertOne(ctx, result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Document: ", insertOne.InsertedID)
	fmt.Println(result)
	return &pb.CreateResponse{ResMessage: 1}, nil
}

func (ts *todoListService)GetFilterTodolist(ctx context.Context,req *pb.GetFilterRequest) (*pb.GetFilterResponse, error) {
	var results []*models.HaveIDTodoList
	filter := bson.D{{Key: "email", Value: req.GetEmail()}}
	if req.GetStatus() != "" && req.GetFinishedCondition() != ""{
		filter =bson.D{{Key: "email", Value: req.GetEmail()},{Key: "status", Value: req.GetStatus()},{Key: "finishedCondition", Value: req.GetFinishedCondition()}}
	}else if req.GetStatus() != ""{
		filter =bson.D{{Key: "email", Value: req.GetEmail()},{Key: "status", Value: req.GetStatus()}}
	}else if req.GetFinishedCondition() != ""{
		filter =bson.D{{Key: "email", Value: req.GetEmail()},{Key: "finishedCondition", Value: req.GetFinishedCondition()}}
	}
	cur, err := connect.TodolistCollection.Find(ctx, filter)
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
	return &pb.GetFilterResponse{GetResult: convertGetToDoListsToPb(results)}, nil
}

func (ts *todoListService) GetAllTodolist(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	var results []*models.HaveIDTodoList
	filter := bson.D{{Key: "email", Value: req.GetEmail()}}

	cur, err := connect.TodolistCollection.Find(ctx, filter)
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
	return &pb.GetResponse{GetResult: convertGetToDoListsToPb(results)}, nil
}

func convertGetToDoListsToPb(in []*models.HaveIDTodoList) []*pb.GetResult {
	ans := make([]*pb.GetResult, len(in))
	for i, v := range in {
		ans[i] = convertGetToDoListToPb(v)

	}
	return ans
}

func convertGetToDoListToPb(in *models.HaveIDTodoList) *pb.GetResult {
	if in == nil {
		return new(pb.GetResult)
	}
	return &pb.GetResult{
		Id:                in.ID,
		Matter:            in.Matter,
		EndTime:           in.EndTime.String(),
		FinishedCondition: in.FinishedCondition,
		CreateTime:        in.CreateTime.String(),
		Status:            in.Status,
		Note:              in.Note,
		Email:             in.Email,
		UpdateTime:        in.UpeateTime.String(),
	}
}

func (ts *todoListService) UpdateTodolist(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	id, _ := primitive.ObjectIDFromHex(req.GetId())

	filter := bson.D{{Key: "_id", Value: id}, {Key: "email", Value: req.GetEmail()}}

	opts := options.Update().SetUpsert(false)
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "finishedCondition", Value: req.GetFinishedCondition()},
			{Key: "note", Value: req.GetNote()},
			{Key: "updateTime", Value: time.Now()},
		},
		}}
	fmt.Println(req.GetEmail(), req.GetId(), req.GetFinishedCondition())
	result, err := connect.TodolistCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Fatal(err)
		return &pb.UpdateResponse{ResMessage: 2}, nil
	}
	if result.MatchedCount != 0 {
		fmt.Printf("Matched %v documents and updated %v documents.\n", result.MatchedCount, result.ModifiedCount)
		return &pb.UpdateResponse{ResMessage: 1}, nil
	}

	fmt.Println(result)
	return &pb.UpdateResponse{ResMessage: 2}, nil
}

func (ts *todoListService) DeleteTodolist(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	id, _ := primitive.ObjectIDFromHex(req.GetId())

	filter := bson.D{{Key: "_id", Value: id}, {Key: "email", Value: req.GetEmail()}}

	deleteResult, err := connect.TodolistCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	if deleteResult.DeletedCount == 0 {
		return &pb.DeleteResponse{ResMessage: 2}, nil
	}
	return &pb.DeleteResponse{ResMessage: 1}, nil
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
