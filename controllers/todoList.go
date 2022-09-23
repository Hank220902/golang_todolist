package controllers

import (
	"fmt"
	"todolist/service"
	"github.com/kataras/iris/v12"
)



func CreateToDoList(ctx iris.Context) {
	fmt.Println("123")
	// service.CreateToDoList()
	fmt.Println(service.CreateToDoList())
	ctx.JSON(service.CreateToDoList())
	
}
