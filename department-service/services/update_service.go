package services

import (
	"context"
	"department/models"
	"errors"

	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/logger"
	"go.uber.org/zap"
)

func UpdateDepartment(c context.Context, id string, input models.UpdateDepartment) (models.Department, error) {
	logger.C(c).Info("Updating department in database", zap.String("department_id", id))

	result := db.DB.Model(&models.Department{}).
		Where("id = ?", id).
		Updates(input)

	if result.Error != nil {
		logger.C(c).Error("Failed to update department in database", zap.Error(result.Error), zap.String("department_id", id))
		return models.Department{}, errors.New("failed to update department in database")
	}

	var updatedDepartment models.Department
	if err := db.DB.Where("id = ?", id).First(&updatedDepartment).Error; err != nil {
		logger.C(c).Error("Failed to fetch updated department", zap.Error(err), zap.String("department_id", id))
		return models.Department{}, errors.New("failed to fetch updated department")
	}

	return updatedDepartment, nil
}
