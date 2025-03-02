package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/habbazettt/jobseek-go/dto"
	"github.com/habbazettt/jobseek-go/services"
	"github.com/habbazettt/jobseek-go/utils"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService}
}

// ✅ Get All Users (Hanya Admin)
func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(ctx, http.StatusOK, "Users retrieved successfully", users)
}

// ✅ Get Current User
func (c *UserController) GetCurrentUser(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")
	user, err := c.userService.GetUserByID(userID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "User not found")
		return
	}
	utils.SuccessResponse(ctx, http.StatusOK, "User retrieved successfully", user)
}

// ✅ Get User by ID (Hanya Admin)
func (c *UserController) GetUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := c.userService.GetUserByID(uint(id))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "User not found")
		return
	}
	utils.SuccessResponse(ctx, http.StatusOK, "User retrieved successfully", user)
}

// ✅ Update User
func (c *UserController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var request dto.UpdateUserRequest
	if err := ctx.ShouldBind(&request); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// ✅ Ambil user_id & role dari token
	userID, _ := ctx.Get("user_id")
	role, _ := ctx.Get("role")

	// ✅ Hanya pemilik akun atau admin yang bisa update
	if userID.(uint) != uint(id) && role.(string) != "admin" {
		utils.ErrorResponse(ctx, http.StatusForbidden, "Unauthorized to update this user")
		return
	}

	// ✅ Ambil file dari form-data jika ada
	file, _ := ctx.FormFile("photo")

	user, err := c.userService.UpdateUser(uint(id), request, file)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "User updated successfully", user)
}

// ✅ Delete User
func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// ✅ Ambil user_id & role dari token
	userID, _ := ctx.Get("user_id")
	role, _ := ctx.Get("role")

	// ✅ Hanya pemilik akun atau admin yang bisa delete
	if userID.(uint) != uint(id) && role.(string) != "admin" {
		utils.ErrorResponse(ctx, http.StatusForbidden, "Unauthorized to delete this user")
		return
	}

	err = c.userService.DeleteUser(uint(id))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "User deleted successfully", nil)
}
