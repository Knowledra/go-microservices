package controllers

import (
	"auth/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/logger"
	"github.com/sushantpardhi/shared/response"
	"go.uber.org/zap"
)

type resetPasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

// This function will take new password in body and then check the strength of the password and then update the password in database after hashing it and then send an email to the user informing them about the password change.
// It will also invalidate all existing tokens for the user so that they have to login again with the new password.
func ResetPassword(c *gin.Context) {
	logger.C(c).Info("ResetPassword endpoint hit")

	var input resetPasswordRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.C(c).Error("Invalid input", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "invalid input", err)
		return
	}
	logger.C(c).Info("JSON input bound successfully")

	if input.NewPassword != input.ConfirmPassword {
		response.Error(c, http.StatusBadRequest, "new_password and confirm_password do not match", nil)
		return
	}

	userID, ok := c.Get("user_id")
	if !ok {
		response.Error(c, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	err := services.ResetPasswordForUser(c, userID.(string), input.CurrentPassword, input.NewPassword)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrUserNotFound):
			response.Error(c, http.StatusNotFound, "user not found", err)
		case errors.Is(err, services.ErrCurrentPasswordIncorrect), errors.Is(err, services.ErrNewPasswordSameAsCurrent):
			response.Error(c, http.StatusBadRequest, err.Error(), err)
		default:
			response.Error(c, http.StatusBadRequest, err.Error(), err)
		}
		return
	}

	// Invalidate current session cookie so user must login with new credentials.
	c.SetCookie("access_token", "", -1, "/", "", true, true)

	response.Success(c, http.StatusOK, "password reset successful, please login again", nil)

}
