package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/habbazettt/jobseek-go/controllers"
	"github.com/habbazettt/jobseek-go/middleware"
)

func ChatRoutes(r *gin.Engine, chatController *controllers.ChatController) {
	chat := r.Group("/api/v1/chat")
	chat.Use(middleware.AuthMiddleware()) // ✅ Semua route chat butuh autentikasi
	{
		chat.POST("/send_message", chatController.SendMessage) // ✅ API untuk mengirim pesan
		chat.GET("/messages", chatController.GetMessages)      // Get all messages (by sender & receiver)
		chat.GET("/my-messages", chatController.GetMyMessages) // Get user messages from token
	}
}
