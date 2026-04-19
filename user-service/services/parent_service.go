package services

import (
	"context"
	"errors"
	"user/models"

	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/logger"
	"go.uber.org/zap"
)

func CreateParentProfile(ctx context.Context, input models.Parent) (models.Parent, error) {
	logger.Ctx(ctx).Info("Creating new parent user in database", zap.String("user_id", input.ID.String()))
	newParent := models.Parent{
		ID:          input.ID,
		Name:        input.Name,
		LastName:    input.LastName,
		PhoneNumber: input.PhoneNumber,
	}

	logger.Ctx(ctx).Info("Saving Parent in Parent Table", zap.String("user_id", newParent.ID.String()))
	if err := db.DB.Create(&newParent).Error; err != nil {
		logger.Ctx(ctx).Error("Failed to create parent in Parent table", zap.Error(err))
		return models.Parent{}, errors.New("failed to create parent in Parent table")
	}
	logger.Ctx(ctx).Info("Parent created successfully in Parent table", zap.String("user_id", newParent.ID.String()))

	return newParent, nil
}
