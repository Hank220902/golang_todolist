package controllers

import (
	"fmt"
	"todolist/service"
	"github.com/kataras/iris/v12"
)



func CreateToDoList(ctx iris.Context) {
	// service.CreateToDoList()
	hello :=service.CreateToDoList(ctx)
	fmt.Println(hello)
	ctx.JSON(hello)
	
}
