package middleWares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleWare(c *gin.Context) {
	//	业务逻辑...
	fmt.Println("当前时间：", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("请求路径：", c.Request.URL)

	c.Set("username", "jones")

	//	定义一个go Routine统计日志
	cCp := c.Copy()
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Done in Path" + cCp.Request.URL.Path)
	}()
}
