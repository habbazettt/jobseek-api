package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/habbazettt/jobseek-go/config"
	"github.com/habbazettt/jobseek-go/controllers"
	"github.com/habbazettt/jobseek-go/repositories"
	"github.com/habbazettt/jobseek-go/routes"
	"github.com/habbazettt/jobseek-go/services"
)

func main() {
	db := config.ConnectDB()

	config.SetupCloudinary()

	r := gin.Default()

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
	fmt.Println("Server running on port " + port)
	log.Fatal(r.Run(":" + port))
}
