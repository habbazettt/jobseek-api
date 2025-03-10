package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/habbazettt/jobseek-go/dto"
	"github.com/habbazettt/jobseek-go/services"
	"github.com/habbazettt/jobseek-go/utils"
)

type ReviewController struct {
	reviewService services.ReviewService
}

func NewReviewController(reviewService services.ReviewService) *ReviewController {
	return &ReviewController{reviewService}
}

// ✅ 1. Tambah Review
func (c *ReviewController) CreateReview(ctx *gin.Context) {
	var request dto.CreateReviewRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	reviewerID, _ := ctx.Get("user_id")
	review, err := c.reviewService.CreateReview(request, reviewerID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Review submitted successfully", review)
}

// ✅ 2. Ambil Review Berdasarkan User ID
func (c *ReviewController) GetReviewsByUser(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID")
		return
	}

	reviews, err := c.reviewService.GetReviewsByUserID(uint(userID))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Reviews retrieved successfully", reviews)
}

// ✅ 3. Ambil Review yang Diberikan oleh Reviewer (Get My Reviews)
func (c *ReviewController) GetMyReviews(ctx *gin.Context) {
	reviewerID, _ := ctx.Get("user_id")

	reviews, err := c.reviewService.GetReviewsByReviewerID(reviewerID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Your reviews retrieved successfully", reviews)
}

// ✅ 4. Update Review
func (c *ReviewController) UpdateReview(ctx *gin.Context) {
	reviewID, err := strconv.Atoi(ctx.Param("review_id"))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid review ID")
		return
	}

	var request dto.UpdateReviewRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userID, _ := ctx.Get("user_id")

	updatedReview, err := c.reviewService.UpdateReview(uint(reviewID), request.Rating, request.Comment, userID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Review updated successfully", updatedReview)
}

// ✅ 5. Hapus Review
func (c *ReviewController) DeleteReview(ctx *gin.Context) {
	reviewID, err := strconv.Atoi(ctx.Param("review_id"))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid review ID")
		return
	}

	reviewerID, _ := ctx.Get("user_id")

	err = c.reviewService.DeleteReview(uint(reviewID), reviewerID.(uint)) // ✅ Sesuaikan parameter
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Review deleted successfully", nil)
}

// ✅ 6. Ambil Rata-Rata Rating User
func (c *ReviewController) GetAverageRating(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID")
		return
	}

	average, err := c.reviewService.GetAverageRating(uint(userID))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Average rating retrieved successfully", average)
}
