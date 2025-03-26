package router

import (
	"gin_demo/controller"
	"gin_demo/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {

	// ! middleware.RecoveryMiddleware() 是调用了recover方法给前端返回了错误信息，测试时有用

	r.Use(middleware.CORSMiddle(), middleware.RecoveryMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	categoryRoutes := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("", categoryController.Create)
	categoryRoutes.PUT("/:id", categoryController.Update)
	categoryRoutes.GET("/:id", categoryController.Show)
	categoryRoutes.DELETE("/:id", categoryController.Delete)

	postRoutes := r.Group("/posts")
	postRoutes.Use(middleware.AuthMiddleware())
	postController := controller.NewPostController()
	postRoutes.POST("", postController.Create)
	postRoutes.PUT("/:id", postController.Update)
	postRoutes.GET("/:id", postController.Show)
	postRoutes.DELETE("/:id", postController.Delete)

	return r
}
