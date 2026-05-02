package services

import (
	"auth/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
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

	createdAuth, responseBody, err = createUserWithProfileAtomic(c, auth, func(userID uuid.UUID) (json.RawMessage, error) {
		payload := map[string]any{
			"id":             userID,
			"name":           name,
			"last_name":      lastName,
			"specialization": specialization,
		}

		responseBody, callErr := callAPI.CallAPI(c, "POST", "http://user-service:8002/api/v1/create/teacher", payload)
		if callErr != nil {
			logger.C(c).Error("Failed to create teacher profile", zap.Error(callErr))
			return nil, callErr
		}

		return responseBody, nil
	}, rollbackTeacherProfile)
	if err != nil {
		return createdAuth, responseBody, err
	}

	sendWelcomeEmailAsync(c, "templates/teacher_welcome.html", "Welcome to the Platform - Teacher Account Created", map[string]string{
		"FirstName":         name,
		"LastName":          lastName,
		"Email":             auth.Email,
		"TemporaryPassword": getString(body, "temporary_password"),
	})

	return createdAuth, responseBody, nil
}

func rollbackTeacherProfile(c context.Context, id any) {
	if _, err := callAPI.CallAPI(c, "DELETE", fmt.Sprintf("http://user-service:8002/api/v1/internal/delete/profile/teacher/%v", id), nil); err != nil {
		logger.C(c).Error("Failed to rollback teacher profile", zap.Error(err), zap.String("teacher_id", fmt.Sprintf("%v", id)))
	}
}
