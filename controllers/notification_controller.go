package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/habbazettt/jobseek-go/services"
	"github.com/habbazettt/jobseek-go/utils"
)

type NotificationController struct {
	notificationService services.NotificationService
}

func NewNotificationController(notificationService services.NotificationService) *NotificationController {
	return &NotificationController{notificationService}
}

// ✅ Ambil semua notifikasi milik user yang login
func (c *NotificationController) GetNotifications(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")

	notifications, err := c.notificationService.GetNotifications(userID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Notifications retrieved successfully", notifications)
}

// ✅ Tandai satu notifikasi sebagai sudah dibaca & tampilkan datanya
func (c *NotificationController) MarkAsRead(ctx *gin.Context) {
	notificationID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid notification ID")
		return
	}

	// Ambil notifikasi sebelum diperbarui
	notification, err := c.notificationService.GetNotificationByID(uint(notificationID))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "Notification not found")
		return
	}

	// Tandai sebagai sudah dibaca
	err = c.notificationService.MarkAsRead(uint(notificationID))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to mark notification as read")
		return
	}

	// Update status is_read pada response
	notification.IsRead = true

	utils.SuccessResponse(ctx, http.StatusOK, "Notification marked as read", notification)
}

// ✅ Tandai semua notifikasi sebagai sudah dibaca & tampilkan datanya
func (c *NotificationController) MarkAllAsRead(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")

	// Ambil semua notifikasi sebelum diperbarui
	notifications, err := c.notificationService.GetNotifications(userID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve notifications")
		return
	}

	// Tandai semua sebagai sudah dibaca
	err = c.notificationService.MarkAllAsRead(userID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to mark all notifications as read")
		return
	}

	// Update status is_read pada response
	for i := range notifications {
		notifications[i].IsRead = true
	}

	utils.SuccessResponse(ctx, http.StatusOK, "All notifications marked as read", notifications)
}

// ✅ Hapus satu notifikasi berdasarkan ID & tampilkan notifikasi yang dihapus
func (c *NotificationController) DeleteNotification(ctx *gin.Context) {
	notificationID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid notification ID")
		return
	}

	// Ambil notifikasi sebelum dihapus
	notification, err := c.notificationService.GetNotificationByID(uint(notificationID))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "Notification not found")
		return
	}

	// Hapus notifikasi
	err = c.notificationService.DeleteNotification(uint(notificationID))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete notification")
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Notification deleted successfully", notification)
}

// ✅ Hapus semua notifikasi user yang login & tampilkan daftar notifikasi yang dihapus
func (c *NotificationController) DeleteAllNotifications(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")

	// Ambil semua notifikasi sebelum dihapus
	notifications, err := c.notificationService.GetNotifications(userID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve notifications")
		return
	}

	// Hapus semua notifikasi
	err = c.notificationService.DeleteAllNotifications(userID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete all notifications")
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "All notifications deleted successfully", notifications)
}
