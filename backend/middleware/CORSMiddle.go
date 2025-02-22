package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// 允许携带 Cookies
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			// c.AbortWithStatus(200)

			// 204 No Content，表示预检请求通过
			c.AbortWithStatus(204)

		} else {
			c.Next()
		}
	}
}
