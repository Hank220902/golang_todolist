package controllers

import (
	"fmt"
	"todolist/service"

	"github.com/kataras/iris/v12"
)

func Register(ctx iris.Context){
	
	fmt.Println(service.Register(ctx))
}

func Login(ctx iris.Context){
	service.Login()
	fmt.Println("login")
}