package main

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

type LoginForm struct {
	UserName string `json:"username" binding:"required,min=3,max=7"`
	PassWord string `json:"password" binding:"required,len=8"`
}

type RegisterForm struct {
	UserName string `json:"username" binding:"required,min=3,max=7"`
	PassWord string `json:"password" binding:"required,len=8"`
	Age      uint32 `json:"age" binding:"required,gte=1,lte=150"`
	Gender   string `json:"gender" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func main() {

	if err := InitializeTranslate(); err != nil {
		fmt.Println(err)
	}

	r := gin.Default()

	r.POST("login", loginHander)
	r.POST("register", registerHander)

	r.Run()
}

func registerHander(c *gin.Context) {
	var rf RegisterForm
	if err := c.ShouldBindJSON(&rf); err != nil {
		err, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(200, gin.H{
				"Code": 40010,
				"Msg":  "注册失败，请确认",
				"Err":  err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"Code": 40004,
			"Msg":  err.Translate(trans),
		})
		return
	}

	c.JSON(200, gin.H{
		"Code": 0,
		"Msg":  "注册成功",
		"data": rf,
	})
}

func loginHander(c *gin.Context) {
	var lg LoginForm
	if err := c.ShouldBindJSON(&lg); err != nil {
		c.JSON(200, gin.H{
			"Code": 40001,
			"Msg":  "登录失败，请确认",
			"Err":  err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"Code": 0,
		"Msg":  "登录成功",
		"data": lg.UserName,
	})
}

//locale string
func InitializeTranslate() (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("json")
			return name
		})

		zh := zh.New()
		uni := ut.New(zh, zh)
		trans, _ = uni.GetTranslator("zh")

		err = zh_translations.RegisterDefaultTranslations(v, trans)
		return

	}
	return
}
