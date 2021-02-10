package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware1, middleware2)
	r.GET("test", func(c *gin.Context) {
		log.Println("func start...")
		k := c.GetInt("key")
		c.Set("key", k+1000)
		log.Println("func end...")

	})

	r.Run()
}

func middleware1(c *gin.Context) {
	log.Println("middleware1 in ...")
	c.Set("key", 500)
	log.Println("middleware1 before next ...")
	c.Next()
	log.Println("middleware1 after next ...")
	log.Println("middleware1 out ...")

	c.JSON(200, gin.H{
		"msg": c.GetInt("key"),
	})
}

func middleware2(c *gin.Context) {
	log.Println("middleware2 in ...")
	log.Println("middleware2 before next ...")
	c.Next()
	log.Println("middleware2 after next ...")
	log.Println("middleware2 out ...")

}
