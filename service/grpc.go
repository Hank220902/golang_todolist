package service

import (
	"fmt"
	"log"
	"time"
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

func GCreateToDoList(ctx iris.Context) int {
	email := middleware.MyAuthenticatedHandler(ctx)
	if email == "token not found" {
		return 3
	}
	var TodoList models.TodoList

	if err := ctx.ReadJSON(&TodoList); err != nil {
		panic(err.Error())
	}

	status := ""
	createTime := time.Now()
	req := pb.CreateRequest{
		Matter:            string(TodoList.Matter),
		EndTime:           string(TodoList.EndTime.String()),
		FinishedCondition: string(TodoList.FinishedCondition),
		Status:            string(status),
		CreateTime:        string(createTime.String()),
	}
	res, err := client.CreateTodolist(ctx, &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.GetResMessage())
	return 1

}
