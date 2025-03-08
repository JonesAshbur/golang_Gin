package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type userList struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	res := gin.Default()
	//配置模板文件
	res.LoadHTMLGlob("views/*")
	res.GET("/", func(c *gin.Context) {
		c.String(200, "value：%v", "Hello Gin")
	})
	res.GET("/json1", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"success": true,
			"msg":     "json1 content",
		})
	})
	res.GET("/json2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"msg":     "json2 content",
		})
	})
	res.GET("/json3", func(c *gin.Context) {
		user1 := userList{
			Title:   "title-1",
			Desc:    "desc-1",
			Content: "content-1",
		}
		c.JSON(200, user1)
	})
	res.GET("/jsonp", func(c *gin.Context) {
		user := userList{
			Title:   "jsonp",
			Desc:    "jsonp",
			Content: "jsonp-content",
		}
		c.JSONP(200, user)
	})
	res.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"success": true,
			"msg":     "xml content",
		})
	})
	res.GET("/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "后台news数据",
		})
	})
	res.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"title": "后台goods数据",
		})
	})
	err := res.Run(":8000")
	if err != nil {
		return
	}
}
