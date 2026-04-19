package services

import (
	"auth/models"
	"context"
	"encoding/json"
	"errors"
	"os"

	"github.com/sushantpardhi/shared/callAPI"
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
