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

// @Summary      Get All Users
// @Description  Get All Users
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   dto.UserResponse
// @Failure      500  {object}  map[string]interface{}
// @Router       /users [get]
func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(ctx, http.StatusOK, "Users retrieved successfully", users)
}

// @Summary      Get Current User
// @Description  Get Current User
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  dto.UserResponse
// @Failure      404  {object}  map[string]interface{}
// @Router       /users/me [get]
func (c *UserController) GetCurrentUser(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")
	user, err := c.userService.GetUserByID(userID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "User not found")
		return
	}
	utils.SuccessResponse(ctx, http.StatusOK, "User retrieved successfully", user)
}

// @Summary      Get User By ID
// @Description  Get User By ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  dto.UserResponse
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Router       /users/{id} [get]
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

// @Summary      Update User
// @Description  Update user details based on the provided user ID. Only the account owner or an admin
//
//	can perform this update. Accepts a JSON payload for user details and optionally a
//	file upload for updating the user's avatar.
//
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Param        photo formData file false "User Avatar"
// @Security     BearerAuth
// @Success      200  {object}  dto.UserResponse
// @Failure      400  {object}  map[string]interface{}
// @Failure      403  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /users/{id} [put]
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

	userID, _ := ctx.Get("user_id")
	role, _ := ctx.Get("role")

	if userID.(uint) != uint(id) && role.(string) != "admin" {
		utils.ErrorResponse(ctx, http.StatusForbidden, "Unauthorized to update this user")
		return
	}

	file, _ := ctx.FormFile("photo")

	user, err := c.userService.UpdateUser(uint(id), request, file)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "User updated successfully", user)
}

// @Summary      Delete User
// @Description  Delete a user based on the provided user ID. Only the account owner or an admin
//
//	can perform this deletion. The function verifies the user identity and role
//	from the token to ensure authorization.
//
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      403  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /users/{id} [delete]
func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID")
		return
	}

	userID, _ := ctx.Get("user_id")
	role, _ := ctx.Get("role")

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
