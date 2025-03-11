package go_func_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gin_demo/controller"
	"gin_demo/middleware"

	"github.com/gin-gonic/gin"
)

// 你的注册生成的 Token
const validToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjgwLCJpc3MiOiJBZG1pbmlzdHJhdG9yIiwic3ViIjoidXNlciB0b2tlbiIsImV4cCI6MTc0MjI4MDU5NywiaWF0IjoxNzQxNjc1Nzk3fQ.FydJcDUuIu1wlErhI3_GnZmW9U9D2DEdOfvHrp4P4vk" // 这里替换成你手动获取的 token

func TestAuthMiddlewareWithValidToken(t *testing.T) {
	gin.SetMode(gin.TestMode) // 设置 Gin 测试模式

	// 初始化 Gin 服务器
	r := gin.Default()

	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	// 构造 HTTP 请求
	req, _ := http.NewRequest("GET", "/api/auth/info", nil)
	req.Header.Set("Authorization", "Bearer "+validToken) // 传入真实 token

	// 记录 HTTP 响应
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// 打印返回的 HTTP 状态码和 Body，方便你手动检查
	t.Log("HTTP 状态码:", w.Code)
	t.Log("Response Body:", w.Body.String())
}
