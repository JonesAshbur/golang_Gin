package api

import "github.com/gin-gonic/gin"

type ApiController struct{}

func (con *ApiController) Index(c *gin.Context) {
	c.String(200, "%v", "api接口")
}
func (con *ApiController) UserList(c *gin.Context) {
	c.String(200, "%v", "用户列表接口")
}
