package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"auth/models"

	"github.com/google/uuid"
	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/logger"

	"go.uber.org/zap"
)

func GetUserByEmail(c context.Context, email string) (models.User, error) {
	var user models.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func GetUserByID(c context.Context, id string) (models.User, error) {
	var user models.User
	if err := db.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func CreateUserByRole(c context.Context, auth models.User, body map[string]any) (models.User, json.RawMessage, error) {
	switch strings.ToLower(auth.Role) {
	case "super_admin":
		return CreateSuperAdmin(c, auth, body)
	case "admin":
		return CreateAdmin(c, auth, body)
	case "teacher":
		return CreateTeacher(c, auth, body)
	case "student":
		return CreateStudent(c, auth, body)
	case "parent":
		return CreateParent(c, auth, body)
	default:
		return models.User{}, nil, errors.New("invalid role")
	}
}

func saveAuthUser(c context.Context, auth *models.User) (models.User, error) {
	if err := db.DB.Create(auth).Error; err != nil {
		logger.C(c).Error("Failed to save auth user", zap.Error(err))
		return models.User{}, errors.New("failed to create auth user")
	}
	return *auth, nil
}

func rollbackAuthUser(c context.Context, id uuid.UUID) {
	if id == uuid.Nil {
		return
	}
	if err := db.DB.Unscoped().Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		logger.C(c).Error("Failed to rollback auth user", zap.Error(err))
	}
}

func getString(body map[string]any, key string) string {
	if val, ok := body[key]; ok && val != nil {
		switch v := val.(type) {
		case string:
			return strings.TrimSpace(v)
		default:
			return strings.TrimSpace(fmt.Sprintf("%v", v))
		}
	}
	return ""
}

func getUUID(body map[string]any, key string) (uuid.UUID, error) {
	if raw, ok := body[key]; ok && raw != nil {
		if str, ok := raw.(string); ok {
			id, err := uuid.Parse(strings.TrimSpace(str))
			if err != nil {
				return uuid.Nil, errors.New("invalid uuid for " + key)
			}
			return id, nil
		}
		return uuid.Nil, errors.New("invalid type for " + key)
	}
	return uuid.Nil, nil
}

func DeleteUserAccount(c context.Context, userId string) error {
	// check if user exists and is not already deleted
	var user models.User
	if err := db.DB.Where("id = ? AND is_deleted = false", userId).First(&user).Error; err != nil {
		logger.C(c).Error("User not found or already deleted", zap.String("user_id", userId), zap.Error(err))
		return errors.New("user not found or already deleted")
	}

	// Mark user as deleted in Auth table
	if err := db.DB.Model(&models.User{}).Where("id = ?", userId).Update("is_deleted", true).Update("deleted_at", time.Now()).Error; err != nil {
		logger.C(c).Error("Failed to mark user as deleted in Auth table", zap.Error(err))
		return errors.New("failed to mark user as deleted")
	}

	logger.C(c).Info("User marked as deleted in Auth table", zap.String("user_id", userId))
	return nil
}
