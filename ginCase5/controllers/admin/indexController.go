package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (con *IndexController) Index(c *gin.Context) {
	c.String(200, "user list\n")
	username, _ := c.Get("username")
	fmt.Println(username)
	value, _ := username.(string)
	c.String(200, "username="+value)
}
