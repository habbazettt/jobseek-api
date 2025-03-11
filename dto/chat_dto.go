package dto

type MessageRequest struct {
	ReceiverID uint   `json:"receiver_id" binding:"required"`
	Message    string `json:"message" binding:"required"`
}
