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

var userClient pb.UserServiceClient

func init() {
	connect, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	userClient = pb.NewUserServiceClient(connect)
}

func GRegister(ctx iris.Context) int32 {

	var user models.User
	if err := ctx.ReadJSON(&user); err != nil {
		panic(err.Error())
	}
	if haveEmail(ctx, user.Email) == 4 {
		return 4
	}
	req := pb.RegisterRequest{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	res, err := userClient.Register(ctx, &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.GetResMessage())

	return res.GetResMessage()
}

func GLogin(ctx iris.Context) string {
	var reqLogin models.Login
	if err := ctx.ReadJSON(&reqLogin); err != nil {
		panic(err.Error())
	}
	email := reqLogin.Email

	password := reqLogin.Password

	req := pb.LoginRequest{
		Email:    email,
		Password: password,
	}
	res, err := userClient.Login(ctx, &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.GetResMessage())
	return res.GetResMessage()

}

func GLogout(ctx iris.Context) int32 {

	email := middleware.MyAuthenticatedHandler(ctx)
	req := pb.LogoutRequest{
		Email: email,
	}
	res, err := userClient.Logout(ctx, &req)
	if err != nil {
		panic(err)
	}
	return res.GetResMessage()

}
