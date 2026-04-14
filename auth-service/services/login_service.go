package services

import (
	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/jwtToken"
	"github.com/sushantpardhi/shared/logger"
	"github.com/sushantpardhi/shared/models"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func Login(email, password string) (string, models.User, error) {
	var user models.User

	// DB query
	logger.Info("Checking if user exists with email", zap.String("email", email))
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", user, err
	}
	logger.Info("User found", zap.String("user_id", user.ID.String()))

	// Password check
	logger.Info("Comparing password for user", zap.String("user_id", user.ID.String()))
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", user, err
	}
	logger.Info("Password compared successfully", zap.String("user_id", user.ID.String()))

	// Token generation
	logger.Info("Generating token for user", zap.String("user_id", user.ID.String()))
	token, err := jwtToken.GenerateToken(user.ID.String(), user.Email, string(user.Role))
	if err != nil {
		return "", user, err
	}
	logger.Info("Token generated successfully", zap.String("user_id", user.ID.String()))

	return token, user, nil
}
