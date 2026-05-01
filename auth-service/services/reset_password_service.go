package services

import (
	"auth/models"
	"context"
	"errors"

	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/logger"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserNotFound = errors.New("user not found")
var ErrCurrentPasswordIncorrect = errors.New("current password is incorrect")
var ErrNewPasswordSameAsCurrent = errors.New("new password must be different from current password")

func ResetPasswordForUser(c context.Context, userID string, currentPassword string, newPassword string) error {
	if err := ValidatePasswordStrength(newPassword); err != nil {
		return err
	}

	user, err := GetUserByID(c, userID)
	if err != nil {
		logger.C(c).Error("Failed to load user for password reset", zap.String("user_id", userID), zap.Error(err))
		return ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(currentPassword)); err != nil {
		return ErrCurrentPasswordIncorrect
	}

	if currentPassword == newPassword {
		return ErrNewPasswordSameAsCurrent
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		logger.C(c).Error("Failed to hash new password", zap.String("user_id", userID), zap.Error(err))
		return errors.New("failed to process new password")
	}

	if err := db.DB.Model(&models.User{}).Where("id = ? AND is_deleted = false", userID).Update("password", string(hashedPassword)).Error; err != nil {
		logger.C(c).Error("Failed to update password", zap.String("user_id", userID), zap.Error(err))
		return errors.New("failed to update password")
	}

	logger.C(c).Info("Password reset successful", zap.String("user_id", userID))
	return nil
}
