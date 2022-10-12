// file: main.go
package main

import (
	"fmt"
	"todolist/cmd"
	"todolist/connect"
	"todolist/routes"

	"github.com/kataras/iris/v12"
	// "todolist/middleware"
)


func init() {

	connect.Connect()
	connect.Redis()
    // cmd.Client()
    go cmd.Server()
    

}

func main() {
   
	app := iris.New()
    fmt.Println("hello")

	routes.TodoList(app)
	routes.User(app)
 

	app.Listen(":3000")


}
