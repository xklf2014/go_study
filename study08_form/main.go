package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/form_post", postHandler)
	r.POST("/form_post_array", postArrayHandler)
	r.POST("/form_post_map", postMapHandler)
	r.Run()
}

func postMapHandler(c *gin.Context) {
	p := c.PostFormMap("person")
	c.JSON(200, gin.H{
		"datas": p,
	})
}

func postArrayHandler(c *gin.Context) {
	datas := c.PostFormArray("datas")
	c.JSON(200, gin.H{
		"datas": datas,
	})
}

func postHandler(c *gin.Context) {
	message := c.PostForm("message")
	name := c.DefaultPostForm("name", "story")
	c.JSON(200, gin.H{
		"message": message,
		"name":    name,
	})
}
