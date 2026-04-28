package services

import (
	"auth/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/sushantpardhi/shared/callAPI"
	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/logger"
	"go.uber.org/zap"
)

func CreateAdmin(c context.Context, auth models.User, body map[string]any) (createdAuth models.User, responseBody json.RawMessage, err error) {
	adminSecret := getString(body, "admin_pass")
	name := getString(body, "name")
	lastName := getString(body, "last_name")

	logger.C(c).Info("Validating admin credentials for CreateAdmin")
	if adminSecret == "" || adminSecret != os.Getenv("ADMIN_PASS") {
		logger.C(c).Error("Invalid admin credentials")
		return models.User{}, nil, errors.New("invalid admin credentials")
	}
	if name == "" || lastName == "" {
		return models.User{}, nil, errors.New("name and last_name are required for admin")
	}

	return createUserWithProfileAtomic(c, auth, func(userID uuid.UUID) (json.RawMessage, error) {
		payload := map[string]any{
			"id":        userID,
			"name":      name,
			"last_name": lastName,
		}

		responseBody, callErr := callAPI.CallAPI(c, "POST", "http://dev-user-service:8002/api/v1/user/create/admin", payload)
		if callErr != nil {
			logger.C(c).Error("Failed to create admin profile", zap.Error(callErr))
			return nil, callErr
		}

		return responseBody, nil
	}, rollbackAdminProfile)
}

func CreateSuperAdmin(c context.Context, auth models.User, body map[string]any) (createdAuth models.User, responseBody json.RawMessage, err error) {
	superAdminSecret := getString(body, "super_admin_pass")
	name := getString(body, "name")
	lastName := getString(body, "last_name")

	logger.C(c).Info("Validating super admin credentials for CreateSuperAdmin")
	if superAdminSecret == "" || superAdminSecret != os.Getenv("SUPER_ADMIN_PASS") {
		logger.C(c).Error("Invalid super admin credentials")
		return models.User{}, nil, errors.New("invalid super admin credentials")
	}
	if name == "" || lastName == "" {
		return models.User{}, nil, errors.New("name and last_name are required for super admin")
	}

	// Check if super admin already exists
	var existingSuperAdmin models.User
	if err := db.DB.Where("role = ? AND is_deleted = false", "super_admin").First(&existingSuperAdmin).Error; err == nil {
		logger.C(c).Error("Super admin already exists")
		return models.User{}, nil, errors.New("super admin already exists")
	}

	return createUserWithProfileAtomic(c, auth, func(userID uuid.UUID) (json.RawMessage, error) {
		payload := map[string]any{
			"id":        userID,
			"name":      name,
			"last_name": lastName,
		}

		responseBody, callErr := callAPI.CallAPI(c, "POST", "http://dev-user-service:8002/api/v1/user/create/super-admin", payload)
		if callErr != nil {
			logger.C(c).Error("Failed to create super admin profile", zap.Error(callErr))
			return nil, callErr
		}

		return responseBody, nil
	}, rollbackSuperAdminProfile)
}

func rollbackAdminProfile(c context.Context, id any) {
	if _, err := callAPI.CallAPI(c, "DELETE", fmt.Sprintf("http://dev-user-service:8002/api/v1/user/internal/delete/profile/admin/%v", id), nil); err != nil {
		logger.C(c).Error("Failed to rollback admin profile", zap.Error(err), zap.String("admin_id", fmt.Sprintf("%v", id)))
	}
}

func rollbackSuperAdminProfile(c context.Context, id any) {
	if _, err := callAPI.CallAPI(c, "DELETE", fmt.Sprintf("http://dev-user-service:8002/api/v1/user/internal/delete/profile/super-admin/%v", id), nil); err != nil {
		logger.C(c).Error("Failed to rollback super admin profile", zap.Error(err), zap.String("super_admin_id", fmt.Sprintf("%v", id)))
	}
}
