package services

import (
	"context"
	"department/models"
	"errors"

	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/logger"
	"go.uber.org/zap"
)

func CreateDepartment(c context.Context, input models.Department) (models.Department, error) {
	logger.C(c).Info("Creating new department in database")
	newDepartment := models.Department{
		Name: input.Name,
		Code: input.Code,
	}

	logger.C(c).Info("Saving Department in Department Table", zap.String("department_id", newDepartment.ID.String()))
	if err := db.DB.Create(&newDepartment).Error; err != nil {
		logger.C(c).Error("Failed to create department in Department table", zap.Error(err))
		return models.Department{}, errors.New("failed to create department in Department table")
	}
	logger.C(c).Info("Department created successfully in Department table", zap.String("department_id", newDepartment.ID.String()))

	return newDepartment, nil
}
