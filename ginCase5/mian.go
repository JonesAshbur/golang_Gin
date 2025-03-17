package main

import (
	"golang_Gin/ginCase5/models"
	"golang_Gin/ginCase5/routers"
	"html/template"

	"github.com/gin-gonic/gin"
)

func main() {
	//Default默认使用了Logger，Recover中间件，如果不想使用可以换做gin.New()
	r := gin.Default()

	//在中间件或者handler中启动新的goroutine时，不能使用原始的上下文（c *gin.Context）,必须使用其只读副本c.Copy()
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
	})
	r.LoadHTMLGlob("views/*/*.html")
	r.Static("/static", "./static")
	routers.DefaultRoutersInit(r)
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)

	//80端口访问时直接访问localhost，无需加端口号
	err := r.Run(":80")
	if err != nil {
		return
	}
}
