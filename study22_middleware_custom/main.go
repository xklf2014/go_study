package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.Use(middleware1)
	//r.Use(middleware2())

	r.Use(RefererMiddleware())

	r.Use(func(c *gin.Context) {
		fmt.Println("我是另外一个中间件")
	})

	r.GET("test", func(c *gin.Context) {

	})

	r.Run()
}

/*func middleware1(c *gin.Context) {

}

func middleware2() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}*/

func RefererMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ref := c.GetHeader("Referer")
		if ref == "" {
			c.AbortWithStatusJSON(200, gin.H{
				"msg": "非法访问",
			})
			return
		}
		c.Next()
	}
}
