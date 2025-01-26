package common

import (
	"fmt"
	"log"

	"gin_demo/model"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义数据库全局变量
var db *gorm.DB

// 连接mysql数据库
func InitDB() {
	username := viper.GetString("datasource.username")
	passwd := viper.GetString("datasource.password")
	databaseName := viper.GetString("datasource.database")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	charset := viper.GetString("datasource.charset")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", username, passwd, host, port, databaseName, charset)
	var err error

	// 初始化数据库连接
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 自动迁移
	db.AutoMigrate(&model.User{})
}

func GetDB() *gorm.DB {
	return db
}
