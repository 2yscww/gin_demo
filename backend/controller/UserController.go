package controller

import (
	"fmt"
	"log"
	"net/http"

	"gin_demo/common"
	"gin_demo/dto"
	"gin_demo/model"
	"gin_demo/response"
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

	var requestUser = model.User{}

	c.Bind(&requestUser)

	//获取参数
	// name := c.PostForm("username")
	// telephone := c.PostForm("telephone")
	// password := c.PostForm("password")

	name := requestUser.Username
	telephone := requestUser.Telephone
	password := requestUser.Password

	fmt.Println("用户名为:", name)
	fmt.Println("电话号码为:", telephone)
	fmt.Println("密码为:", password)

	// ! 后端无法正常获取到数据
	// TODO  修复后端无法正常获取到数据的问题
	// 引入数据库实例
	db := common.GetDB()

	// 验证手机号长度
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		// c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}

	// 验证密码长度
	if len(password) < 8 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于8位")
		// c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于8位"})
		return
	}

	// 判断姓名是否为空
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println("name:", name, "telephone:", telephone, "password:", password)

	//判断手机号是否存在

	if isTelephoneExist(db, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "该手机号已注册用户")
		// c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "该手机号已注册用户"})
		return
	}

	//创建用户

	// 密码不能明文保存,对密码加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {

		response.Response(c, http.StatusUnprocessableEntity, 500, nil, "系统错误")
		// c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500, "msg": "加密错误"})
	}

	newUser := model.User{
		Username:  name,
		Telephone: telephone,
		Password:  string(hashPassword),
	}

	db.Create(&newUser)

	// 返回结果
	// c.JSON(200, gin.H{
	// 	"code": "200",
	// 	"msg":  "注册成功"})

	response.Success(c, nil, "注册成功")
}

func Login(c *gin.Context) {

	db := common.GetDB()
	// 获取参数

	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	// 数据验证

	// 验证手机号长度
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		// c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}

	// 验证密码长度
	if len(password) < 8 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于8位")
		// c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于8位"})
		return
	}

	// 判断手机号是否存在

	var user model.User
	db.Where("telephone = ?", telephone).First(&user)

	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		// c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 400, nil, "密码错误")
		// c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 400, "msg": "密码错误"})
		return
	}

	// 发放token

	token, err := common.ReleaseToken(user)

	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 500, nil, "系统异常")
		// c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error: %v", err)
		return
	}

	// 返回结果
	// c.JSON(200, gin.H{
	// 	"code": "200",
	// 	"data": gin.H{"token": token},
	// 	"msg":  "登录成功"})

	response.Success(c, gin.H{"token": token}, "登录成功")
}

// 返回账户信息

func Info(c *gin.Context) {
	user, _ := c.Get("user")

	// ? 使用DTO来封装传输 来确保不会有不相关的敏感信息也被返回
	c.JSON(200, gin.H{
		"code": "200",
		"data": gin.H{"user": dto.ToUserDto(user.(model.User))}})

}
