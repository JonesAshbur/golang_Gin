package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建默认路由引擎
	res := gin.Default()
	//配置路由
	res.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "%v", "hello gin")
	})
	res.GET("/home", func(c *gin.Context) {
		c.String(http.StatusOK, "%v", "home page")
	})
	res.POST("/add", func(c *gin.Context) {
		c.String(http.StatusOK, "%v", "post请求用来增加数据")
	})
	res.PUT("/edit", func(c *gin.Context) {
		c.String(http.StatusOK, "%v", "PUT请求用来编辑数据")
	})
	res.DELETE("/del", func(c *gin.Context) {
		c.String(http.StatusOK, "%v", "DELETE请求用来删除数据")
	})
	//启动一个web服务
	err := res.Run(":8000")
	if err != nil {
		return
	}
}
