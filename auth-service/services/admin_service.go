package services

import (
	"auth/models"
	"context"
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/sushantpardhi/shared/callAPI"
	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/logger"
	"go.uber.org/zap"
)

func CreateAdmin(ctx context.Context, auth models.User, body map[string]any) (createdAuth models.User, responseBody json.RawMessage, err error) {
	adminSecret := os.Getenv("ADMIN_PASS")
	name := getString(body, "name")
	lastName := getString(body, "last_name")

	logger.Ctx(ctx).Info("Validating admin credentials for CreateAdmin")
	if adminSecret == "" || adminSecret != os.Getenv("ADMIN_PASS") {
		logger.Ctx(ctx).Error("Invalid admin credentials")
		return models.User{}, nil, errors.New("invalid admin credentials")
	}
	if name == "" || lastName == "" {
		return models.User{}, nil, errors.New("name and last_name are required for admin")
	}

	createdAuth, err = saveAuthUser(ctx, &auth)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			rollbackAuthUser(ctx, createdAuth.ID)
		}
	}()

	payload := map[string]any{
		"id":        createdAuth.ID,
		"name":      name,
		"last_name": lastName,
	}

	responseBody, err = callAPI.CallAPI(ctx, "POST", "http://user-service:8002/api/v1/profile/create/admin", payload)
	if err != nil {
		logger.Ctx(ctx).Error("Failed to create admin profile", zap.Error(err))
		return
	}

	return
}

func DeleteUserAccount(ctx context.Context, userId string) error {
	// check if user exists and is not already deleted
	var user models.User
	if err := db.DB.Where("id = ? AND is_deleted = false", userId).First(&user).Error; err != nil {
		logger.Ctx(ctx).Error("User not found or already deleted", zap.String("user_id", userId), zap.Error(err))
		return errors.New("user not found or already deleted")
	}

	// Mark user as deleted in Auth table
	if err := db.DB.Model(&models.User{}).Where("id = ?", userId).Update("is_deleted", true).Update("deleted_at", time.Now()).Error; err != nil {
		logger.Ctx(ctx).Error("Failed to mark user as deleted in Auth table", zap.Error(err))
		return errors.New("failed to mark user as deleted")
	}

	logger.Ctx(ctx).Info("User marked as deleted in Auth table", zap.String("user_id", userId))
	return nil
}
