package controller

import (
	"log"
	"net/http"

	"gin_demo/common"
	"gin_demo/model"
	"gin_demo/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 验证手机号是否存在
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func Register(c *gin.Context) {
	// c.String(http.StatusOK, "hello World!")
	//获取参数

	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	// 引入数据库实例
	db := common.GetDB()

	// 验证手机号长度
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}

	// 验证密码长度
	if len(password) < 8 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于8位"})
		return
	}

	// 判断姓名是否为空
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println("name:", name, "telephone:", telephone, "password:", password)

	//判断手机号是否存在

	if isTelephoneExist(db, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "该手机号已注册用户"})
		return
	}

	//创建用户
	newUser := model.User{
		Username:  name,
		Telephone: telephone,
		Password:  password,
	}

	db.Create(&newUser)

	// 返回结果
	c.JSON(200, gin.H{"msg": "注册成功"})
}
