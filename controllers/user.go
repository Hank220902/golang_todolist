package controllers

import (
	"fmt"
	"todolist/service"

	"github.com/kataras/iris/v12"
)

func Register(ctx iris.Context){
	result :=service.Register(ctx)
	fmt.Println(result)
	if result ==1{
		ctx.JSON("success")
	}else{
		ctx.JSON("emailExists")
	}

}

func Login(ctx iris.Context){
	
	result := service.Login(ctx)
	fmt.Println(result)
	ctx.JSON(result)
}

func Logout(ctx iris.Context){
	result := service.Logout(ctx)
	if result ==1{
		ctx.JSON("success")
	}else{
		ctx.JSON("error")
	}
}