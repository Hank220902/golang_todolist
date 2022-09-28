package routes

import (
	"todolist/controllers"
	

	"github.com/kataras/iris/v12"

)

func User (app *iris.Application){
	app.Post("/register",controllers.Register)
	app.Post("/login",controllers.Login)
}