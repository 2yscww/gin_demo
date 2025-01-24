package controller

import (
	"log"
	"net/http"

	"gin_demo/common"
	"gin_demo/model"
	"gin_demo/util"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	// 密码不能明文保存,对密码加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		// TODO 稍后要把这段反馈修改为其他的理由
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500, "msg": "加密错误"})
	}

	newUser := model.User{
		Username:  name,
		Telephone: telephone,
		Password:  string(hashPassword),
	}

	db.Create(&newUser)

	// 返回结果
	c.JSON(200, gin.H{
		"code": "200",
		"msg":  "注册成功"})
}

func Login(c *gin.Context) {

	db := common.GetDB()
	// 获取参数

	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	// 数据验证

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

	// 判断手机号是否存在

	var user model.User
	db.Where("telephone = ?", telephone).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 400, "msg": "密码错误"})
		return
	}

	// 发放token

	token, err := common.ReleaseToken(user)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error: %v", err)
		return
	}

	// 返回结果
	c.JSON(200, gin.H{
		"code": "200",
		"data": gin.H{"token": token},
		"msg":  "登录成功"})
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(200, gin.H{
		"code": "200",
		"data": gin.H{"user": user}})

}
