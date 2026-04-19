package services

import (
	"auth/models"
	"context"
	"encoding/json"
	"errors"
	"os"

	"github.com/sushantpardhi/shared/callAPI"
	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/logger"
	"go.uber.org/zap"
)

func CreateAdmin(ctx context.Context, auth models.AuthUser, body map[string]any) (models.AuthUser, json.RawMessage, error) {
	adminSecret := os.Getenv("ADMIN_PASS")
	adminPassword := getString(body, "admin_password")
	name := getString(body, "name")
	lastName := getString(body, "last_name")

	logger.Ctx(ctx).Info("Validating admin credentials for CreateAdmin")
	if adminPassword == "" || adminPassword != adminSecret {
		logger.Ctx(ctx).Error("Invalid admin credentials")
		return models.AuthUser{}, nil, errors.New("invalid admin credentials")
	}
	if name == "" || lastName == "" {
		return models.AuthUser{}, nil, errors.New("name and last_name are required for admin")
	}

	createdAuth, err := saveAuthUser(ctx, &auth)
	if err != nil {
		return models.AuthUser{}, nil, err
	}

	payload := map[string]any{
		"id":        createdAuth.ID,
		"name":      name,
		"last_name": lastName,
	}

	responseBody, err := callAPI.CallAPI(ctx, "POST", "http://user-service:8002/api/v1/profile/create/admin", payload)
	if err != nil {
		rollbackAuthUser(ctx, createdAuth.ID)
		logger.Ctx(ctx).Error("Failed to create admin profile", zap.Error(err))
		return models.AuthUser{}, nil, err
	}

	return createdAuth, responseBody, nil
}

func DeleteAdminUser(ctx context.Context, userId string) error {
	if err := db.DB.Where("id = ?", userId).Delete(&models.AuthUser{}).Error; err != nil {
		logger.Ctx(ctx).Error("Failed to delete user from Auth table", zap.Error(err))
		return errors.New("failed to delete user from Auth table")
	}
	logger.Ctx(ctx).Info("User deleted successfully from Auth table", zap.String("user_id", userId))

	logger.Ctx(ctx).Info("Calling User-Service API to delete admin in User table", zap.String("user_id", userId))
	_, err := callAPI.CallAPI(ctx, "DELETE", "http://user-service:8002/api/v1/profile/delete/admin/"+userId, nil)
	if err != nil {
		logger.Ctx(ctx).Error("Failed to delete admin in User table", zap.Error(err))
		return errors.New("failed to delete admin in User table")
	}

	return nil
}
