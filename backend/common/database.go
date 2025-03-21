package common

import (
	"fmt"
	"log"
	"net/url"

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
	loc := viper.GetString("datasource.loc")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s", username, passwd, host, port, databaseName, charset, url.QueryEscape(loc))
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
