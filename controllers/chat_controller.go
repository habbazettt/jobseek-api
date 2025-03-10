package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/habbazettt/jobseek-go/dto"
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

// @Summary     Send Message
// @Description Send a message to another user
// @Tags        chat
// @Accept      json
// @Produce     json
// @Param       body  body      dto.MessageRequest  true  "Message request"
// @Security    BearerAuth
// @Success     200   {object}  models.ChatMessage
// @Failure     400   {object}  map[string]interface{}
// @Failure     401   {object}  map[string]interface{}
// @Failure     500   {object}  map[string]interface{}
// @Router      /chat/send_message [post]
func (c *ChatController) SendMessage(ctx *gin.Context) {
	var request dto.MessageRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	senderID, exists := ctx.Get("user_id")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized: No user ID found in token")
		return
	}

	sender, err := c.chatService.GetUserByID(senderID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve sender details")
		return
	}

	chat, err := c.chatService.SendMessage(senderID.(uint), request.ReceiverID, request.Message)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	notificationMessage := "Anda menerima pesan baru dari " + sender.FullName
	_, err = c.notificationService.CreateNotification(request.ReceiverID, notificationMessage)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create notification")
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Message sent successfully", chat)
}

// @Summary     Get Messages
// @Description Get messages by sender and receiver
// @Tags        chat
// @Accept      json
// @Produce     json
// @Param       sender_id  query     uint    false  "Sender ID"
// @Param       receiver_id  query     uint    false  "Receiver ID"
// @Security    BearerAuth
// @Success     200   {object}  []models.ChatMessage
// @Failure     400   {object}  map[string]interface{}
// @Failure     401   {object}  map[string]interface{}
// @Failure     500   {object}  map[string]interface{}
// @Router      /chat/messages [get]
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


// @Summary     Get My Messages
// @Description Get messages of current user
// @Tags        chat
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Success     200   {object}  []models.ChatMessage
// @Failure     400   {object}  map[string]interface{}
// @Failure     401   {object}  map[string]interface{}
// @Failure     500   {object}  map[string]interface{}
// @Router      /chat/my-messages [get]
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
