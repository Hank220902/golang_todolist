package connect

import (
	"google.golang.org/grpc"
	pb "todolist/protobuf"
)

var client pb.TodoListServiceClient

func Grpc() {
	cli, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	client =pb.NewTodoListServiceClient(cli)
}
