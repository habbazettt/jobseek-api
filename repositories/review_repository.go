package repositories

import (
	"github.com/habbazettt/jobseek-go/dto"
	"github.com/habbazettt/jobseek-go/models"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	CreateReview(review *models.Review) error
	GetReviewsByUserID(userID uint) ([]dto.ReviewResponse, error)
	GetReviewsByReviewerID(reviewerID uint) ([]dto.ReviewResponse, error)
	GetReviewByID(reviewID uint) (*models.Review, error)
	UpdateReview(reviewID uint, rating float64, comment string) error
	DeleteReview(reviewID uint, reviewerID uint) error
	GetAverageRating(userID uint) (float64, int, error)
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db}
}

// ✅ 1. Simpan review baru
func (r *reviewRepository) CreateReview(review *models.Review) error {
	return r.db.Create(review).Error
}

// ✅ 2. Ambil semua review yang diterima oleh user tertentu
func (r *reviewRepository) GetReviewsByUserID(userID uint) ([]dto.ReviewResponse, error) {
	var reviews []dto.ReviewResponse
	err := r.db.Table("reviews").
		Where("reviewed_id = ?", userID).
		Scan(&reviews).Error
	return reviews, err
}

// ✅ 3. Ambil semua review yang diberikan oleh reviewer tertentu
func (r *reviewRepository) GetReviewsByReviewerID(reviewerID uint) ([]dto.ReviewResponse, error) {
	var reviews []dto.ReviewResponse
	err := r.db.Table("reviews").
		Where("reviewer_id = ?", reviewerID).
		Scan(&reviews).Error
	return reviews, err
}

// ✅ 4. Ambil satu review berdasarkan ID
func (r *reviewRepository) GetReviewByID(reviewID uint) (*models.Review, error) {
	var review models.Review
	err := r.db.Where("id = ?", reviewID).First(&review).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}

// ✅ 5. Update review (hanya reviewer yang bisa)
func (r *reviewRepository) UpdateReview(reviewID uint, rating float64, comment string) error {
	return r.db.Model(&models.Review{}).
		Where("id = ?", reviewID).
		Updates(map[string]interface{}{"rating": rating, "comment": comment}).Error
}

// ✅ 6. Hapus review (hanya reviewer yang bisa)
func (r *reviewRepository) DeleteReview(reviewID, reviewerID uint) error {
	return r.db.Where("id = ? AND reviewer_id = ?", reviewID, reviewerID).
		Delete(&models.Review{}).Error
}

// ✅ 7. Hitung rata-rata rating user
func (r *reviewRepository) GetAverageRating(userID uint) (float64, int, error) {
	var result struct {
		Average float64
		Total   int
	}
	err := r.db.Table("reviews").
		Select("AVG(rating) AS average, COUNT(*) AS total").
		Where("reviewed_id = ?", userID).
		Scan(&result).Error

	return result.Average, result.Total, err
}
