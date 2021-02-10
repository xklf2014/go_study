package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, "get")
	})

	r.POST("/post", func(c *gin.Context) {
		c.String(http.StatusOK, "post")
	})

	r.PUT("/put/:id", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("put id: %s", c.Param("id")))
	})

	r.DELETE("/delete/:id", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("delete id: %s", c.Param("id")))
	})

	r.PATCH("/patch/:id", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("patch id: %s", c.Param("id")))
	})

	r.Any("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "match any methods",
			"code":    200,
			"data":    "any data",
		})
	})

	r.Run()
}
