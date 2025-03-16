package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_Gin/ginCase4/controllers/admin"
)

// 单一路由中间件
func initMiddleWareCase(c *gin.Context) {
	fmt.Println("中间件信息-1")

	//	调用该请求的剩余处理程序
	c.Next()

	//	终止调用该请求的剩余处理程序
	//c.Abort()

	//	剩余处理程序
	fmt.Println("中间件信息-2")

}

func initMiddleWare1(c *gin.Context) {
	fmt.Println("中间件-1-信息1")
	c.Next()
	fmt.Println("中间件-1-信息2")
}

func initMiddleWare2(c *gin.Context) {
	fmt.Println("中间件-2-信息1")
	c.Next()
	fmt.Println("中间件-2-信息2")
}
func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", initMiddleWareCase, func(c *gin.Context) {
			fmt.Println("This is home page")
		}, (&admin.IndexController{}).Index)
		adminRouters.GET("/user", initMiddleWare1, initMiddleWare2, func(c *gin.Context) {
			fmt.Println("This is user page")
		}, (&admin.UserController{}).Index)
		adminRouters.GET("/user/add", (&admin.UserController{}).Add)
		adminRouters.GET("/user/edit", (&admin.UserController{}).Edit)
		adminRouters.GET("/article", (&admin.ArticleController{}).Index)
		adminRouters.GET("/article/add", (&admin.ArticleController{}).Add)
		adminRouters.GET("/article/edit", (&admin.ArticleController{}).Edit)
	}
}
