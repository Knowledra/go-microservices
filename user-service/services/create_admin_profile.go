package services

import (
	"errors"
	"user/models"

	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/logger"
	"go.uber.org/zap"
)

func CreateAdminProfile(input models.Admin) (models.Admin, error) {
	// Create user
	logger.Info("Creating new admin user in database", zap.String("user_id", input.ID.String()))
	newAdmin := models.Admin{
		Name:     input.Name,
		LastName: input.LastName,
		ID:       input.ID,
	}

	// Save user in User table
	logger.Info("Saving Admin in Admin Table", zap.String("user_id", newAdmin.ID.String()))
	if err := db.DB.Create(&newAdmin).Error; err != nil {
		logger.Error("Failed to create admin in Admin table", zap.Error(err))
		return models.Admin{}, errors.New("failed to create admin in Admin table")
	}
	logger.Info("Admin created successfully in Admin table", zap.String("user_id", newAdmin.ID.String()))

	return newAdmin, nil
}
