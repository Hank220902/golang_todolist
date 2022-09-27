package controllers

import (
	"fmt"
	"todolist/service"

	"github.com/kataras/iris/v12"
)

func CreateToDoList(ctx iris.Context) {
	// service.CreateToDoList()

	hello := service.CreateToDoList(ctx)
	fmt.Println(hello)
	ctx.JSON(hello)

}

func GetToDoList(ctx iris.Context) {
	hello := service.GetOneToDoList(ctx)
	fmt.Println(hello)
	ctx.JSON(hello)
}

func GetManyToDoList(ctx iris.Context) {
	hello := service.GetManyToDoList(ctx)
	fmt.Println(hello)
	ctx.JSON(hello)
}

func UpdateToDoList(ctx iris.Context) {
	service.UpdateToDoList(ctx)
	// hello := service.UpdateToDoList(ctx)
	// fmt.Println()
	// ctx.JSON(hello)
}

func DeleteToDoList(ctx iris.Context) {
	service.DeleteToDoList(ctx)
}
