package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("ttt", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
		})
	})

	r.Run()
}
