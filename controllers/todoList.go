package controllers

import (
	"fmt"
	"todolist/service"

	"github.com/kataras/iris/v12"
)

func CreateToDoList(ctx iris.Context) {
	// service.CreateToDoList()

	result := service.CreateToDoList(ctx)
	fmt.Println(result)
	ctx.JSON(result)

}

func GetToDoList(ctx iris.Context) {
	result := service.GetOneToDoList(ctx)
	fmt.Println(result)
	ctx.JSON(result)
}

func GetManyToDoList(ctx iris.Context) {
	result := service.GetManyToDoList(ctx)
	fmt.Println(result)
	ctx.JSON(result)
}

func UpdateToDoList(ctx iris.Context) {
	result := service.UpdateToDoList(ctx)
	fmt.Println()
	ctx.JSON(result)
}

func DeleteToDoList(ctx iris.Context) {
	result := service.DeleteToDoList(ctx)
	ctx.JSON(result)
}
