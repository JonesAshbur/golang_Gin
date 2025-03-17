package routers

import (
	"github.com/gin-gonic/gin"
	"golang_Gin/ginCase5/controllers/home"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/default")
	{
		defaultRouters.GET("/", (&home.DefaultController{}).Index)
		defaultRouters.GET("/news", (&home.DefaultController{}).News)
		defaultRouters.GET("/deletecookie", (&home.DefaultController{}).DeleteCookie)
	}
}
