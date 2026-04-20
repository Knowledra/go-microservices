package services

import (
	"context"
	"errors"
	"user/models"

	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/logger"
	"go.uber.org/zap"
)

func CreateAdminProfile(C context.Context, input models.Admin) (models.Admin, error) {
	logger.C(C).Info("Creating new admin user in database", zap.String("user_id", input.ID.String()))
	newAdmin := models.Admin{
		ID:       input.ID,
		Name:     input.Name,
		LastName: input.LastName,
	}

	logger.C(C).Info("Saving Admin in Admin Table", zap.String("user_id", newAdmin.ID.String()))
	if err := db.DB.Create(&newAdmin).Error; err != nil {
		logger.C(C).Error("Failed to create admin in Admin table", zap.Error(err))
		return models.Admin{}, errors.New("failed to create admin in Admin table")
	}
	logger.C(C).Info("Admin created successfully in Admin table", zap.String("user_id", newAdmin.ID.String()))

	return newAdmin, nil
}

func UpdateAdminUser(c context.Context, adminID string, input models.UpdateAdmin) error {
	logger.C(c).Info("Updating admin user in database", zap.String("user_id", adminID))

	if err := db.DB.Model(&models.Admin{}).
		Where("id = ?", adminID).
		Updates(input).Error; err != nil {

		logger.C(c).Error("Failed to update admin in database", zap.Error(err))
		return errors.New("failed to update admin in database")
	}
	return nil
}
