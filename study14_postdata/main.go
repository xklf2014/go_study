package main

import "github.com/gin-gonic/gin"

type User struct {
	ID string `form:"id" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("bindquery", func(c *gin.Context) {
		var user User
		// get : form
		// post : Content-Type : json / xml / form-data
		if err := c.BindQuery(&user); err != nil {
			c.JSON(200, gin.H{
				"Code": 200,
				"Msg":  err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"Code": 0,
			"Id":   user.ID,
		})
	})

	r.Run()
}
