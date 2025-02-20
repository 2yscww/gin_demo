package main

import (
	"log"
	"os"

	"gin_demo/common"
	"gin_demo/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 定义数据库全局变量
// var db *gorm.DB

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {

	// *读取配置文件
	InitConfig()

	common.InitDB()
	db := common.GetDB() // 初始化数据库

	log.Println("Connected to database:", db)

	// 1.创建路由
	r := gin.Default()

	r = router.CollectRtoutes(r)

	// 读取端口配置
	port := viper.GetString("server.port")

	if port != "" {
		panic(r.Run(":" + port))
	}

	// * 不启用HTTPS
	// 启用HTTPS
	// certFile := "./verify/cert.pem"
	// keyFile := "./verify/key.pem"

	// err := r.RunTLS(":"+port, certFile, keyFile)
	// if err != nil {
	// 	log.Fatal("Failed to start HTTPS server:", err)
	// }
	// * 不启用HTTPS

	panic(r.Run())

	// r.Run(":8081")

	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response

	// api注册
	// r.POST("/api/auth/register", controller.Register)

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	// panic(r.Run(":8081"))
}
