package common

import (
	"fmt"
	"log"

	"gin_demo/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义数据库全局变量
var db *gorm.DB

// 连接mysql数据库
func InitDB() {
	username := "root"
	passwd := "w86#qNwV"
	databaseName := "go_test"

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, passwd, databaseName)
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
