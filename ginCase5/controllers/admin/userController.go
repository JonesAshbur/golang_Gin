package admin

import "github.com/gin-gonic/gin"

type UserController struct {
	BaseController
}

func (con *UserController) Index(c *gin.Context) {
	c.String(200, "user list")
	con.Success(c)
}

func (con *UserController) Add(c *gin.Context) {
	c.String(200, "add user")
}

func (con *UserController) Edit(c *gin.Context) {
	c.String(200, "edit user")
}
