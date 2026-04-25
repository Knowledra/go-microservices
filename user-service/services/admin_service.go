package services

import (
	"context"
	"errors"
	"strings"
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

func CreateSuperAdminProfile(C context.Context, input models.SuperAdmin) (models.SuperAdmin, error) {
	logger.C(C).Info("Creating new super admin user in database", zap.String("user_id", input.ID.String()))
	newSuperAdmin := models.SuperAdmin{
		ID:       input.ID,
		Name:     input.Name,
		LastName: input.LastName,
	}

	logger.C(C).Info("Saving SuperAdmin in SuperAdmin Table", zap.String("user_id", newSuperAdmin.ID.String()))
	if err := db.DB.Create(&newSuperAdmin).Error; err != nil {
		logger.C(C).Error("Failed to create super admin in SuperAdmin table", zap.Error(err))
		return models.SuperAdmin{}, errors.New("failed to create super admin in SuperAdmin table")
	}
	logger.C(C).Info("SuperAdmin created successfully in SuperAdmin table", zap.String("user_id", newSuperAdmin.ID.String()))

	return newSuperAdmin, nil
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

func UpdateSuperAdminUser(c context.Context, superAdminID string, input models.UpdateSuperAdmin) error {
	logger.C(c).Info("Updating super admin user in database", zap.String("user_id", superAdminID))

	if err := db.DB.Model(&models.SuperAdmin{}).
		Where("id = ?", superAdminID).
		Updates(input).Error; err != nil {

		logger.C(c).Error("Failed to update super admin in database", zap.Error(err))
		return errors.New("failed to update super admin in database")
	}
	return nil
}

func DeleteUserProfile(C context.Context, role, userId string) error {
	logger.C(C).Info("Deleting user profile", zap.String("role", role), zap.String("user_id", userId))

	switch strings.ToLower(role) {
	case "admin":
		return db.DB.Where("id = ?", userId).Delete(&models.Admin{}).Error
	case "super_admin":
		return db.DB.Where("id = ?", userId).Delete(&models.SuperAdmin{}).Error
	case "teacher":
		return db.DB.Where("id = ?", userId).Delete(&models.Teacher{}).Error
	case "student":
		return db.DB.Where("id = ?", userId).Delete(&models.Student{}).Error
	case "parent":
		return db.DB.Where("id = ?", userId).Delete(&models.Parent{}).Error
	default:
		return errors.New("invalid role for profile deletion")
	}
}
