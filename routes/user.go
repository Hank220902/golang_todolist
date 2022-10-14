package routes

import (
	"todolist/controllers"
	"todolist/middleware"

	"github.com/kataras/iris/v12"
)

func User(app *iris.Application) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	user := app.Party("/logout", middleware.J.Serve)
	{
		user.Get("", controllers.Logout)
	}
	
	grpc:=app.Party("/grpc")
	{
		grpc.Post("/register",controllers.GrpcRegister)
		grpc.Post("/login",controllers.GrpcLogin)
	}
	grpcUser :=app.Party("/gepc",middleware.J.Serve)
	{
		grpcUser.Get("/logout",controllers.GrpcLogout)
	}
}
