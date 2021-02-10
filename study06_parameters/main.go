package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Id   int    `uri:"id"`
	Name string `uri:"name"`
}

func main() {
	r := gin.Default()

	/*	r.GET("/posts", func(c *gin.Context) {
			c.String(200, "post")
		})

		r.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(http.StatusOK, gin.H{
				"id": id,
			})
		})*/

	r.GET("/:id/:name", func(c *gin.Context) {

		var p Person
		if err := c.ShouldBindUri(&p); err != nil {
			c.Status(404)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"name": p.Name,
			"id":   p.Id,
		})
	})

	/*	r.DELETE("/posts/:id", func(c *gin.Context) {

		})*/

	r.Run()
}
