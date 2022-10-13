package routes

import (
	"log"
	"todolist/controllers"
	"todolist/middleware"
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

func TodoList(app *iris.Application) {

	api := app.Party("/api", middleware.J.Serve)
	{
		api.Get("/todolist", controllers.GetFilterToDoList)
		api.Post("/todolist", controllers.CreateToDoList)
		api.Get("/manytodolist", controllers.GetAllToDoList)
		api.Put("/todolist", controllers.UpdateToDoList)
		api.Delete("/todolist", controllers.DeleteToDoList)

	}
	grpc:=app.Party("/grpc",middleware.J.Serve)
	{
		grpc.Post("/todolist", controllers.GrpcCreateToDoList)
		grpc.Get("/manytodolist", controllers.GrpcGetManyToDoList)
		grpc.Put("/todolist", controllers.GrpcUpdateToDoList)
		grpc.Delete("/todolist", controllers.GrpcDeleteToDoList)
	}


	app.Get("/", func(ctx iris.Context) {
		ctx.JSON("Hello World")
	})
}

// func test(ctx iris.Context) {
	// req := pb.CreateRequest{Matter: string("hello"), EndTime: string("10/10"), FinishedCondition: string("not")}
// 	res, err := client.CreateTodolist(ctx, &req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(res.GetResMessage())
// }
