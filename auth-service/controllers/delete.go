package controllers

import (
	"auth/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/logger"
	"github.com/sushantpardhi/shared/response"
	"go.uber.org/zap"
)

func DeleteSelf(c *gin.Context) {
	logger.C(c).Info("DeleteSelf endpoint hit")
	userId := c.MustGet("user_id").(string)

	err := services.DeleteUserAccount(c, userId)
	if err != nil {
		logger.C(c).Error("Failed to delete user", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "Failed to delete user", err)
		return
	}

	logger.Info("Removing cookies")
	c.SetCookie("access_token", "", -1, "/", "", true, true)
	logger.Info("Cookies removed successfully")

	logger.C(c).Info("User deleted successfully", zap.String("user_id", userId))
	response.Success(c, http.StatusOK, "User deleted", gin.H{"user": userId})
}
