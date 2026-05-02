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

func CreateParent(c context.Context, auth models.User, body map[string]any) (createdAuth models.User, responseBody json.RawMessage, err error) {
	name := getString(body, "name")
	lastName := getString(body, "last_name")
	if name == "" || lastName == "" {
		return models.User{}, nil, errors.New("name and last_name are required for parent")
	}

	return createParentAccount(c, auth, name, lastName, getString(body, "phone_number"), getString(body, "temporary_password"))
}

func createParentAccount(c context.Context, auth models.User, name string, lastName string, phoneNumber string, temporaryPassword string) (createdAuth models.User, responseBody json.RawMessage, err error) {
	createdAuth, responseBody, err = createUserWithProfileAtomic(c, auth, func(userID uuid.UUID) (json.RawMessage, error) {
		payload := map[string]any{
			"id":           userID,
			"name":         name,
			"last_name":    lastName,
			"phone_number": phoneNumber,
		}

		responseBody, callErr := callAPI.CallAPI(c, "POST", "http://user-service:8002/api/v1/create/parent", payload)
		if callErr != nil {
			logger.C(c).Error("Failed to create parent profile", zap.Error(callErr))
			return nil, callErr
		}

		return responseBody, nil
	}, rollbackParentProfileAccount)
	if err != nil {
		return createdAuth, responseBody, err
	}

	sendWelcomeEmailAsync(c, "templates/parent_welcome.html", "Welcome to the Platform - Parent Account Created", map[string]string{
		"FirstName":         name,
		"LastName":          lastName,
		"Email":             auth.Email,
		"TemporaryPassword": temporaryPassword,
	})

	return createdAuth, responseBody, nil
}

func rollbackParentProfileAccount(c context.Context, id any) {
	if _, err := callAPI.CallAPI(c, "DELETE", fmt.Sprintf("http://user-service:8002/api/v1/internal/delete/profile/parent/%v", id), nil); err != nil {
		logger.C(c).Error("Failed to rollback parent profile", zap.Error(err), zap.String("parent_id", fmt.Sprintf("%v", id)))
	}
}
