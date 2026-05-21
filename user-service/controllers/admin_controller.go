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

func CreateAdmin(C *gin.Context) {
	logger.C(C).Info("Creating Admin Profile")

	var input models.Admin
	if err := C.ShouldBindJSON(&input); err != nil {
		logger.C(C).Error("Invalid Input", zap.Error(err))
		response.Error(C, http.StatusBadRequest, "Invalid Input", err)
		return
	}

	admin, err := services.CreateAdminProfile(C, input)
	if err != nil {
		response.Error(C, http.StatusInternalServerError, "Failed to create admin", err)
		return
	}

	logger.C(C).Info("Admin Profile created successfully", zap.String("user_id", admin.ID.String()))
	response.Success(C, http.StatusCreated, "Admin created", gin.H{"admin": admin})
}

func CreateSuperAdmin(C *gin.Context) {
	logger.C(C).Info("Creating Super Admin Profile")

	var input models.SuperAdmin
	if err := C.ShouldBindJSON(&input); err != nil {
		logger.C(C).Error("Invalid Input", zap.Error(err))
		response.Error(C, http.StatusBadRequest, "Invalid Input", err)
		return
	}

	superAdmin, err := services.CreateSuperAdminProfile(C, input)
	if err != nil {
		response.Error(C, http.StatusInternalServerError, "Failed to create super admin", err)
		return
	}

	logger.C(C).Info("Super Admin Profile created successfully", zap.String("user_id", superAdmin.ID.String()))
	response.Success(C, http.StatusCreated, "Super Admin created", gin.H{"super_admin": superAdmin})
}

func DeleteUserProfile(c *gin.Context) {
	role := c.Param("role")
	userId := c.Param("user_id")

	if role == "" || userId == "" {
		logger.C(c).Error("Missing role or user_id for profile deletion")
		response.Error(c, http.StatusBadRequest, "role and user_id are required", nil)
		return
	}

	if err := services.DeleteUserProfile(c, role, userId); err != nil {
		logger.C(c).Error("Failed to delete user profile", zap.String("role", role), zap.String("user_id", userId), zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "Failed to delete user profile", err)
		return
	}

	response.Success(c, http.StatusOK, "User profile deleted", gin.H{"role": role, "user_id": userId})
}

func UpdateAdmin(c *gin.Context, userId string) {
	logger.C(c).Info("UpdateAdmin endpoint hit")

	// Bind json
	var input models.UpdateAdmin
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.C(c).Error("Invalid Input", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "Invalid Input", err)
		return
	}

	// Update admin
	err := services.UpdateAdminUser(c, userId, input)
	if err != nil {
		logger.C(c).Error("Failed to update admin", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "Failed to update admin", err)
		return
	}

	logger.C(c).Info("Admin updated successfully", zap.String("user_id", userId))
}
