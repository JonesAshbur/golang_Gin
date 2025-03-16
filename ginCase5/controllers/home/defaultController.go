package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DefaultController struct{}

func (con *DefaultController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "default/index.html", gin.H{
		"msg":  "msg信息",
		"time": 1629788418,
	})
}
func (con *DefaultController) News(c *gin.Context) {
	c.String(200, "%v", "新闻页面")
}
