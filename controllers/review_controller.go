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

// CreateReview godoc
// @Summary      Create a new review
// @Description  Create a new review. Only authenticated reviewer can create a review.
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Param        request  body     dto.CreateReviewRequest  true  "Request body"
// @Success      201      {object} dto.ReviewResponse        "Review created successfully"
// @Failure      400      {object} utils.ErrorResponseSwagger       "Bad request"
// @Failure      500      {object} utils.ErrorResponseSwagger       "Internal server error"
// @Router       /reviews [post]
// @Security     BearerAuth
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


// GetReviewsByUser retrieves all reviews received by a specific user.
// @Summary      Get Reviews By User
// @Description  Retrieve all reviews for a particular user based on the user ID provided in the path parameter.
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Param        user_id path int true "User ID"
// @Success      200 {array} dto.ReviewResponse "Reviews retrieved successfully"
// @Failure      400 {object} utils.ErrorResponseSwagger "Invalid user ID"
// @Failure      500 {object} utils.ErrorResponseSwagger "Internal server error"
// @Router       /reviews/{user_id} [get]
// @Security     BearerAuth

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


// GetMyReviews retrieves all reviews submitted by the currently authenticated user.
// @Summary      Get My Reviews
// @Description  Retrieve all reviews submitted by the currently authenticated user.
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Success      200 {array} dto.ReviewResponse "Reviews retrieved successfully"
// @Failure      401 {object} utils.ErrorResponseSwagger "Unauthorized: No user ID found in token"
// @Failure      500 {object} utils.ErrorResponseSwagger "Internal server error"
// @Router       /reviews/me [get]
// @Security     BearerAuth
func (c *ReviewController) GetMyReviews(ctx *gin.Context) {
	reviewerID, _ := ctx.Get("user_id")

	reviews, err := c.reviewService.GetReviewsByReviewerID(reviewerID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Your reviews retrieved successfully", reviews)
}


// UpdateReview updates an existing review. Only the reviewer who submitted the
// review can make the update. The review ID is provided as a path parameter.
// @Summary      Update Review
// @Description  Update an existing review. Only the reviewer who submitted the
// review can make the update.
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Param        review_id path int true "Review ID"
// @Param        request body dto.UpdateReviewRequest true "Review details"
// @Success      200 {object} dto.ReviewResponse "Review updated successfully"
// @Failure      400 {object} utils.ErrorResponseSwagger "Invalid review ID"
// @Failure      401 {object} utils.ErrorResponseSwagger "Unauthorized: You can only update your own reviews"
// @Failure      500 {object} utils.ErrorResponseSwagger "Internal server error"
// @Router       /reviews/{review_id} [put]
// @Security     BearerAuth
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


// DeleteReview deletes a review. Only the reviewer who submitted the review
// can delete it. The review ID is provided as a path parameter.
// @Summary      Delete Review
// @Description  Delete a review. Only the reviewer who submitted the review
// can delete it.
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Param        review_id path int true "Review ID"
// @Success      200 {object} dto.ReviewResponse "Review deleted successfully"
// @Failure      400 {object} utils.ErrorResponseSwagger "Invalid review ID"
// @Failure      401 {object} utils.ErrorResponseSwagger "Unauthorized: You can only delete your own reviews"
// @Failure      500 {object} utils.ErrorResponseSwagger "Internal server error"
// @Router       /reviews/{review_id} [delete]
// @Security     BearerAuth
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


// GetAverageRating retrieves the average rating of a user.
// @Summary      Get Average Rating
// @Description  Retrieve the average rating of a user.
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Param        user_id path int true "User ID"
// @Success      200 {object} dto.AverageRatingResponse "Average rating retrieved successfully"
// @Failure      400 {object} utils.ErrorResponseSwagger "Invalid user ID"
// @Failure      500 {object} utils.ErrorResponseSwagger "Internal server error"
// @Router       /reviews/average/{user_id} [get]
// @Security     BearerAuth
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
