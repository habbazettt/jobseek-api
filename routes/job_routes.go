package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/habbazettt/jobseek-go/controllers"
	"github.com/habbazettt/jobseek-go/middleware"
	"github.com/habbazettt/jobseek-go/repositories"
	"github.com/habbazettt/jobseek-go/services"
	"gorm.io/gorm"
)

func JobRoutes(r *gin.Engine, db *gorm.DB) {
	jobRepo := repositories.NewJobRepository(db)
	jobService := services.NewJobService(jobRepo)
	jobController := controllers.NewJobController(jobService)

	job := r.Group("/api/v1/jobs")
	job.Use(middleware.AuthMiddleware())
	{
		job.POST("/", jobController.CreateJob)
		job.GET("/", jobController.GetJobs)
		job.GET("/:id", jobController.GetJobByID)
		job.PUT("/:id", jobController.UpdateJob)
		job.DELETE("/:id", jobController.DeleteJob)
	}
}
