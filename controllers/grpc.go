package controllers

import (
	"fmt"
	"todolist/service"
	"github.com/kataras/iris/v12"
)

func GrpcCreateToDoList(ctx iris.Context) {
	result:= service.GCreateToDoList(ctx)
	fmt.Println(result)
	if result == 1 {
		ctx.JSON("新增成功")
	} else if result == 3 {
		ctx.JSON("tokenError")
	} else {
		ctx.JSON("error")
	}
}

func GrpcGetFilterToDoList(ctx iris.Context) {
	result,num := service.GGetFilterTodoList(ctx)
	fmt.Println(result)
	if num == 1{
		ctx.JSON(result.GetResult)
	}else if num == 3{
		ctx.JSON("tokenError")
	}else{
		ctx.JSON("error")
	}
}

func GrpcGetAllToDoList(ctx iris.Context){
	result,num := service.GGetAllToDoList(ctx)
	fmt.Println(result)
	if num == 1{
		ctx.JSON(result.GetResult)
	}else if num == 3{
		ctx.JSON("tokenError")
	}else{
		ctx.JSON("error")
	}
}

func GrpcUpdateToDoList(ctx iris.Context) {
	result := service.GUpdateToDoList(ctx)
	if result ==1{
		ctx.JSON("success")
	}else if result ==3{
		ctx.JSON("tokenError")
	}else{
		ctx.JSON("error")
	}
}

func GrpcDeleteToDoList(ctx iris.Context) {
	result := service.GDeleteToDoList(ctx)
	if result ==1{
		ctx.JSON("success")
	}else if result ==3{
		ctx.JSON("tokenError")
	}else{
		ctx.JSON("error")
	}

}
