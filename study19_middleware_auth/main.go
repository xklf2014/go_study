package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//全局调用
	//r.Use(loginAuth)

	//single作用域
	r.GET("test", singleAuth, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
		})
	})

	r.POST("login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "token",
		})
	})

	r.POST("register", func(c *gin.Context) {

	})

	//组作用域
	user := r.Group("user", loginAuth)
	{
		user.POST(":id", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "登录保护",
			})
		})

		user.PUT(":id", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "修改登录保护",
			})
		})
	}

	r.Run()

}

func loginAuth(c *gin.Context) {
	fmt.Println("我是登录中间件")
}

func singleAuth(c *gin.Context) {
	fmt.Println("我是single作用域")
}
