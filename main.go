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
	chatService := services.NewChatService(chatRepo)

	notificationRepo := repositories.NewNotificationRepository(db)
	notificationService := services.NewNotificationService(notificationRepo)
	notificationController := controllers.NewNotificationController(notificationService)

	chatController := controllers.NewChatController(chatService, notificationService)

	routes.AuthRoutes(r)
	routes.JobRoutes(r, db)
	routes.UserRoutes(r, db)
	routes.JobRoutes(r, db)
	routes.UserRoutes(r, db)
	routes.ChatRoutes(r, chatController)
	routes.NotificationRoutes(r, notificationController)

	port := "8080"
	fmt.Println("Server running on port " + port)
	log.Fatal(r.Run(":" + port))
}
