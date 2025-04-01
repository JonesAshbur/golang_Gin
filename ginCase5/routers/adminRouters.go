package routers

import (
	"golang_Gin/ginCase5/controllers/admin"
	"golang_Gin/ginCase5/middleWares"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin", middleWares.InitMiddleWare)
	//使用中间件方法2
	//adminRouters.Use(middleWares.InitMiddleWare)
	{
		adminRouters.GET("/", (&admin.IndexController{}).Index)
		adminRouters.GET("/user", (&admin.UserController{}).Index)
		adminRouters.GET("/user/add", (&admin.UserController{}).Add)
		adminRouters.GET("/user/edit", (&admin.UserController{}).Edit)
		adminRouters.GET("/user/delete", (&admin.UserController{}).Delete)
		adminRouters.POST("/user/doUpload", (&admin.UserController{}).DoUpload)
		adminRouters.POST("/user/doMultipleUpload", (&admin.UserController{}).DoMultipleUpload)
		adminRouters.GET("/article", (&admin.ArticleController{}).Index)
		adminRouters.GET("/article/add", (&admin.ArticleController{}).Add)
		adminRouters.GET("/article/edit", (&admin.ArticleController{}).Edit)
	}
}
