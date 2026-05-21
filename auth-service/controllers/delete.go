package controllers

import (
	"auth/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/logger"
	"github.com/sushantpardhi/shared/response"
	"go.uber.org/zap"
)

func DeleteUser(c *gin.Context) {
	logger.C(c).Info("DeleteUser endpoint hit")
	userId := c.Param("user_id")

	if userId == "" {
		logger.C(c).Error("Missing user_id param")
		response.Error(c, http.StatusBadRequest, "user_id is required", nil)
		return
	}

	err := services.DeleteUserAccount(c, userId)
	if err != nil {
		logger.C(c).Error("Failed to delete user", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "Failed to delete user", err)
		return
	}

	logger.C(c).Info("User deleted successfully", zap.String("user_id", userId))
	response.Success(c, http.StatusOK, "User deleted", gin.H{"user_id": userId})
}
