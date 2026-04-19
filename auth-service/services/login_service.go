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

func Login(ctx context.Context, email, password string) (string, models.AuthUser, error) {
	var user models.AuthUser

	// DB query
	logger.Ctx(ctx).Info("Checking if user exists with email", zap.String("email", email))
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", user, err
	}
	logger.Ctx(ctx).Info("User found", zap.String("user_id", user.ID.String()))

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
