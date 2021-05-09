package main

import (
	v1 "gin_demo/study31_login_project/api/v1"
	"gin_demo/study31_login_project/model"

	"github.com/gin-gonic/gin"
)

func main() {

	err := model.InitializeDatabase()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	v := r.Group("api/v1")
	{
		v.POST("user/register", v1.UserRegisterHandler)

		v.POST("user/login", v1.UserLoginHandler)
	}
	r.Run()
}
