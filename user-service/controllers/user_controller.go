package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/logger"
	"github.com/sushantpardhi/shared/response"
	"go.uber.org/zap"
)

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
