package services

import (
	"context"
	"encoding/json"
	"errors"
	"os"

	"auth/models"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/callAPI"
	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/logger"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateAdminUser(ctx context.Context, input models.RegisterAdmin) (models.User, json.RawMessage, error) {
	// Password check
	adminSecret := os.Getenv("ADMIN_PASS")

	logger.Ctx(ctx).Info("Validating admin credentials for CreateAdmin")
	if input.AdminPassword == "" || input.AdminPassword != adminSecret {
		logger.Ctx(ctx).Error("Invalid admin credentials")
		return models.User{}, nil, errors.New("invalid admin credentials")
	}
	logger.Ctx(ctx).Info("Admin credentials validated successfully")

	// Check if user already exists
	logger.Ctx(ctx).Info("Checking if user already exists with email", zap.String("email", input.Email))
	var existingUser models.User
	err := db.DB.Where("email = ?", input.Email).First(&existingUser).Error

	if err == nil {
		// user found → already exists
		logger.Ctx(ctx).Error("User already exists with email", zap.String("email", input.Email))
		return models.User{}, nil, errors.New("user already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// real DB error
		logger.Ctx(ctx).Error("Database error while checking existing user", zap.Error(err))
		return models.User{}, nil, errors.New("database error")
	}
	logger.Ctx(ctx).Info("No existing user found with email, proceeding to hash admin password", zap.String("email", input.Email))

	// Hash password
	logger.Ctx(ctx).Info("Hashing password for new admin user")
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(input.Password), 10)

	if hashErr != nil {
		logger.Ctx(ctx).Error("Password hashing failed", zap.Error(hashErr))
		return models.User{}, nil, errors.New("password hashing failed")
	}
	logger.Ctx(ctx).Info("Password hashed successfully for new admin user")

	// Create user
	logger.Ctx(ctx).Info("Creating new admin user in database", zap.String("email", input.Email))
	newUser := models.User{
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     "admin",
	}

	// Save user in User table
	logger.Ctx(ctx).Info("Saving admin in User table", zap.String("email", newUser.Email))
	if err := db.DB.Create(&newUser).Error; err != nil {
		logger.Ctx(ctx).Error("Failed to create admin in User table", zap.Error(err))
		return models.User{}, nil, errors.New("failed to create admin in User table")
	}
	logger.Ctx(ctx).Info("Admin created successfully in User table", zap.String("user_id", newUser.ID.String()))

	logger.Ctx(ctx).Info("Calling User-Service API to create admin in Admin table", zap.String("user_id", newUser.ID.String()))
	adminBytes, err := callAPI.CallAPI(ctx, "POST", "http://user-service:8002/api/v1/profile/create/admin", gin.H{
		"user_id":   newUser.ID,
		"name":      input.Name,
		"last_name": input.LastName,
	})

	if err != nil {
		logger.Ctx(ctx).Error("Failed to create admin in Admin table", zap.Error(err))
		return models.User{}, nil, errors.New("failed to create admin in Admin table")
	}
	logger.Ctx(ctx).Info("Admin created successfully in Admin table", zap.String("user_id", newUser.ID.String()))

	var apiResponse struct {
		Data struct {
			Admin json.RawMessage `json:"admin"`
		} `json:"data"`
	}

	var adminData json.RawMessage = adminBytes
	if err := json.Unmarshal(adminBytes, &apiResponse); err == nil && len(apiResponse.Data.Admin) > 0 {
		adminData = apiResponse.Data.Admin
	}

	return newUser, adminData, nil
}
