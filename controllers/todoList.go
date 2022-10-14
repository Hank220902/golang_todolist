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
	if result == 1 {
		ctx.JSON("success")
	} else if result == 3 {
		ctx.JSON("tokenError")
	} else {
		ctx.JSON("error")
	}

}

func GetFilterToDoList(ctx iris.Context) {
	result, num := service.GetFilterToDoList(ctx)
	fmt.Println(result)
	if num == 1 {
		ctx.JSON(result)
	} else if num == 3 {
		ctx.JSON("tokenError")
	} else {
		ctx.JSON("error")
	}

}

func GetAllToDoList(ctx iris.Context) {
	result, num := service.GetAllToDoList(ctx)
	if num == 1 {
		ctx.JSON(result)
	} else if num == 3 {
		ctx.JSON("tokenError")
	} else {
		ctx.JSON("error")
	}
}

func UpdateToDoList(ctx iris.Context) {
	result := service.UpdateToDoList(ctx)
	fmt.Println()
	ctx.JSON(result)
}

func DeleteToDoList(ctx iris.Context) {
	result := service.DeleteToDoList(ctx)
	if result == 1 {
		ctx.JSON("success")
	} else if result == 3 {
		ctx.JSON("tokenError")
	} else {
		ctx.JSON("error")
	}

}
