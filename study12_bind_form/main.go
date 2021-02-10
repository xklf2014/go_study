package main

import "github.com/gin-gonic/gin"

type User struct {
	ID       string `form:"id" binding:"required"`
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	r.POST("user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(200, gin.H{
				"Code": 200,
				"Msg":  err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"Code":     0,
			"ID":       user.ID,
			"Username": user.Username,
			"Password": user.Password,
		})

	})

	r.Run()
}
