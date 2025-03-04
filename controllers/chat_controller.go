package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/habbazettt/jobseek-go/services"
	"github.com/habbazettt/jobseek-go/utils"
)

type ChatController struct {
	chatService         services.ChatService
	notificationService services.NotificationService
}

func NewChatController(chatService services.ChatService, notificationService services.NotificationService) *ChatController {
	return &ChatController{chatService, notificationService}
}

// ✅ Kirim pesan
func (c *ChatController) SendMessage(ctx *gin.Context) {
	var request struct {
		ReceiverID uint   `json:"receiver_id" binding:"required"`
		Message    string `json:"message" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// ✅ Ambil sender_id dari token JWT
	senderID, exists := ctx.Get("user_id")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized: No user ID found in token")
		return
	}

	// ✅ Ambil informasi user (full_name) berdasarkan sender_id
	sender, err := c.chatService.GetUserByID(senderID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve sender details")
		return
	}

	// ✅ Kirim pesan
	chat, err := c.chatService.SendMessage(senderID.(uint), request.ReceiverID, request.Message)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// ✅ Notifikasi dengan full_name pengirim
	notificationMessage := "Anda menerima pesan baru dari " + sender.FullName
	_, err = c.notificationService.CreateNotification(request.ReceiverID, notificationMessage)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create notification")
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Message sent successfully", chat)
}

// ✅ Ambil pesan berdasarkan filter sender_id dan receiver_id (atau semua pesan jika tidak ada filter)
func (c *ChatController) GetMessages(ctx *gin.Context) {
	var senderID, receiverID *uint

	// Ambil sender_id dari query jika ada
	if senderIDParam := ctx.Query("sender_id"); senderIDParam != "" {
		parsedID, err := strconv.Atoi(senderIDParam)
		if err == nil {
			id := uint(parsedID)
			senderID = &id
		}
	}

	// Ambil receiver_id dari query jika ada
	if receiverIDParam := ctx.Query("receiver_id"); receiverIDParam != "" {
		parsedID, err := strconv.Atoi(receiverIDParam)
		if err == nil {
			id := uint(parsedID)
			receiverID = &id
		}
	}

	// Panggil service untuk mengambil pesan
	messages, err := c.chatService.GetMessages(senderID, receiverID)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve messages")
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Messages retrieved successfully", messages)
}

// ✅ Ambil semua pesan dari user yang sedang login
func (c *ChatController) GetMyMessages(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized: No user ID found in token")
		return
	}

	messages, err := c.chatService.GetMessagesByUser(userID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "User messages retrieved successfully", messages)
}
