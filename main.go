// file: main.go
package main

import (
	"todolist/controllers"
	"todolist/database"

	"github.com/kataras/iris/v12"

)

func init() {
    database.Connect()
}
    
func main() {
    
    app := iris.New()
    app.Get("/", func(ctx iris.Context) {
        ctx.JSON("Hello World")
      })
    app.Get("/todolist", controllers.GetToDoList)
    app.Post("/todolist", controllers.CreateToDoList)
    app.Get("/manytodolist", controllers.GetManyToDoList)
    app.Put("/todolist", controllers.UpdateToDoList)
    app.Delete("/todolist", controllers.DeleteToDoList)
    app.Listen(":3000")


}




