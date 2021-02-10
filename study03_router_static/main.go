package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//加载静态图片
	r.Static("/images", "./images")
	r.StaticFS("/static", http.Dir("./static"))
	r.StaticFile("index", "index.html")
	r.Run()
}
