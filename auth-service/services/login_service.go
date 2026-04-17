package services

import (
	"auth/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sushantpardhi/shared/db"

	"github.com/sushantpardhi/shared/logger"

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
	token, err := generateToken(user.ID.String(), user.Email, string(user.Role))
	if err != nil {
		return "", user, err
	}
	logger.Info("Token generated successfully", zap.String("user_id", user.ID.String()))

	return token, user, nil
}

type claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func getSecretKey() []byte {
	return []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
}

func generateToken(userID, email, role string) (string, error) {
	// Create claims
	claims := claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "eduplat",
		},
	}

	// Sign token
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(getSecretKey())
}
