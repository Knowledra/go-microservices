package services

import (
	"errors"
	"os"

	"auth/models"

	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/logger"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateAdminUser(input models.RegisterAdmin) (models.User, error) {
	// Password check
	adminSecret := os.Getenv("ADMIN_PASS")

	logger.Info("Validating admin credentials for CreateAdmin")
	if input.AdminPassword == "" || input.AdminPassword != adminSecret {
		logger.Error("Invalid admin credentials")
		return models.User{}, errors.New("invalid admin credentials")
	}
	logger.Info("Admin credentials validated successfully")

	// Check if user already exists
	logger.Info("Checking if user already exists with email", zap.String("email", input.Email))
	var existingUser models.User
	err := db.DB.Where("email = ?", input.Email).First(&existingUser).Error

	if err == nil {
		// user found → already exists
		logger.Error("User already exists with email", zap.String("email", input.Email))
		return models.User{}, errors.New("user already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// real DB error
		logger.Error("Database error while checking existing user", zap.Error(err))
		return models.User{}, errors.New("database error")
	}
	logger.Info("No existing user found with email, proceeding to hash admin password", zap.String("email", input.Email))

	// Hash password
	logger.Info("Hashing password for new admin user")
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(input.Password), 10)

	if hashErr != nil {
		logger.Error("Password hashing failed", zap.Error(hashErr))
		return models.User{}, errors.New("password hashing failed")
	}
	logger.Info("Password hashed successfully for new admin user")

	// Create user
	logger.Info("Creating new admin user in database", zap.String("email", input.Email))
	newUser := models.User{
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     "admin",
	}

	// Save user in User table
	logger.Info("Saving admin in User table", zap.String("email", newUser.Email))
	if err := db.DB.Create(&newUser).Error; err != nil {
		logger.Error("Failed to create admin in User table", zap.Error(err))
		return models.User{}, errors.New("failed to create admin in User table")
	}
	logger.Info("Admin created successfully in User table", zap.String("user_id", newUser.ID.String()))

	// TODO: Call the User-Service API and add Admin data to Admin table in users db

	return newUser, nil
}

func CreateAdminProfile(input models.RegisterAdmin) (models.User, error) {
	return models.User{}, nil
}
