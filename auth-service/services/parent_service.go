package services

import (
	"auth/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sushantpardhi/shared/callAPI"
	"github.com/sushantpardhi/shared/logger"
	"go.uber.org/zap"
)

func CreateParent(c context.Context, auth models.User, body map[string]any) (createdAuth models.User, responseBody json.RawMessage, err error) {
	name := getString(body, "name")
	lastName := getString(body, "last_name")
	if name == "" || lastName == "" {
		return models.User{}, nil, errors.New("name and last_name are required for parent")
	}

	createdAuth, err = saveAuthUser(c, &auth)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			rollbackAuthUser(c, createdAuth.ID)
		}
	}()

	payload := map[string]any{
		"id":           createdAuth.ID,
		"name":         name,
		"last_name":    lastName,
		"phone_number": getString(body, "phone_number"),
	}

	responseBody, err = callAPI.CallAPI(c, "POST", "http://dev-user-service:8002/api/v1/user/create/parent", payload)
	if err != nil {
		logger.C(c).Error("Failed to create parent profile", zap.Error(err))
		return
	}

	return
}

func rollbackParentProfileAccount(c context.Context, id any) {
	if _, err := callAPI.CallAPI(c, "DELETE", fmt.Sprintf("http://dev-user-service:8002/api/v1/user/internal/delete/profile/parent/%v", id), nil); err != nil {
		logger.C(c).Error("Failed to rollback parent profile", zap.Error(err), zap.String("parent_id", fmt.Sprintf("%v", id)))
	}
}
