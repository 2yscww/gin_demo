package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

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

func main() {
	// 1.创建路由
	r := gin.Default()

	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
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

		// 姓名为空
		if len(name) == 0 {
			name = RandomString(10)
		}

		log.Println("name:", name, "telephone:", telephone, "password:", password)

		//判断手机号是否存在

		//创建用户

		// 返回结果
		c.JSON(200, gin.H{"msg": "注册成功"})
	})

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	// panic(r.Run(":8081"))
	r.Run(":8081")
}
