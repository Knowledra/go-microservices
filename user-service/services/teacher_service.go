package services

import (
	"context"
	"errors"
	"user/models"

	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/logger"
	"go.uber.org/zap"
)

func CreateTeacherProfile(ctx context.Context, input models.Teacher) (models.Teacher, error) {
	logger.Ctx(ctx).Info("Creating new teacher user in database", zap.String("user_id", input.ID.String()))
	newTeacher := models.Teacher{
		ID:             input.ID,
		Name:           input.Name,
		LastName:       input.LastName,
		Specialization: input.Specialization,
	}

	logger.Ctx(ctx).Info("Saving Teacher in Teacher Table", zap.String("user_id", newTeacher.ID.String()))
	if err := db.DB.Create(&newTeacher).Error; err != nil {
		logger.Ctx(ctx).Error("Failed to create teacher in Teacher table", zap.Error(err))
		return models.Teacher{}, errors.New("failed to create teacher in Teacher table")
	}
	logger.Ctx(ctx).Info("Teacher created successfully in Teacher table", zap.String("user_id", newTeacher.ID.String()))

	return newTeacher, nil
}

func DeleteTeacherProfile(ctx context.Context, userId string) error {
	if err := db.DB.Where("id = ?", userId).Delete(&models.Teacher{}).Error; err != nil {
		logger.Ctx(ctx).Error("Failed to delete teacher from User table", zap.Error(err))
		return errors.New("failed to delete teacher from User table")
	}
	return nil
}
