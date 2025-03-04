package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/habbazettt/jobseek-go/controllers"
	"github.com/habbazettt/jobseek-go/middleware"
)

func ReviewRoutes(r *gin.Engine, reviewController *controllers.ReviewController) {
	reviews := r.Group("/api/v1/reviews")
	reviews.Use(middleware.AuthMiddleware())
	{
		reviews.POST("/", reviewController.CreateReview)                   // Tambah review
		reviews.GET("/:user_id", reviewController.GetReviewsByUser)        // Lihat review berdasarkan user ID
		reviews.GET("/me", reviewController.GetMyReviews)                  // Lihat review yang saya buat
		reviews.PUT("/:review_id", reviewController.UpdateReview)          // Update review
		reviews.DELETE("/:review_id", reviewController.DeleteReview)       // Hapus review
		reviews.GET("/rating/:user_id", reviewController.GetAverageRating) // Ambil rata-rata rating user
	}
}
