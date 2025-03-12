package routers

import "github.com/gin-gonic/gin"

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", func(c *gin.Context) {
			c.String(200, "%v", "api接口")
		})
		apiRouters.GET("/userlist", func(c *gin.Context) {
			c.String(200, "%v", "用户列表接口")
		})

	}
}
