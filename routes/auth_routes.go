package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/habbazettt/jobseek-go/controllers"
	"github.com/habbazettt/jobseek-go/middleware"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/register", controllers.RegisterUser)
		auth.POST("/login", controllers.LoginUser)
	}

	protected := r.Group("/api/v1/user")
	protected.Use(middleware.AuthMiddleware()) // Middleware JWT diaktifkan di sini
	{
		protected.GET("/profile", controllers.UserProfile)
	}
}
