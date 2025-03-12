package middleware

import (
	"gin_demo/common"
	"gin_demo/model"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取authorization header
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}
		// TODO 完善中间件直接返回敏感信息的问题
		// ! 不能将token的敏感信息直接返回

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)

		log.Println("claims中的userid为:", claims.UserId)

		// ? token.Valid 为 true 时表示token有效

		// 2025/02/27 22:40:10 Token claims: &{0 {Administrator user token []

		// if err != nil || !token.Valid {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "token 出错"})
		// 	c.Abort()
		// 	return
		// }

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}
		if !token.Valid {
			// ! token 无效
			c.JSON(http.StatusUnauthorized, gin.H{"code": 500, "msg": "内部服务器错误"})
			c.Abort()
			return
		}

		// token通过了验证,开始获取token中的ID

		userID := claims.UserId

		log.Printf("AuthMiddleware - 解析出 UserId: %d\n", userID)
		db := common.GetDB()

		if db == nil {
			log.Println("数据库连接为空")
		}

		var user model.User

		// ? 搞清楚这句话做什么的
		// db.First(&user, userID)

		result := db.First(&user, userID)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				// 没有找到对应用户
				c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "无此用户"})
			} else {
				// 其他错误
				c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "数据库查询出错"})
			}
			c.Abort()
			return
		}

		// 验证用户
		if userID == 0 {
			// c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "无此用户"})
			c.Abort()
			return
		}

		// 用户存在 将 user 的信息写入上下文

		c.Set("user", user)

		c.Next()

	}
}
