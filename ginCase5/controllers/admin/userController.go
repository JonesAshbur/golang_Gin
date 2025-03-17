package admin

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
)

type UserController struct {
	BaseController
}

func (con *UserController) Index(c *gin.Context) {
	c.String(200, "user list")
	con.Success(c)
}

func (con *UserController) Add(c *gin.Context) {
	//c.String(200, "add user")
	c.HTML(200, "admin/userAdd.html", gin.H{})
}

func (con *UserController) Edit(c *gin.Context) {
	c.String(200, "edit user")
}

// 单文件上传
func (con *UserController) DoUpload(c *gin.Context) {
	username := c.PostForm("username")
	file, err := c.FormFile("face")

	//获取文件名称file.Filename
	dst := path.Join("./static/upload", file.Filename)
	if err == nil {
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			return
		}
	}
	log.Println(username, file.Filename)
	c.JSONP(http.StatusOK, gin.H{
		"success":  true,
		"username": username,
		"dst":      dst,
	})
}

// 多文件上传
func (con *UserController) DoMultipleUpload(c *gin.Context) {
	username := c.PostForm("username")

	form, _ := c.MultipartForm()
	files, _ := form.File["face[]"]

	for _, file := range files {
		dst := path.Join("./static/upload", file.Filename)
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			return
		}
	}
	c.JSONP(http.StatusOK, gin.H{
		"success":  true,
		"username": username,
	})
}
