package home

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DefaultController struct{}

func (con *DefaultController) Index(c *gin.Context) {
	c.SetCookie("username", "jones", 3600, "/", "localhost", false, false)
	c.HTML(http.StatusOK, "default/index.html", gin.H{
		"msg":  "msg信息",
		"time": 1629788418,
	})
}
func (con *DefaultController) DeleteCookie(c *gin.Context) {
	//删除cookie将maxAge设置为小于零的数
	c.SetCookie("username", "jones", -1, "/", "localhost", false, false)
	c.String(http.StatusOK, "删除成功")
}
func (con *DefaultController) News(c *gin.Context) {
	cookie, err := c.Cookie("username")
	if err != nil {
		return
	}
	fmt.Println("设置在index中的cookie:", cookie)
	c.String(200, "%v", "新闻页面")
}
