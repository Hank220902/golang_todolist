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
	result,str := service.GetOneToDoList(ctx)
	fmt.Println(result)
	if str == "success"{
		ctx.JSON(result)
	}else{
		ctx.JSON(str)
	}
	
}

func GetManyToDoList(ctx iris.Context) {
	result,str := service.GetManyToDoList(ctx)
	if str == "success"{
		ctx.JSON(result)
	}else{
		ctx.JSON(str)
	}
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
