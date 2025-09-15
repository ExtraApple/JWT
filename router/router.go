package router

import (
	"QQZone/controllers"
	"QQZone/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/user")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
		// 受保护路由
		api.GET("/admin", middleware.AuthMiddleware(), middleware.AdminOnly(), controllers.AdminOnly)
	}
	return r
}
