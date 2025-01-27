package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	protected := router.Group("/protected")
	protected.Use(AuthMiddleware()) // Middleware untuk autentikasi
	{
		protected.GET("/dashboard", controllers.Dashboard)
	}
}
