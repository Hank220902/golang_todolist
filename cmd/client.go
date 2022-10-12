package cmd

import (
	"context"
	"fmt"

	pb "todolist/protobuf"

	"google.golang.org/grpc"
)

func Client() {
	cli, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	c := pb.NewTodoListServiceClient(cli)

	req := pb.CreateRequest{Matter: string("hello"), EndTime: string("10/10"), FinishedCondition: string("not")}
	res, err := c.CreateTodolist(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.GetResMessage())

}
