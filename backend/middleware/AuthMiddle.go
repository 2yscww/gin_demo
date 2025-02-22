package middleware

import (
	"gin_demo/common"
	"gin_demo/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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

		// ? token.Valid 为 true 时表示token有效

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}

		// token通过了验证,开始获取token中的ID

		userID := claims.UserId
		db := common.GetDB()

		var user model.User

		// ? 搞清楚这句话做什么的
		db.First(&user, userID)

		// 验证用户
		if userID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}

		// 用户存在 将 user 的信息写入上下文

		c.Set("user", user)

		c.Next()

	}
}
