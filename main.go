package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type myType map[string]interface{}

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello gin demo",
	})
}
func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}
func hi(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hi",
	})
}

func mapTest(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "自己写的map，没有使用gin.H",
		"code":    200,
	})
}

func mapTest1(c *gin.Context) {
	c.JSON(http.StatusOK, myType{
		"message": "自己写的map",
		"code":    200,
		"date":    "",
	})
}

func main() {
	//创建gin的默认engine实例
	r := gin.Default()

	/*	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello gin api",
		})
	})*/

	r.GET("/test", test)
	r.GET("/hello", hello)
	r.GET("/hi", hi)
	r.GET("/mapTest", mapTest)
	r.GET("/mapTest1", mapTest1)
	//如果需要修改端口 r.Run(":8888")
	r.Run()
}
