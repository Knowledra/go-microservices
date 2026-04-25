package services

import (
	"context"
	"errors"
	"user/models"

	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/logger"
	"go.uber.org/zap"
)

func CreateParentProfile(C context.Context, input models.Parent) (models.Parent, error) {
	logger.C(C).Info("Creating new parent user in database", zap.String("user_id", input.ID.String()))
	newParent := models.Parent{
		ID:          input.ID,
		Name:        input.Name,
		LastName:    input.LastName,
		PhoneNumber: input.PhoneNumber,
	}

	logger.C(C).Info("Saving Parent in Parent Table", zap.String("user_id", newParent.ID.String()))
	if err := db.DB.Create(&newParent).Error; err != nil {
		logger.C(C).Error("Failed to create parent in Parent table", zap.Error(err))
		return models.Parent{}, errors.New("failed to create parent in Parent table")
	}
	logger.C(C).Info("Parent created successfully in Parent table", zap.String("user_id", newParent.ID.String()))

	return newParent, nil
}

func UpdateParentUser(C context.Context, parentId string, input models.UpdateParent) error {
	logger.C(C).Info("Updating parent user in database", zap.String("user_id", parentId))

	if err := db.DB.Model(&models.Parent{}).
		Where("id = ?", parentId).
		Updates(input).Error; err != nil {

		logger.C(C).Error("Failed to update parent in database", zap.Error(err))
		return errors.New("failed to update parent in database")
	}
	return nil
}
