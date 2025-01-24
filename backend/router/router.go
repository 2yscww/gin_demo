package router

import (
	"gin_demo/controller"

	"github.com/gin-gonic/gin"
)

func CollectRtoutes(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	// r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	return r
}
