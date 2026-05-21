package controllers

import (
	"net/http"
	"user/models"
	"user/services"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/logger"
	"github.com/sushantpardhi/shared/response"
	"go.uber.org/zap"
)

// UpdateUser is admin-only: updates a user's profile by role and user_id params.
func UpdateUser(c *gin.Context) {
	role := c.Param("role")
	userId := c.Param("user_id")

	if role == "" || userId == "" {
		response.Error(c, http.StatusBadRequest, "role and user_id are required", nil)
		return
	}

	logger.C(c).Info("Admin updating user profile", zap.String("role", role), zap.String("user_id", userId))

	input, ok := bindUpdateInput(c, role)
	if !ok {
		return
	}

	if err := services.UpdateUserByRole(c, role, userId, input); err != nil {
		logger.C(c).Error("Failed to update user", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "Failed to update user profile", err)
		return
	}

	response.Success(c, http.StatusOK, "User profile updated", gin.H{"role": role, "user_id": userId})
}

// UpdateSelf allows the logged-in user to update their own profile based on their role.
func UpdateSelf(c *gin.Context) {
	userId := c.MustGet("user_id").(string)
	role := c.MustGet("role").(string)

	logger.C(c).Info("User updating self", zap.String("role", role), zap.String("user_id", userId))

	input, ok := bindUpdateInput(c, role)
	if !ok {
		return
	}

	if err := services.UpdateSelfByRole(c, role, userId, input); err != nil {
		logger.C(c).Error("Failed to update self", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "Failed to update profile", err)
		return
	}

	response.Success(c, http.StatusOK, "Profile updated", gin.H{"user_id": userId})
}

// bindUpdateInput binds the request body to the appropriate update model for the given role.
func bindUpdateInput(c *gin.Context, role string) (any, bool) {
	switch role {
	case "admin":
		var input models.UpdateAdmin
		if err := c.ShouldBindJSON(&input); err != nil {
			response.Error(c, http.StatusBadRequest, "Invalid input", err)
			return nil, false
		}
		return input, true
	case "teacher":
		var input models.UpdateTeacher
		if err := c.ShouldBindJSON(&input); err != nil {
			response.Error(c, http.StatusBadRequest, "Invalid input", err)
			return nil, false
		}
		return input, true
	case "student":
		var input models.UpdateStudent
		if err := c.ShouldBindJSON(&input); err != nil {
			response.Error(c, http.StatusBadRequest, "Invalid input", err)
			return nil, false
		}
		return input, true
	case "parent":
		var input models.UpdateParent
		if err := c.ShouldBindJSON(&input); err != nil {
			response.Error(c, http.StatusBadRequest, "Invalid input", err)
			return nil, false
		}
		return input, true
	default:
		response.Error(c, http.StatusBadRequest, "Invalid role", nil)
		return nil, false
	}
}

