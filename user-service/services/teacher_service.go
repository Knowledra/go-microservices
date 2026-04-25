package services

import (
	"context"
	"errors"
	"user/models"

	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/logger"
	"go.uber.org/zap"
)

func CreateTeacherProfile(C context.Context, input models.Teacher) (models.Teacher, error) {
	logger.C(C).Info("Creating new teacher user in database", zap.String("user_id", input.ID.String()))
	newTeacher := models.Teacher{
		ID:             input.ID,
		Name:           input.Name,
		LastName:       input.LastName,
		Specialization: input.Specialization,
	}

	logger.C(C).Info("Saving Teacher in Teacher Table", zap.String("user_id", newTeacher.ID.String()))
	if err := db.DB.Create(&newTeacher).Error; err != nil {
		logger.C(C).Error("Failed to create teacher in Teacher table", zap.Error(err))
		return models.Teacher{}, errors.New("failed to create teacher in Teacher table")
	}
	logger.C(C).Info("Teacher created successfully in Teacher table", zap.String("user_id", newTeacher.ID.String()))

	return newTeacher, nil
}

func UpdateTeacherUser(C context.Context, teacherID string, input models.UpdateTeacher) error {
	logger.C(C).Info("Updating teacher user in database", zap.String("user_id", teacherID))

	if err := db.DB.Model(&models.Teacher{}).
		Where("id = ?", teacherID).
		Updates(input).Error; err != nil {

		logger.C(C).Error("Failed to update teacher in database", zap.Error(err))
		return errors.New("failed to update teacher in database")
	}
	return nil
}
