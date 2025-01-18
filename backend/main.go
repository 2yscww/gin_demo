package main

import (
	"log"

	"gin_demo/common"
	"gin_demo/router"

	"github.com/gin-gonic/gin"
)

// 定义数据库全局变量
// var db *gorm.DB

func main() {

	common.InitDB()
	db := common.GetDB() // 初始化数据库

	log.Println("Connected to database:", db)

	// 1.创建路由
	r := gin.Default()

	r = router.CollectRtoutes(r)

	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response

	// api注册
	// r.POST("/api/auth/register", controller.Register)

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	// panic(r.Run(":8081"))
	r.Run(":8081")
}
