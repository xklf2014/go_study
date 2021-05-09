package v1

import (
	"gin_demo/study31_login_project/service"

	"github.com/gin-gonic/gin"
)

func UserLoginHandler(c *gin.Context) {
	var s service.UserLoginService
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(200, gin.H{
			"msg": "login err",
		})
	} else {
		res := s.Login()
		c.JSON(200, res)
	}
}

func UserRegisterHandler(c *gin.Context) {
	var s service.UserRegisterService
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(200, gin.H{
			"msg": "register err",
		})
	} else {
		res := s.Register()
		c.JSON(200, res)
	}
}
