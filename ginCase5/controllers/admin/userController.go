package admin

import (
	"fmt"
	"golang_Gin/ginCase5/models"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (con *UserController) Index(c *gin.Context) {
	userList := []models.User{}
	models.DB.Where("age > 18").Find(&userList)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    userList,
	})
}

func (con *UserController) Add(c *gin.Context) {
	//c.String(200, "add user")
	user := models.User{
		UserName: "ashbur",
		Age:      23,
		Email:    "ashbur@163.com",
		AddTime:  int(time.Now().Unix()),
	}
	err := models.DB.Create(&user).Error
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("user:", user)
	c.JSON(http.StatusOK, gin.H{
		"message": "update user success",
		"data":    user,
	})
}

func (con *UserController) Edit(c *gin.Context) {
	// 更新用户所有信息
	// user := models.User{Id: 3}
	// models.DB.Find(&user)
	// fmt.Println("user id=3:", user)
	// user.UserName = "new user name"
	// user.Age = 00
	// user.Email = "new email"
	// user.AddTime = int(time.Now().Unix())
	// err := models.DB.Save(&user).Error
	// if err != nil {
	// 	log.Panic(err)
	// }

	// 更新用户部分信息(单例，多列)

	// 直接修改
	// user := models.User{}
	// models.DB.Model(&user).Where("id = ?", 2).Update("username", "new name")

	// 获取，修改，保存
	user := models.User{}
	models.DB.Where("id = ?", 2).Find(&user)
	user.UserName = "name"
	models.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "update user info success",
		"data":    user,
	})
}

func (con *UserController) Delete(c *gin.Context) {
	// 删除一条数据
	user := models.User{}
	models.DB.Where("id = ?", 2).Delete(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "delete user success",
	})

	// 指定条件删除数据
	models.DB.Where("username = ?", "new user name").Delete(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "delete user success",
	})
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
	files := form.File["face[]"]

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
