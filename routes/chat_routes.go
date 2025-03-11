package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/habbazettt/jobseek-go/controllers"
	"github.com/habbazettt/jobseek-go/middleware"
	"github.com/habbazettt/jobseek-go/services"
	"github.com/habbazettt/jobseek-go/websocketgo"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ChatRoutes(r *gin.Engine, chatController *controllers.ChatController, chatService services.ChatService) {
	chat := r.Group("/api/v1/chat")
	chat.Use(middleware.AuthMiddleware())
	{
		// Routes untuk REST API Chat
		chat.POST("/send_message", chatController.SendMessage)
		chat.GET("/messages", chatController.GetMessages)
		chat.GET("/my-messages", chatController.GetMyMessages)

		// WebSocket Route
		chat.GET("/ws", func(c *gin.Context) {
			userID, exists := c.Get("user_id") // Ambil user_id dari middleware
			if !exists {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}

			conn, err := upgrader.Upgrade(c.Writer, c.Request, nil) // Upgrade ke WebSocket
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade WebSocket"})
				return
			}

			websocketgo.HandleConnections(conn, userID.(uint), chatService)
		})
	}
}
