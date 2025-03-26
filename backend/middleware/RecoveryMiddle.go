package middleware

import (
	"fmt"
	"gin_demo/response"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				response.Fail(c, nil, fmt.Sprint(err))
			}
		}()

		c.Next()
	}
}
