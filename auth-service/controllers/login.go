package controllers

import (
	"auth/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/logger"
	"github.com/sushantpardhi/shared/models"
	"github.com/sushantpardhi/shared/response"
	"go.uber.org/zap"
)

func Login(c *gin.Context) {
	logger.Info("Login endpoint hit")
	var input models.Login

	logger.Info("Binding JSON input for Login")
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Error("Invalid input", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "Invalid input", err)
		return
	}
	logger.Info("JSON input bound successfully")

	token, user, err := services.Login(input.Email, input.Password)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Invalid credentials", err)
		return
	}

	logger.Info("Setting cookie for user", zap.String("user_id", user.ID.String()))
	c.SetCookie("access_token", token, 3600*24, "/", "", true, true)
	logger.Info("Cookie set successfully")

	logger.Info("Login successful", zap.String("user_id", user.ID.String()))
	response.Success(c, http.StatusOK, "Login successful", gin.H{"user": user})
}
