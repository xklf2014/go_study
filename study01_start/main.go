package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello gin",
			"code":    200,
			"data":    "my data",
		})
	})

	r.Run(":9090")
}
