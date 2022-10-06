package routes

import (
	"todolist/controllers"
	"todolist/middleware"

	"github.com/kataras/iris/v12"
)

func TodoList(app *iris.Application) {


	api := app.Party("/api", middleware.J.Serve)
	{
		api.Get("/todolist", controllers.GetToDoList)
		api.Post("/todolist", controllers.CreateToDoList)
		api.Get("/manytodolist", controllers.GetManyToDoList)
		api.Put("/todolist", controllers.UpdateToDoList)
		api.Delete("/todolist", controllers.DeleteToDoList)

	}

	app.Get("/", func(ctx iris.Context) {
		ctx.JSON("Hello World")
	})
}
