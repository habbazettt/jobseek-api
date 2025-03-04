package repositories

import (
	"github.com/habbazettt/jobseek-go/models"
	"gorm.io/gorm"
)

type ChatRepository interface {
	SaveMessage(message *models.ChatMessage) error
	GetMessages(senderID, receiverID *uint) ([]models.ChatMessage, error)
	GetMessagesByUser(userID uint) ([]models.ChatMessage, error)
	GetUserByID(userID uint) (*models.User, error)
}

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) ChatRepository {
	return &chatRepository{db}
}

// ✅ Simpan pesan ke database
func (r *chatRepository) SaveMessage(message *models.ChatMessage) error {
	return r.db.Create(message).Error
}

// ✅ Ambil pesan berdasarkan filter (atau semua pesan jika filter kosong)
func (r *chatRepository) GetMessages(senderID, receiverID *uint) ([]models.ChatMessage, error) {
	var messages []models.ChatMessage
	query := r.db.Order("created_at ASC")

	// Tambahkan filter jika ada sender_id dan receiver_id
	if senderID != nil && receiverID != nil {
		query = query.Where("sender_id = ? AND receiver_id = ?", *senderID, *receiverID)
	}

	err := query.Find(&messages).Error
	return messages, err
}

// ✅ Ambil semua pesan berdasarkan user_id
func (r *chatRepository) GetMessagesByUser(userID uint) ([]models.ChatMessage, error) {
	var messages []models.ChatMessage
	err := r.db.
		Where("sender_id = ? OR receiver_id = ?", userID, userID).
		Order("created_at DESC").
		Find(&messages).Error
	return messages, err
}

func (r *chatRepository) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
