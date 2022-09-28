package routes

import (
	"todolist/controllers"
	

	"github.com/kataras/iris/v12"

)

func Route (app *iris.Application){
	app.Get("/", func(ctx iris.Context) {
        ctx.JSON("Hello World")
      })
	app.Get("/todolist", controllers.GetToDoList)
    app.Post("/todolist", controllers.CreateToDoList)
    app.Get("/manytodolist", controllers.GetManyToDoList)
    app.Put("/todolist", controllers.UpdateToDoList)
    app.Delete("/todolist", controllers.DeleteToDoList)
}