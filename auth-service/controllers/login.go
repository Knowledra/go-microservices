package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/jwtToken"
	"github.com/sushantpardhi/shared/logger"
	"github.com/sushantpardhi/shared/models"
	"github.com/sushantpardhi/shared/response"
	"go.uber.org/zap"

	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	// Bind json
	var input models.Login

	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Error("Failed to bind login input", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Find user by email
	var user models.User
	if err := db.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		logger.Error("User not found with email", zap.String("email", input.Email), zap.Error(err))
		response.Error(c, http.StatusUnauthorized, "Invalid credentials", err)
		return
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		logger.Error("Password mismatch for user", zap.String("email", input.Email), zap.Error(err))
		response.Error(c, http.StatusUnauthorized, "Invalid credentials", err)
		return
	}

	// Generate access token
	access_token, err := jwtToken.GenerateToken(user.ID.String(), user.Email, string(user.Role))
	if err != nil {
		logger.Error("Failed to generate access token", zap.String("user_id", user.ID.String()), zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "Failed to generate token", err)
		return
	}

	// Set cookie
	c.SetCookie(
		"access_token",
		access_token,
		3600*24,
		"/",
		"",
		true,
		true,
	)

	logger.Info("Login successful", zap.String("user_id", user.ID.String()))
	response.Success(c, http.StatusOK, "Login successful", gin.H{"user": user})
}
