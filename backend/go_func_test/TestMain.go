package go_func_test

import (
	"log"
	"os"
	"testing"

	"gin_demo/common"

	"github.com/spf13/viper"
)

// 初始化 Viper
func initConfig() {
	viper.SetConfigName("config")    // 配置文件名（不带扩展名）
	viper.SetConfigType("yml")       // 配置文件格式
	viper.AddConfigPath("../config") // 配置文件路径（当前目录）
	viper.AddConfigPath("../")       // 兼容 go test 运行路径

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("加载配置文件失败: %v", err)
	}

	log.Println("配置文件加载成功")
}

// TestMain 运行所有测试前执行
func TestMain(m *testing.M) {
	initConfig()    // 加载配置
	common.InitDB() // 初始化数据库

	db := common.GetDB() // 初始化数据库

	log.Println("Connected to database:", db)

	// 运行测试
	code := m.Run()
	os.Exit(code)
}
