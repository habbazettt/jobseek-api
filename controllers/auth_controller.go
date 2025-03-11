package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/habbazettt/jobseek-go/config"
	"github.com/habbazettt/jobseek-go/dto"
	"github.com/habbazettt/jobseek-go/models"
	"github.com/habbazettt/jobseek-go/utils"
)

// @Summary Register new user
// @Description Register new user, only available for guest
// @Tags auth
// @Accept  json
// @Produce  json
// @Param request body dto.RegisterRequest true "Register request"
// @Success 201 {object} dto.UserResponse "User registered successfully"
// @Failure 400 {object} utils.ErrorResponseSwagger "Invalid request format"
// @Failure 500 {object} utils.ErrorResponseSwagger "Error saving user"
// @Router /auth/register [post]
func RegisterUser(c *gin.Context) {
	var request dto.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request format",
			"errors":  err.Error(),
		})
		return
	}

	if err := utils.ValidateStruct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Validation error",
			"errors":  err.Error(),
		})
		return
	}

	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Error hashing password",
		})
		return
	}

	user := models.User{
		FullName: request.FullName,
		Email:    request.Email,
		Password: hashedPassword,
		Role:     request.Role,
		Phone:    request.Phone,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Error saving user",
		})
		return
	}

	response := dto.UserResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Role:      user.Role,
		Phone:     &user.Phone,
		AvatarURL: &user.AvatarURL,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "User registered successfully",
		"data":    response,
	})
}

// @Summary Login user
// @Description Login user, only available for guest
// @Tags auth
// @Accept  json
// @Produce  json
// @Param request body dto.LoginRequest true "Login request"
// @Success 200 {object} dto.LoginResponse "User logged in successfully"
// @Failure 400 {object} utils.ErrorResponseSwagger "Invalid request"
// @Failure 401 {object} utils.ErrorResponseSwagger "Invalid email or password"
// @Failure 500 {object} utils.ErrorResponseSwagger "Failed to login user"
// @Router /auth/login [post]
func LoginUser(c *gin.Context) {
	var request dto.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request",
			"errors":  err.Error(),
		})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Invalid email or password",
		})
		return
	}

	if !utils.CheckPassword(request.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Invalid email or password",
		})
		return
	}

	expirationTime := time.Now().Add(time.Hour * 24) // 24 hours
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"email":   user.Email,
		"exp":     expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Error generating token",
		})
		return
	}

	response := dto.UserResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Role:      user.Role,
		Phone:     &user.Phone,
		AvatarURL: &user.AvatarURL,
		CreatedAt: user.CreatedAt,
	}

	loginResponse := dto.LoginResponse{
		Status:  "success",
		Message: "Login successful",
		Data:    response,
		Token:   tokenString,
	}

	c.JSON(http.StatusOK, loginResponse)
}
