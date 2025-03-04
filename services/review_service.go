package services

import (
	"errors"

	"github.com/habbazettt/jobseek-go/dto"
	"github.com/habbazettt/jobseek-go/models"
	"github.com/habbazettt/jobseek-go/repositories"
)

type ReviewService interface {
	CreateReview(request dto.CreateReviewRequest, reviewerID uint) (*dto.ReviewResponse, error)
	GetReviewsByUserID(userID uint) ([]dto.ReviewResponse, error)
	GetReviewsByReviewerID(reviewerID uint) ([]dto.ReviewResponse, error)
	UpdateReview(reviewID uint, rating float64, comment string, userID uint) (*dto.ReviewResponse, error)
	DeleteReview(reviewID uint, reviewerID uint) error
	GetAverageRating(userID uint) (*dto.AverageRatingResponse, error)
}

type reviewService struct {
	reviewRepo repositories.ReviewRepository
}

func NewReviewService(reviewRepo repositories.ReviewRepository) ReviewService {
	return &reviewService{reviewRepo}
}

// ✅ 1. Buat Review
func (s *reviewService) CreateReview(request dto.CreateReviewRequest, reviewerID uint) (*dto.ReviewResponse, error) {
	if reviewerID == request.ReviewedID {
		return nil, errors.New("you cannot review yourself")
	}

	review := models.Review{
		ReviewerID: reviewerID,
		ReviewedID: request.ReviewedID,
		Rating:     request.Rating,
		Comment:    request.Comment,
	}

	err := s.reviewRepo.CreateReview(&review)
	if err != nil {
		return nil, err
	}

	response := dto.ReviewResponse{
		ID:         review.ID,
		ReviewerID: review.ReviewerID,
		ReviewedID: review.ReviewedID,
		Rating:     review.Rating,
		Comment:    review.Comment,
		CreatedAt:  review.CreatedAt,
	}

	return &response, nil
}

// ✅ 2. Ambil Review Berdasarkan User ID
func (s *reviewService) GetReviewsByUserID(userID uint) ([]dto.ReviewResponse, error) {
	return s.reviewRepo.GetReviewsByUserID(userID)
}

// ✅ 3. Ambil Review yang Diberikan oleh Reviewer
func (s *reviewService) GetReviewsByReviewerID(reviewerID uint) ([]dto.ReviewResponse, error) {
	return s.reviewRepo.GetReviewsByReviewerID(reviewerID)
}

// ✅ 4. Update Review
func (s *reviewService) UpdateReview(reviewID uint, rating float64, comment string, userID uint) (*dto.ReviewResponse, error) {
	// Ambil review yang akan diperbarui
	review, err := s.reviewRepo.GetReviewByID(reviewID)
	if err != nil {
		return nil, errors.New("review not found")
	}

	// Pastikan hanya reviewer yang bisa mengedit review mereka
	if review.ReviewerID != userID {
		return nil, errors.New("unauthorized: you can only update your own reviews")
	}

	// Lakukan update pada review
	err = s.reviewRepo.UpdateReview(reviewID, rating, comment)
	if err != nil {
		return nil, err
	}

	// Ambil kembali data review yang telah diperbarui
	updatedReview, err := s.reviewRepo.GetReviewByID(reviewID)
	if err != nil {
		return nil, err
	}

	// Buat response DTO
	response := &dto.ReviewResponse{
		ID:         updatedReview.ID,
		ReviewerID: updatedReview.ReviewerID,
		ReviewedID: updatedReview.ReviewedID,
		Rating:     updatedReview.Rating,
		Comment:    updatedReview.Comment,
		CreatedAt:  updatedReview.CreatedAt,
	}

	return response, nil
}

// ✅ 5. Hapus Review
func (s *reviewService) DeleteReview(reviewID uint, reviewerID uint) error {
	review, err := s.reviewRepo.GetReviewByID(reviewID)
	if err != nil {
		return errors.New("review not found")
	}

	if review.ReviewerID != reviewerID {
		return errors.New("unauthorized: you can only delete your own reviews")
	}

	return s.reviewRepo.DeleteReview(reviewID, reviewerID) // ✅ Sesuaikan parameter
}

// ✅ 6. Hitung Rata-Rata Rating User
func (s *reviewService) GetAverageRating(userID uint) (*dto.AverageRatingResponse, error) {
	average, count, err := s.reviewRepo.GetAverageRating(userID)
	if err != nil {
		return nil, err
	}

	return &dto.AverageRatingResponse{
		ReviewedID: userID,
		Average:    average,
		TotalCount: count,
	}, nil
}
