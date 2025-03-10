package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/habbazettt/jobseek-go/config"
	"github.com/habbazettt/jobseek-go/controllers"
	_ "github.com/habbazettt/jobseek-go/docs"
	"github.com/habbazettt/jobseek-go/repositories"
	"github.com/habbazettt/jobseek-go/routes"
	"github.com/habbazettt/jobseek-go/services"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           JobSeeker API
// @version         1.0
// @description     API untuk manajemen pencari kerja dan lowongan pekerjaan.
// @host      localhost:8080
// @BasePath  /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Masukkan token dalam format "Bearer <token>"
func main() {
	db := config.ConnectDB()
	config.SetupCloudinary()

	r := gin.Default()
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	chatRepo := repositories.NewChatRepository(db)
	jobRepo := repositories.NewJobRepository(db)
	userRepo := repositories.NewUserRepository(db)
	savedRepo := repositories.NewSavedRepository(db)
	notificationRepo := repositories.NewNotificationRepository(db)
	proposalRepo := repositories.NewProposalRepository(db)
	reviewRepo := repositories.NewReviewRepository(db)

	chatService := services.NewChatService(chatRepo)
	notificationService := services.NewNotificationService(notificationRepo)
	proposalService := services.NewProposalService(proposalRepo, jobRepo, userRepo)
	reviewService := services.NewReviewService(reviewRepo)
	savedService := services.NewSavedService(savedRepo, jobRepo, userRepo)

	notificationController := controllers.NewNotificationController(notificationService)
	chatController := controllers.NewChatController(chatService, notificationService)
	proposalController := controllers.NewProposalController(proposalService)
	reviewController := controllers.NewReviewController(reviewService)
	savedController := controllers.NewSavedController(savedService)

	routes.AuthRoutes(r)
	routes.JobRoutes(r, db)
	routes.UserRoutes(r, db)
	routes.ChatRoutes(r, chatController)
	routes.NotificationRoutes(r, notificationController)
	routes.ProposalRoutes(r, proposalController)
	routes.ReviewRoutes(r, reviewController)
	routes.SavedRoutes(r, savedController)

	port := "8080"
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(r.Run(":" + port))
}
