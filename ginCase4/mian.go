package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_Gin/ginCase4/routers"
)

// 全局路由中间件
func initMiddleWare(*gin.Context) {
	fmt.Println("这是全局路由中间件")
}
func main() {
	r := gin.Default()

	r.LoadHTMLGlob("views/**/*")
	r.Static("/static", "./static")
	//全局中间件
	r.Use(initMiddleWare)
	routers.DefaultRoutersInit(r)
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)

	err := r.Run(":8000")
	if err != nil {
		return
	}
}
