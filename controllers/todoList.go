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
	if result ==1{
		ctx.JSON("新增成功")
	}else if result ==3{
		ctx.JSON("tokenError")
	}else{
		ctx.JSON("error")
	}


}



func GetToDoList(ctx iris.Context) {
	result,num := service.GetOneToDoList(ctx)
	fmt.Println(result)
	if num == 1{
		ctx.JSON(result)
	}else if num == 3{
		ctx.JSON("tokenError")
	}else{
		ctx.JSON("error")
	}
	
}

func GetManyToDoList(ctx iris.Context) {
	result,num := service.GetManyToDoList(ctx)
	if num == 1{
		ctx.JSON(result)
	}else if num==3{
		ctx.JSON("tokenError")
	}else{
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
	if result ==1{
		ctx.JSON("刪除成功")
	}else if result ==3{
		ctx.JSON("tokenError")
	}else{
		ctx.JSON("error")
	}

}
