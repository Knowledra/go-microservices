package services

import (
	"auth/models"

	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/jwtToken"

	"github.com/sushantpardhi/shared/logger"

	"context"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx context.Context, email, password string) (string, models.User, error) {
	var user models.User

	// DB query (include soft-deleted records)
	logger.Ctx(ctx).Info("Checking if user exists with email", zap.String("email", email))
	if err := db.DB.Unscoped().Where("email = ?", email).First(&user).Error; err != nil {
		return "", user, err
	}
	logger.Ctx(ctx).Info("User found", zap.String("user_id", user.ID.String()))

	// Check if user is deleted, if yes then restore them
	logger.Ctx(ctx).Info("Checking if user is deleted", zap.String("user_id", user.ID.String()))
	if user.IsDeleted {
		logger.Ctx(ctx).Info("User is deleted, restoring user", zap.String("user_id", user.ID.String()))
		if err := db.DB.Unscoped().Model(&user).Update("is_deleted", false).Update("deleted_at", nil).Error; err != nil {
			logger.Ctx(ctx).Error("Failed to restore user", zap.String("user_id", user.ID.String()), zap.Error(err))
			return "", user, err
		}
		logger.Ctx(ctx).Info("User restored successfully", zap.String("user_id", user.ID.String()))
	}

	// Password check
	logger.Ctx(ctx).Info("Comparing password for user", zap.String("user_id", user.ID.String()))
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", user, err
	}
	logger.Ctx(ctx).Info("Password compared successfully", zap.String("user_id", user.ID.String()))

	// Token generation
	logger.Ctx(ctx).Info("Generating token for user", zap.String("user_id", user.ID.String()))
	token, err := jwtToken.GenerateToken(user.ID.String(), user.Email, string(user.Role))
	if err != nil {
		return "", user, err
	}
	logger.Ctx(ctx).Info("Token generated successfully", zap.String("user_id", user.ID.String()))

	return token, user, nil
}
