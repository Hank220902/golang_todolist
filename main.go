// file: main.go
package main

import (
    "todolist/routes"
	"todolist/database"
	"github.com/kataras/iris/v12"

)

func init() {
    database.Connect()
}
    
func main() {
    
    app := iris.New()
    
    routes.Route(app)
    app.Listen(":3000")


}




