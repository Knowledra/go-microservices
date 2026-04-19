package services

import (
	"context"
	"errors"
	"user/models"

	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/logger"
	"go.uber.org/zap"
)

func CreateStudentProfile(ctx context.Context, input models.Student) (models.Student, error) {
	logger.Ctx(ctx).Info("Creating new student user in database", zap.String("user_id", input.ID.String()))
	newStudent := models.Student{
		ID:       input.ID,
		Name:     input.Name,
		LastName: input.LastName,
		Class:    input.Class,
		ParentID: input.ParentID,
	}

	logger.Ctx(ctx).Info("Saving Student in Student Table", zap.String("user_id", newStudent.ID.String()))
	if err := db.DB.Create(&newStudent).Error; err != nil {
		logger.Ctx(ctx).Error("Failed to create student in Student table", zap.Error(err))
		return models.Student{}, errors.New("failed to create student in Student table")
	}
	logger.Ctx(ctx).Info("Student created successfully in Student table", zap.String("user_id", newStudent.ID.String()))

	return newStudent, nil
}
