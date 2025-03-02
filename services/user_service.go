package services

import (
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/habbazettt/jobseek-go/config"
	"github.com/habbazettt/jobseek-go/dto"
	"github.com/habbazettt/jobseek-go/repositories"
)

type UserService interface {
	GetAllUsers() ([]dto.UserResponse, error)
	GetUserByID(id uint) (*dto.UserResponse, error)
	UpdateUser(id uint, request dto.UpdateUserRequest, file *multipart.FileHeader) (*dto.UserResponse, error)
	DeleteUser(id uint) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) GetAllUsers() ([]dto.UserResponse, error) {
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var responses []dto.UserResponse
	for _, user := range users {
		responses = append(responses, dto.UserResponse{
			ID:        user.ID,
			FullName:  user.FullName,
			Email:     user.Email,
			Phone:     &user.Phone,
			AvatarURL: &user.AvatarURL,
			Role:      user.Role,
		})
	}
	return responses, nil
}

func (s *userService) GetUserByID(id uint) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	response := dto.UserResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Phone:     &user.Phone,
		AvatarURL: &user.AvatarURL,
		Role:      user.Role,
	}
	return &response, nil
}

func (s *userService) UpdateUser(id uint, request dto.UpdateUserRequest, file *multipart.FileHeader) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if request.FullName != "" {
		user.FullName = request.FullName
	}
	if request.Phone != "" {
		user.Phone = request.Phone
	}

	if file != nil {
		src, err := file.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open file: %v", err)
		}
		defer src.Close()

		imageURL, err := config.UploadImage(src)
		if err != nil {
			return nil, fmt.Errorf("failed to upload image: %v", err)
		}
		user.AvatarURL = imageURL
	}

	err = s.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	response := dto.UserResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Role:      user.Role,
		Phone:     &user.Phone,
		AvatarURL: &user.AvatarURL,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return &response, nil
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.DeleteUser(id)
}
