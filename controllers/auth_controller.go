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

// RegisterUser menangani pendaftaran user baru
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

	// Validasi tambahan menggunakan validator dari utils
	if err := utils.ValidateStruct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Validation error",
			"errors":  err.Error(),
		})
		return
	}

	// Hash password sebelum disimpan
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Error hashing password",
		})
		return
	}

	// Buat user baru
	user := models.User{
		FullName: request.FullName,
		Email:    request.Email,
		Password: hashedPassword,
		Role:     request.Role,
		Phone:    request.Phone,
	}

	// Simpan user ke database
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Error saving user",
		})
		return
	}

	// Response
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

// LoginUser menangani autentikasi user
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

	// Verifikasi password
	if !utils.CheckPassword(request.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Invalid email or password",
		})
		return
	}

	// Buat token JWT
	expirationTime := time.Now().Add(time.Hour * 24)
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

	// Buat response user
	response := dto.UserResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Role:      user.Role,
		Phone:     &user.Phone,
		AvatarURL: &user.AvatarURL,
		CreatedAt: user.CreatedAt,
	}

	// Response login dengan LoginResponse DTO
	loginResponse := dto.LoginResponse{
		Status:  "success",
		Message: "Login successful",
		Data:    response,
		Token:   tokenString,
	}

	c.JSON(http.StatusOK, loginResponse)
}
