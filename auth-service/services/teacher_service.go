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

func CreateTeacher(c context.Context, auth models.User, body map[string]any) (createdAuth models.User, responseBody json.RawMessage, err error) {
	name := getString(body, "name")
	lastName := getString(body, "last_name")
	specialization := getString(body, "specialization")
	if specialization == "" {
		return models.User{}, nil, errors.New("specialization is required for teacher")
	}
	if name == "" || lastName == "" {
		return models.User{}, nil, errors.New("name and last_name are required for teacher")
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
		"id":             createdAuth.ID,
		"name":           name,
		"last_name":      lastName,
		"specialization": specialization,
	}

	responseBody, err = callAPI.CallAPI(c, "POST", "http://dev-user-service:8002/api/v1/user/create/teacher", payload)
	if err != nil {
		logger.C(c).Error("Failed to create teacher profile", zap.Error(err))
		rollbackTeacherProfile(c, createdAuth.ID)
		return
	}

	return
}

func rollbackTeacherProfile(c context.Context, id any) {
	if _, err := callAPI.CallAPI(c, "DELETE", fmt.Sprintf("http://dev-user-service:8002/api/v1/user/delete/profile/teacher/%v", id), nil); err != nil {
		logger.C(c).Error("Failed to rollback teacher profile", zap.Error(err), zap.String("teacher_id", fmt.Sprintf("%v", id)))
	}
}
