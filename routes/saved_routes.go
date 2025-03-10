package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/habbazettt/jobseek-go/controllers"
	"github.com/habbazettt/jobseek-go/middleware"
)

func SavedRoutes(r *gin.Engine, savedController *controllers.SavedController) {
	saved := r.Group("/api/v1/saved")
	saved.Use(middleware.AuthMiddleware())
	{
		// Freelancer menyimpan & menghapus job favorit
		saved.POST("/jobs/:job_id", savedController.SaveJob)          // Simpan job favorit
		saved.GET("/jobs", savedController.GetSavedJobs)              // Ambil daftar job favorit
		saved.DELETE("/jobs/:job_id", savedController.RemoveSavedJob) // Hapus job favorit

		// Perusahaan menyimpan & menghapus freelancer favorit
		saved.POST("/freelancers/:freelancer_id", savedController.SaveFreelancer)          // Simpan freelancer favorit
		saved.GET("/freelancers", savedController.GetSavedFreelancers)                     // Ambil daftar freelancer favorit
		saved.DELETE("/freelancers/:freelancer_id", savedController.RemoveSavedFreelancer) // Hapus freelancer favorit
	}
}
