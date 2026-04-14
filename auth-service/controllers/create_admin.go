package controllers

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/logger"
	"github.com/sushantpardhi/shared/models"
	"github.com/sushantpardhi/shared/response"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateAdmin(c *gin.Context) {
	logger.Info("CreateAdmin endpoint hit")
	// Bind json
	var input models.RegisterAdmin

	logger.Info("Binding JSON input for CreateAdmin")
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Error("Invalid Input", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "Invalid Input", err)
		return
	}
	logger.Info("JSON input bound successfully")

	// Password check
	adminSecret := os.Getenv("ADMIN_PASS")

	logger.Info("Validating admin credentials for CreateAdmin")
	if input.AdminPassword == "" || input.AdminPassword != adminSecret {
		logger.Error("Invalid admin credentials")
		response.Error(c, http.StatusUnauthorized, "Invalid admin credentials", nil)
		return
	}
	logger.Info("Admin credentials validated successfully")

	// Check if user already exists
	logger.Info("Checking if user already exists with email", zap.String("email", input.Email))
	var existingUser models.User
	err := db.DB.Where("email = ?", input.Email).First(&existingUser).Error

	if err == nil {
		// user found → already exists
		logger.Error("User already exists with email", zap.String("email", input.Email))
		response.Error(c, http.StatusBadRequest, "User already exists", nil)
		return
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// real DB error
		logger.Error("Database error while checking existing user", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "Database error", nil)
		return
	}
	logger.Info("No existing user found with email, proceeding to hash admin password", zap.String("email", input.Email))

	// Hash password
	logger.Info("Hashing password for new admin user")
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(input.Password), 10)

	if hashErr != nil {
		logger.Error("Password hashing failed", zap.Error(hashErr))
		response.Error(c, http.StatusInternalServerError, "Password hashing failed", nil)
		return
	}
	logger.Info("Password hashed successfully for new admin user")

	// Create user
	logger.Info("Creating new admin user in database", zap.String("email", input.Email))
	newUser := models.User{
		Name:     input.Name,
		LastName: input.LastName,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     "admin",
	}

	logger.Info("Beginning database transaction to create admin user and admin record")
	err = db.DB.Transaction(func(tx *gorm.DB) error {
		logger.Info("Creating user record for admin", zap.String("email", newUser.Email))
		if err := tx.Create(&newUser).Error; err != nil {
			logger.Error("Failed to create user record for admin", zap.Error(err))
			return err
		}

		logger.Info("Creating admin record linked to user", zap.String("user_id", newUser.ID.String()))
		admin := models.Admin{
			UserID: newUser.ID,
			// User:   newUser,
		}

		logger.Info("Saving admin record to database", zap.String("user_id", newUser.ID.String()))
		if err := tx.Create(&admin).Error; err != nil {
			logger.Error("Failed to create admin record", zap.Error(err))
			return err
		}

		return nil
	})
	logger.Info("Database transaction completed for creating admin user and admin record")

	if err != nil {
		logger.Error("Failed to create admin", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "Failed to create admin", nil)
		return
	}

	logger.Info("Admin created successfully", zap.String("user_id", newUser.ID.String()))
	response.Success(c, http.StatusCreated, "Admin created", gin.H{"admin": newUser})
}
