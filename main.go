// file: main.go
package main

import (
	"todolist/controllers"
	"todolist/database"

	"github.com/kataras/iris/v12"
	//"github.com/kataras/iris/v12/context"
)

func init() {
    database.Connect()
}
    
func main() {
    
    app := iris.New()
    app.Get("/", func(ctx iris.Context) {
        ctx.JSON("Hello World")
      })
    app.Get("/todoList", controllers.CreateToDoList)
    app.Post("/todoList", controllers.GetToDoList)
    app.Get("/ManytodoList", controllers.GetManyToDoList)
    app.Put("/todoList", controllers.UpdateToDoList)
    app.Delete("/todoList", controllers.DeleteToDoList)
    app.Listen(":3000")


}




