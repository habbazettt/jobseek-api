package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/habbazettt/jobseek-go/controllers"
	"github.com/habbazettt/jobseek-go/middleware"
	"github.com/habbazettt/jobseek-go/repositories"
	"github.com/habbazettt/jobseek-go/services"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	user := r.Group("/api/v1/users")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/", middleware.AdminMiddleware(), userController.GetAllUsers)
		user.GET("/me", userController.GetCurrentUser)
		user.GET("/:id", middleware.AdminMiddleware(), userController.GetUserByID)
		user.PUT("/:id", userController.UpdateUser)
		user.DELETE("/:id", userController.DeleteUser)
	}
}
