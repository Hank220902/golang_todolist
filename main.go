// file: main.go
package main

import (
    "todolist/controllers"
    "todolist/database"
	"github.com/kataras/iris/v12")

func main() {
    database.Database()
    app := iris.New()
    app.Get("/", func(ctx iris.Context) {
        ctx.JSON("Hello World")
      })
    app.Get("/todoList", controllers.CreateToDoList)
    // app.Post("/todoList", posting)
    // app.Put("/todoList", putting)
    // app.Delete("/todoList", deleting)
    app.Listen(":8080")


}




