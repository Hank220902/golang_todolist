package controllers

import (
	"fmt"
	"todolist/service"

	"github.com/kataras/iris/v12"
)

func Register(ctx iris.Context){
	result :=service.Register(ctx)
	fmt.Println(result)
	ctx.JSON(result)
}

func Login(ctx iris.Context){
	
	result := service.Login(ctx)
	fmt.Println(result)
	ctx.JSON(result)
}