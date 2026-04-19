package services

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"auth/models"

	"github.com/google/uuid"
	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/logger"

	"go.uber.org/zap"
)

func GetUserByEmail(ctx context.Context, email string) (models.AuthUser, error) {
	var user models.AuthUser
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return models.AuthUser{}, err
	}
	return user, nil
}

func CreateUserByRole(ctx context.Context, auth models.AuthUser, body map[string]any) (models.AuthUser, json.RawMessage, error) {
	switch strings.ToLower(auth.Role) {
	case "admin":
		return CreateAdmin(ctx, auth, body)
	case "teacher":
		return CreateTeacher(ctx, auth, body)
	case "student":
		return CreateStudent(ctx, auth, body)
	case "parent":
		return CreateParent(ctx, auth, body)
	default:
		return models.AuthUser{}, nil, errors.New("invalid role")
	}
}

func saveAuthUser(ctx context.Context, auth *models.AuthUser) (models.AuthUser, error) {
	if err := db.DB.Create(auth).Error; err != nil {
		logger.Ctx(ctx).Error("Failed to save auth user", zap.Error(err))
		return models.AuthUser{}, errors.New("failed to create auth user")
	}
	return *auth, nil
}

func rollbackAuthUser(ctx context.Context, id uuid.UUID) {
	if err := db.DB.Where("id = ?", id).Delete(&models.AuthUser{}).Error; err != nil {
		logger.Ctx(ctx).Error("Failed to rollback auth user", zap.Error(err))
	}
}

func getString(body map[string]any, key string) string {
	value, _ := body[key].(string)
	return strings.TrimSpace(value)
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
