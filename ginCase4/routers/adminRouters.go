package routers

import "github.com/gin-gonic/gin"

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", func(c *gin.Context) {
			c.String(200, "%v", "后台首页")
		})
		adminRouters.GET("/add", func(c *gin.Context) {
			c.String(200, "%v", "增加用户")
		})
		adminRouters.GET("/edit", func(c *gin.Context) {
			c.String(200, "%v", "修改用户")
		})
		adminRouters.GET("/article", func(c *gin.Context) {
			c.String(200, "%v", "admin新闻列表")
		})
	}
}
