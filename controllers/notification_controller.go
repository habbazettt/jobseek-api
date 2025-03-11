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

// GetNotifications retrieves all notifications for the currently authenticated user.
// @Summary Get Notifications
// @Description Fetches all notifications associated with the logged-in user.
// @Tags notifications
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Notification "Notifications retrieved successfully"
// @Failure 500 {object} utils.ErrorResponseSwagger "Failed to retrieve notifications"
// @Router /notifications [get]

func (c *NotificationController) GetNotifications(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")

	notifications, err := c.notificationService.GetNotifications(userID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Notifications retrieved successfully", notifications)
}

// MarkAsRead marks a specific notification as read based on the provided notification ID.
// @Summary Mark Notification As Read
// @Description Marks a single notification as read using its ID, updating its status.
// @Tags notifications
// @Accept json
// @Produce json
// @Param id path int true "Notification ID"
// @Security BearerAuth
// @Success 200 {object} models.Notification "Notification marked as read"
// @Failure 400 {object} utils.ErrorResponseSwagger "Invalid notification ID"
// @Failure 404 {object} utils.ErrorResponseSwagger "Notification not found"
// @Failure 500 {object} utils.ErrorResponseSwagger "Failed to mark notification as read"
// @Router /notifications/{id}/read [patch]

func (c *NotificationController) MarkAsRead(ctx *gin.Context) {
	notificationID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid notification ID")
		return
	}

	notification, err := c.notificationService.GetNotificationByID(uint(notificationID))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "Notification not found")
		return
	}

	err = c.notificationService.MarkAsRead(uint(notificationID))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to mark notification as read")
		return
	}

	notification.IsRead = true

	utils.SuccessResponse(ctx, http.StatusOK, "Notification marked as read", notification)
}

// MarkAllAsRead marks all notifications as read for the currently authenticated user.
// @Summary Mark All Notifications As Read
// @Description Marks all notifications associated with the logged-in user as read.
// @Tags notifications
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Notification "All notifications marked as read"
// @Failure 500 {object} utils.ErrorResponseSwagger "Failed to mark all notifications as read"
// @Router /notifications/read-all [patch]
func (c *NotificationController) MarkAllAsRead(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")

	notifications, err := c.notificationService.GetNotifications(userID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve notifications")
		return
	}

	err = c.notificationService.MarkAllAsRead(userID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to mark all notifications as read")
		return
	}

	for i := range notifications {
		notifications[i].IsRead = true
	}

	utils.SuccessResponse(ctx, http.StatusOK, "All notifications marked as read", notifications)
}

// DeleteNotification deletes a notification by ID. Only the notification owner can
// delete it. The function verifies the user identity from the token to ensure
// authorization.
//
// @Summary Delete Notification
// @Description Deletes a notification by ID.
// @Tags notifications
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Notification ID"
// @Success 200 {object} models.Notification "Notification deleted successfully"
// @Failure 400 {object} utils.ErrorResponseSwagger "Invalid notification ID"
// @Failure 404 {object} utils.ErrorResponseSwagger "Notification not found"
// @Failure 500 {object} utils.ErrorResponseSwagger "Failed to delete notification"
// @Router /notifications/{id} [delete]
func (c *NotificationController) DeleteNotification(ctx *gin.Context) {
	notificationID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid notification ID")
		return
	}

	notification, err := c.notificationService.GetNotificationByID(uint(notificationID))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "Notification not found")
		return
	}

	err = c.notificationService.DeleteNotification(uint(notificationID))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete notification")
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Notification deleted successfully", notification)
}

// DeleteAllNotifications deletes all notifications associated with the currently
// authenticated user. Only the user that owns the notifications can delete them.
// The function verifies the user identity from the token to ensure authorization.
//
// @Summary Delete All Notifications
// @Description Deletes all notifications associated with the logged-in user.
// @Tags notifications
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Notification "All notifications deleted successfully"
// @Failure 500 {object} utils.ErrorResponseSwagger "Failed to delete all notifications"
// @Router /notifications/delete-all [delete]
func (c *NotificationController) DeleteAllNotifications(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")

	notifications, err := c.notificationService.GetNotifications(userID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve notifications")
		return
	}

	err = c.notificationService.DeleteAllNotifications(userID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete all notifications")
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "All notifications deleted successfully", notifications)
}
