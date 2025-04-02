package admin

import (
	"fmt"
	"golang_Gin/ginCase5/models"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	// 指定条件查询
	userList = []models.User{}
	models.DB.Where("age > ? AND age < ?", 18, 23).Find(&userList)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    userList,
	})

	// 指定条件查询id在1-3之间的
	userList = []models.User{}
	models.DB.Where("id BETWEEN ? AND ?", 1, 3).Find(&userList)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    userList,
	})

	// 指定条件查询 id在1，3，5中的数据
	userList = []models.User{}
	models.DB.Where("id IN (?)", []int{1, 3, 5}).Find(&userList)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    userList,
	})

	// 模糊查询
	userList = []models.User{}
	models.DB.Where("username LIKE ?", "%ashbur").Find(&userList)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    userList,
	})

	// 查询id=1或2的数据
	userList = []models.User{}
	models.DB.Where("id = ? OR id = ?", 1, 2).Find(&userList)
	// models.DB.Where("id = ? ", 1).Or("id = ?", 2).Find(&userList)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    userList,
	})

	// 使用slect返回指定字段
	userList = []models.User{}
	models.DB.Select("id, username, age").Find(&userList)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    userList,
	})

	// order排序
	userList = []models.User{}
	// 按年龄降序，id升序,取前10条
	models.DB.Order("age desc").Order("id asc").Limit(10).Find(&userList)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    userList,
	})

	// offset分页
	userList = []models.User{}
	// 从第10条开始，取10条
	models.DB.Offset(10).Limit(10).Find(&userList)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    userList,
	})

	// count统计
	var count int64
	models.DB.Where("age > 18").Count(&count)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"count":   count,
	})

	// 原生sql
	var result []models.User
	models.DB.Raw("SELECT * FROM user WHERE age > ?", 18).Scan(&result)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result,
	})

	// 增删改查使用Exec（执行SQL），获取数据使用Raw结合Scan

	// 原生sql统计数量
	var count2 int64
	models.DB.Raw("SELECT COUNT(*) FROM user WHERE age > ?", 18).Scan(&count2)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"count":   count2,
	})

	// benlogs to关联查询,一对一查询
	var result2 []models.User
	models.DB.Preload("User").Find(&result2)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result2,
	})

	// 如果不知道外键，可以使用gorm重写外键： `gorm:"foreignKey:CompanyRefer"`

	// 一对多查询，Has many

	// 多对多查询 many to many

	// 预加载SQL，设置preload第二个参数为func
	models.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Order("id desc")
	}).Find(&result2)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result2,
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

	// 原生sql删除
	models.DB.Exec("DELETE FROM user WHERE username = ?", "new user name")
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
