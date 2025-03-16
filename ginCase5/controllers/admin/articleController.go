package admin

import "github.com/gin-gonic/gin"

type ArticleController struct {
	BaseController
}

func (con *ArticleController) Index(c *gin.Context) {
	con.Success(c)
	c.String(200, "article list")
}

func (con *ArticleController) Add(c *gin.Context) {
	c.String(200, "add article")
}

func (con *ArticleController) Edit(c *gin.Context) {
	c.String(200, "edit article")
}
