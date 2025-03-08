package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	res := gin.Default()
	res.LoadHTMLGlob("views/**/*")
	//default
	res.GET("/default", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title":    "默认首页",
			"score":    89,
			"hobby":    []string{"eat", "drink", "play"},
			"testList": []string{},
			"list": &Article{
				Title:   "新闻标题",
				Content: "新闻内容",
			},
		})
	})
	res.GET("/default/news", func(c *gin.Context) {
		news := &Article{
			Title:   "default新闻",
			Content: "default新闻内容",
		}
		c.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "default news page",
			"news":  news,
		})
	})
	//admin
	res.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "admin首页",
		})
	})
	res.GET("/admin/news", func(c *gin.Context) {
		news := &Article{
			Title:   "admin新闻",
			Content: "admin新闻内容",
		}
		c.HTML(http.StatusOK, "admin/news.html", gin.H{
			"title": "admin news page",
			"news":  news,
		})
	})
	err := res.Run(":8000")
	if err != nil {
		return
	}
}
