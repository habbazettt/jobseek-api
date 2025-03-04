package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/habbazettt/jobseek-go/controllers"
	"github.com/habbazettt/jobseek-go/middleware"
)

func NotificationRoutes(r *gin.Engine, notificationController *controllers.NotificationController) {
	notifications := r.Group("/api/v1/notifications")
	notifications.Use(middleware.AuthMiddleware())
	{
		notifications.GET("/", notificationController.GetNotifications)                    // ✅ Ambil semua notifikasi milik user yang login
		notifications.PATCH("/:id/read", notificationController.MarkAsRead)                // ✅ Tandai satu notifikasi sebagai sudah dibaca
		notifications.PATCH("/read-all", notificationController.MarkAllAsRead)             // ✅ Tandai semua notifikasi sebagai sudah dibaca
		notifications.DELETE("/:id", notificationController.DeleteNotification)            // ✅ Hapus satu notifikasi berdasarkan ID
		notifications.DELETE("/delete-all", notificationController.DeleteAllNotifications) // ✅ Hapus semua notifikasi user yang login
	}
}
