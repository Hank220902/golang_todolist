// file: main.go
package main

import (
    "todolist/routes"
	"todolist/database"
	"github.com/kataras/iris/v12"
    // "todolist/middleware"

)

func init() {
    database.Connect()
    database.Redis()
}
    
func main() {
    
    app := iris.New()
    
    routes.TodoList(app)
    routes.User(app)
    app.Listen(":3000")


}




