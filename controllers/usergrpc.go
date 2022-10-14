package controllers

import (
	"fmt"
	"todolist/service"

	"github.com/kataras/iris/v12"
)

func GrpcRegister(ctx iris.Context){
	result :=service.GRegister(ctx)
	fmt.Println(result)
	if result ==1{
		ctx.JSON("success")
	}else{
		ctx.JSON("emailExists")
	}

}

func GrpcLogin(ctx iris.Context){
	
	result := service.GLogin(ctx)
	fmt.Println(result)
	ctx.JSON(result)
}

func GrpcLogout(ctx iris.Context){
	result := service.GLogout(ctx)
	if result ==1{
		ctx.JSON("success")
	}else{
		ctx.JSON("error")
	}
}