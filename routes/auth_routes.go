package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/habbazettt/jobseek-go/controllers"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/register", controllers.RegisterUser)
		auth.POST("/login", controllers.LoginUser)
	}
}
