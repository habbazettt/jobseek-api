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

func NewChatService(chatRepo repositories.ChatRepository) ChatService {
	return &chatService{chatRepo}
}

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

func (s *chatService) GetMessages(senderID, receiverID *uint) ([]models.ChatMessage, error) {
	return s.chatRepo.GetMessages(senderID, receiverID)
}

func (s *chatService) GetMessagesByUser(userID uint) ([]models.ChatMessage, error) {
	return s.chatRepo.GetMessagesByUser(userID)
}

func (s *chatService) GetUserByID(userID uint) (*models.User, error) {
	return s.chatRepo.GetUserByID(userID)
}
