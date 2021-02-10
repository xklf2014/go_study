package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/gets", GetHandler)
	r.POST("/post", PostHandler)
	r.DELETE("/delete/1", DeleteHandler)
	//路由分组
	p := r.Group("/posts")
	{
		p.GET("/", GetHandler)
		p.POST("/", PostHandler)
		p.DELETE("/1", DeleteHandler)
	}

	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	v2 := r.Group("/api/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	api := r.Group("/api")
	{
		v3 := api.Group("/v3")
		{
			v3.POST("/login", loginEndpoint)
			v3.POST("/submit", submitEndpoint)
			v3.POST("/read", readEndpoint)
		}
	}

	r.Run()
}

func readEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "read",
	})
}

func submitEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "submit",
	})
}

func loginEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "login",
	})
}

func GetHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get",
	})
}

func PostHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "post",
	})
}

func DeleteHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "delete",
	})
}
