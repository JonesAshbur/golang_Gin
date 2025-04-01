package main

import (
	"golang_Gin/ginCase5/models"
	"golang_Gin/ginCase5/routers"
	"html/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	//Default默认使用了Logger，Recover中间件，如果不想使用可以换做gin.New()
	r := gin.Default()
	//配置session中间件，创建一个基于cookie的存储引擎，secret是加密密钥
	store := cookie.NewStore([]byte("secret"))
	//store是存储引擎，可以替换
	r.Use(sessions.Sessions("mysession", store))
	//在中间件或者handler中启动新的goroutine时，不能使用原始的上下文（c *gin.Context）,必须使用其只读副本c.Copy()
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
	})
	r.LoadHTMLGlob("views/*/*.html")
	r.Static("/static", "./static")
	routers.DefaultRoutersInit(r)
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)

	//设置为80端口访问时直接访问localhost，无需加端口号
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
