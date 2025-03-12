package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/default")
	{
		defaultRouters.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "default/index.html", gin.H{
				"msg": "msg信息",
			})
		})
		defaultRouters.GET("/news", func(c *gin.Context) {
			c.String(200, "%v", "新闻页面")
		})
	}
}
