package repositories

import (
	"github.com/habbazettt/jobseek-go/models"
	"gorm.io/gorm"
)

type NotificationRepository interface {
	CreateNotification(notification *models.Notification) error
	GetNotificationsByUser(userID uint) ([]models.Notification, error)
	MarkAsRead(notificationID uint) error
	MarkAllAsRead(userID uint) error
	DeleteNotification(notificationID uint) error
	DeleteAllNotifications(userID uint) error
	GetNotificationByID(notificationID uint) (*models.Notification, error)
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db}
}

func (r *notificationRepository) CreateNotification(notification *models.Notification) error {
	return r.db.Create(notification).Error
}

func (r *notificationRepository) GetNotificationsByUser(userID uint) ([]models.Notification, error) {
	var notifications []models.Notification
	err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&notifications).Error
	return notifications, err
}

func (r *notificationRepository) MarkAsRead(notificationID uint) error {
	return r.db.Model(&models.Notification{}).Where("id = ?", notificationID).Update("is_read", true).Error
}

func (r *notificationRepository) MarkAllAsRead(userID uint) error {
	return r.db.Model(&models.Notification{}).Where("user_id = ?", userID).Update("is_read", true).Error
}

func (r *notificationRepository) DeleteNotification(notificationID uint) error {
	return r.db.Delete(&models.Notification{}, notificationID).Error
}

func (r *notificationRepository) DeleteAllNotifications(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.Notification{}).Error
}

func (r *notificationRepository) GetNotificationByID(notificationID uint) (*models.Notification, error) {
	var notification models.Notification
	err := r.db.Where("id = ?", notificationID).First(&notification).Error
	if err != nil {
		return nil, err
	}
	return &notification, nil
}
