package services

import (
	"github.com/habbazettt/jobseek-go/models"
	"github.com/habbazettt/jobseek-go/repositories"
)

type NotificationService interface {
	CreateNotification(userID uint, message string) (*models.Notification, error)
	GetNotifications(userID uint) ([]models.Notification, error)
	MarkAsRead(notificationID uint) error
	MarkAllAsRead(userID uint) error
	DeleteNotification(notificationID uint) error
	DeleteAllNotifications(userID uint) error
	GetNotificationByID(notificationID uint) (*models.Notification, error)
}

type notificationService struct {
	notificationRepo repositories.NotificationRepository
}

func NewNotificationService(notificationRepo repositories.NotificationRepository) NotificationService {
	return &notificationService{notificationRepo}
}

func (s *notificationService) CreateNotification(userID uint, message string) (*models.Notification, error) {
	notification := models.Notification{
		UserID:  userID,
		Message: message,
		IsRead:  false,
	}

	err := s.notificationRepo.CreateNotification(&notification)
	if err != nil {
		return nil, err
	}

	return &notification, nil
}

func (s *notificationService) GetNotifications(userID uint) ([]models.Notification, error) {
	return s.notificationRepo.GetNotificationsByUser(userID)
}

func (s *notificationService) MarkAsRead(notificationID uint) error {
	return s.notificationRepo.MarkAsRead(notificationID)
}

func (s *notificationService) MarkAllAsRead(userID uint) error {
	return s.notificationRepo.MarkAllAsRead(userID)
}

func (s *notificationService) DeleteNotification(notificationID uint) error {
	return s.notificationRepo.DeleteNotification(notificationID)
}

func (s *notificationService) DeleteAllNotifications(userID uint) error {
	return s.notificationRepo.DeleteAllNotifications(userID)
}

func (s *notificationService) GetNotificationByID(notificationID uint) (*models.Notification, error) {
	return s.notificationRepo.GetNotificationByID(notificationID)
}
