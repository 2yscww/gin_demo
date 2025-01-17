package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义数据库全局变量
var db *gorm.DB

// 连接mysql数据库
func InitDB() {
	username := "root"
	passwd := "w86#qNwV"

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local", username, passwd)
	var err error

	// 初始化数据库连接
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&User{})
}

// 验证手机号是否存在
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

// 随机分配字符串
func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiop")
	result := make([]byte, n)

	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

type User struct {
	gorm.Model

	Username  string `"gorm:'type:varchar(20);not null'"`
	Telephone string `"gorm:'type:varchar(11);not null;unique'"`
	Password  string `"gorm:'type:size:255;not null'"`
}

func main() {
	// 1.创建路由
	r := gin.Default()

	InitDB() // 初始化数据库
	log.Println("Connected to database:", db)

	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response

	// api注册
	r.POST("/api/auth/register", func(c *gin.Context) {
		// c.String(http.StatusOK, "hello World!")
		//获取参数

		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")

		//数据验证

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
			name = RandomString(10)
		}

		log.Println("name:", name, "telephone:", telephone, "password:", password)

		//判断手机号是否存在

		if isTelephoneExist(db, telephone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "该手机号已注册用户"})
			return
		}

		//创建用户
		newUser := User{
			Username:  name,
			Telephone: telephone,
			Password:  password,
		}

		db.Create(&newUser)

		// 返回结果
		c.JSON(200, gin.H{"msg": "注册成功"})
	})

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	// panic(r.Run(":8081"))
	r.Run(":8081")
}
