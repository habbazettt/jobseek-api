package services

import (
	"github.com/habbazettt/jobseek-go/models"
	"github.com/habbazettt/jobseek-go/repositories"
)

type ChatService interface {
	SendMessage(senderID, receiverID uint, message string) (*models.ChatMessage, error)
	GetMessages(senderID, receiverID *uint) ([]models.ChatMessage, error)
	GetMessagesByUser(userID uint) ([]models.ChatMessage, error)
	GetUserByID(userID uint) (*models.User, error)
}

type chatService struct {
	chatRepo repositories.ChatRepository
}

// ✅ Constructor ChatService
func NewChatService(chatRepo repositories.ChatRepository) ChatService {
	return &chatService{chatRepo}
}

// ✅ Kirim pesan
func (s *chatService) SendMessage(senderID, receiverID uint, message string) (*models.ChatMessage, error) {
	chat := models.ChatMessage{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Message:    message,
	}

	err := s.chatRepo.SaveMessage(&chat)
	if err != nil {
		return nil, err
	}

	return &chat, nil
}

// ✅ Ambil pesan antara dua user (atau semua jika filter kosong)
func (s *chatService) GetMessages(senderID, receiverID *uint) ([]models.ChatMessage, error) {
	return s.chatRepo.GetMessages(senderID, receiverID)
}

// ✅ Ambil semua pesan berdasarkan user_id
func (s *chatService) GetMessagesByUser(userID uint) ([]models.ChatMessage, error) {
	return s.chatRepo.GetMessagesByUser(userID)
}

func (s *chatService) GetUserByID(userID uint) (*models.User, error) {
	return s.chatRepo.GetUserByID(userID)
}
