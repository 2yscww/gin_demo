package common

import (
	"gin_demo/model"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	// v5版本使用的新方法
	jwt.RegisteredClaims
}

func ReleaseToken(user model.User) (string, error) {

	// * 先去掉，重新使用user
	// db := GetDB() // 获取数据库连接
	// var dbUser model.User
	// err := db.Where("telephone = ?", user.Telephone).First(&dbUser).Error

	// ? 从数据库中查询电话号码是否有注册，也就是用户是否存在

	// ? 检查查询是否出错
	// if err != nil {
	// 	if err == gorm.ErrRecordNotFound {
	// 		log.Println("未找到用户:", user.Telephone)
	// 		return "", fmt.Errorf("未找到用户")
	// 	}
	// 	log.Println("查询用户时发生错误:", err)
	// 	return "", err
	// }

	// 打印查询到的 ID
	// log.Println("查询到的用户 ID:", dbUser.ID)

	// log.Printf("数据库查询到的用户: %+v\n", dbUser)

	log.Println("现在jwt封装token,user的ID为:", user.ID)

	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{

		UserId: user.ID,
		// UserId: dbUser.ID,
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

	log.Println("生成的 JWT:", tokenString)

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	log.Printf("解析的 JWT: %+v\n", token)      // 打印 token 结构
	log.Printf("解析出的 Claims: %+v\n", claims) // 打印 claims 具体内容

	return token, claims, err
}

// ! 需要解决严重问题
// TODO 解决用户ID不匹配的问题
// 2025/03/06 12:06:30 C:/project_web/gin_vue/backend/controller/UserController.go:22 record not found
// [0.537ms] [rows:0] SELECT * FROM `users` WHERE telephone = '12345678911' AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
// ! 2025/03/06 12:06:30 查询到的用户 ID: 69
// [GIN] 2025/03/06 - 12:06:30 | 200 |     65.7859ms |       127.0.0.1 | POST     "/api/auth/register"
// ! 2025/03/06 12:06:30 claims中的userid为: 67
// ! 2025/03/06 12:06:30 claims赋值的userid为: 67
