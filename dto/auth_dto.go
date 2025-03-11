package dto

import (
	"time"

	"github.com/habbazettt/jobseek-go/utils"
)

type RegisterRequest struct {
	FullName  string `json:"full_name" binding:"required,min=3,max=100"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8,max=50" validate:"strong_password"`
	Role      string `json:"role" binding:"required,oneof=admin freelancer perusahaan"`
	Phone     string `json:"phone,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Phone     *string   `json:"phone,omitempty"`
	AvatarURL *string   `json:"avatar_url,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginResponse struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Data    UserResponse `json:"data"`
	Token   string       `json:"token"`
}

func init() {
	utils.InitValidator()
}
