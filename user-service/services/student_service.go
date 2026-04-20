package services

import (
	"context"
	"errors"
	"user/models"

	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/logger"
	"go.uber.org/zap"
)

func CreateStudentProfile(C context.Context, input models.Student) (models.Student, error) {
	logger.C(C).Info("Creating new student user in database", zap.String("user_id", input.ID.String()))
	newStudent := models.Student{
		ID:       input.ID,
		Name:     input.Name,
		LastName: input.LastName,
		Class:    input.Class,
		ParentID: input.ParentID,
		Relation: input.Relation,
	}

	logger.C(C).Info("Saving Student in Student Table", zap.String("user_id", newStudent.ID.String()))
	if err := db.DB.Create(&newStudent).Error; err != nil {
		logger.C(C).Error("Failed to create student in Student table", zap.Error(err))
		return models.Student{}, errors.New("failed to create student in Student table")
	}
	logger.C(C).Info("Student created successfully in Student table", zap.String("user_id", newStudent.ID.String()))

	return newStudent, nil
}

func UpdateStudentUser(C context.Context, studentID string, input models.UpdateStudent) error {
	logger.C(C).Info("Updating student profile in database", zap.String("student_id", studentID))

	if err := db.DB.Model(&models.Student{}).
		Where("id = ?", studentID).
		Updates(input).Error; err != nil {

		logger.C(C).Error("Failed to update student profile in database", zap.Error(err))
		return errors.New("failed to update student profile in database")
	}
	return nil
}
