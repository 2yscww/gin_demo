package common

import (
	"fmt"
	"gin_demo/model"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	// v5版本使用的新方法
	jwt.RegisteredClaims
}

func ReleaseToken(user model.User) (string, error) {

	db := GetDB() // 获取数据库连接
	var dbUser model.User
	err := db.Where("telephone = ?", user.Telephone).First(&dbUser).Error

	// ? 从数据库中查询电话号码是否有注册，也就是用户是否存在

	// 检查查询是否出错
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("未找到用户:", user.Telephone)
			return "", fmt.Errorf("未找到用户")
		}
		log.Println("查询用户时发生错误:", err)
		return "", err
	}

	// 打印查询到的 ID
	log.Println("生成的用户 ID:", dbUser.ID)

	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{

		// UserId: user.ID,
		UserId: dbUser.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "Administrator",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}
