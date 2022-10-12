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
