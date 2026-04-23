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

func UpdateUser(c *gin.Context) {
	user_id := c.Param("user_id")
	role := c.Param("role")
	logger.Info("Updating user with ID: ", zap.String("user_id", user_id), zap.String("role", role))

	switch role {
	case "admin":
		logger.Info("Updating admin user with ID: ", zap.String("user_id", user_id))
		UpdateAdmin(c, user_id)
		response.Success(c, http.StatusOK, "Admin user updated successfully", gin.H{"user_id": user_id})
	case "student":
		logger.Info("Updating student user with ID: ", zap.String("user_id", user_id))
		UpdateStudent(c, user_id)
		response.Success(c, http.StatusOK, "Student user updated successfully", gin.H{"user_id": user_id})
	case "teacher":
		logger.Info("Updating teacher user with ID: ", zap.String("user_id", user_id))
		UpdateTeacher(c, user_id)
		response.Success(c, http.StatusOK, "Teacher user updated successfully", gin.H{"user_id": user_id})
	case "parent":
		logger.Info("Updating parent user with ID: ", zap.String("user_id", user_id))
		UpdateParent(c, user_id)
		response.Success(c, http.StatusOK, "Parent user updated successfully", gin.H{"user_id": user_id})
	default:
		logger.Warn("Invalid role provided for user update: ", zap.String("role", role))
		response.Error(c, http.StatusBadRequest, "Invalid role provided", nil)
	}
}
