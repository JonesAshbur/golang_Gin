package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

type Article1 struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
type Article2 struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}

func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}
func main() {
	res := gin.Default()
	//自定义模板函数，放在加载模板前
	res.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})
	res.LoadHTMLGlob("views/**/*")
	res.Static("/static", "./static")
	//get请求传值
	res.GET("/", func(c *gin.Context) {
		username := c.Query("username")
		age := c.Query("age")
		//无参数默认值
		sex := c.DefaultQuery("sex", "none")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"sex":      sex,
		})
	})
	res.GET("/article", func(c *gin.Context) {
		id := c.DefaultQuery("id", "1")
		c.JSON(http.StatusOK, gin.H{
			"id":  id,
			"msg": "news detail",
		})
	})
	//post请求传值
	res.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/user.html", gin.H{})
	})
	//获取表达post数据
	res.POST("/doAddUser1", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		age := c.DefaultPostForm("age", "18")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})
	//default
	res.GET("/default", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title":    "默认首页",
			"score":    89,
			"hobby":    []string{"eat", "drink", "play"},
			"testList": []string{},
			"list": &Article1{
				Title:   "新闻标题",
				Content: "新闻内容",
			},
			"date": 1741486256,
		})
	})
	res.GET("/default/news", func(c *gin.Context) {
		news := &Article1{
			Title:   "default新闻",
			Content: "default新闻内容",
		}
		c.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "default news page",
			"news":  news,
		})
	})
	//获取get，post传递的数据绑定到结构体
	res.GET("/getUserInfo", func(c *gin.Context) {
		user := &UserInfo{}
		err := c.ShouldBind(&user)
		if err == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
		fmt.Printf("%#v\n", user)
	})
	res.POST("/doAddUser2", func(c *gin.Context) {
		user := &UserInfo{}
		err := c.ShouldBind(&user)
		if err == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
		fmt.Printf("%#v\n", user)
	})
	//获取post xml数据
	res.POST("/xml", func(c *gin.Context) {
		Artilce2 := &Article2{}
		xmlSliceData, _ := c.GetRawData() //返回byte类型切片
		fmt.Println(string(xmlSliceData))
		if err := xml.Unmarshal(xmlSliceData, &Artilce2); err == nil {
			c.JSON(http.StatusOK, Artilce2)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
	})
	//动态路由传值
	res.GET("/list/:cid", func(c *gin.Context) {
		cid := c.Param("cid")
		c.String(http.StatusOK, "%v", cid)
	})
	//admin
	res.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "admin首页",
		})
	})
	res.GET("/admin/news", func(c *gin.Context) {
		news := &Article1{
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
