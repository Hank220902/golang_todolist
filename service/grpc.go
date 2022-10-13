package service

import (
	"fmt"
	"log"
	"todolist/middleware"
	"todolist/models"
	pb "todolist/protobuf"

	"github.com/kataras/iris/v12"
	"google.golang.org/grpc"
)

var client pb.TodoListServiceClient

func init() {
	connect, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	client = pb.NewTodoListServiceClient(connect)

}

func GCreateToDoList(ctx iris.Context) int32 {
	email := middleware.MyAuthenticatedHandler(ctx)
	if email == "token not found" {
		return 3
	}

	var TodoList models.Todo

	if err := ctx.ReadJSON(&TodoList); err != nil {
		panic(err.Error())
	}
	endtime := TodoList.EndTime
	fmt.Println(endtime)
	status := ""

	req := pb.CreateRequest{
		Matter:            TodoList.Matter,
		EndTime:           TodoList.EndTime,
		FinishedCondition: TodoList.FinishedCondition,
		Status:            status,
		Email:            email,
	}
	res, err := client.CreateTodolist(ctx, &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.GetResMessage())
	return res.GetResMessage()

}

func GGetManyToDoList(ctx iris.Context)(*pb.GetResponse, int){

	email := middleware.MyAuthenticatedHandler(ctx)
	if email == "token not found" {
		return nil,3
	}
	updateStatus(ctx,email)
	req :=pb.GetRequest{
		Email: email,
	}
	res,err :=client.GetManyTodolist(ctx,&req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	return res,1
}

func GUpdateToDoList(ctx iris.Context)int32{
	var TodoList models.HaveIDTodoList
	email := middleware.MyAuthenticatedHandler(ctx)
	if email == "token not found" {
		return 3
	}
	if err := ctx.ReadJSON(&TodoList); err != nil {
		panic(err.Error())
	}
	req :=pb.UpdateRequest{
		Id: TodoList.ID,
		FinishedCondition: TodoList.FinishedCondition,
		Note:TodoList.Note,
		Email: email,
	}
	res,err :=client.UpdateTodolist(ctx,&req) 
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	return res.ResMessage
}

func GDeleteToDoList(ctx iris.Context)int32{
	email := middleware.MyAuthenticatedHandler(ctx)
	if email == "token not found" {
		return 3
	}


    paramsId := ctx.URLParam("id")
	fmt.Println(paramsId)
	req:=pb.DeleteRequest{
		Id: paramsId,
		Email: email,
	}
	res,err := client.DeleteTodolist(ctx,&req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	return res.ResMessage

}
