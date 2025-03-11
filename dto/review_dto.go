package dto

import "time"

type CreateReviewRequest struct {
	ReviewedID uint    `json:"reviewed_id" binding:"required"` // ID User yang akan diberikan review
	Rating     float64 `json:"rating" binding:"required,min=1,max=5"`
	Comment    string  `json:"comment" binding:"required"`
}

type UpdateReviewRequest struct {
	Rating  float64 `json:"rating" binding:"required,min=1,max=5"`
	Comment string  `json:"comment" binding:"required"`
}

type ReviewResponse struct {
	ID         uint      `json:"id"`
	ReviewerID uint      `json:"reviewer_id"`
	ReviewedID uint      `json:"reviewed_id"`
	Rating     float64   `json:"rating"`
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `json:"created_at"`
}

type AverageRatingResponse struct {
	ReviewedID uint    `json:"reviewed_id"`
	Average    float64 `json:"average_rating"`
	TotalCount int     `json:"total_reviews"`
}
