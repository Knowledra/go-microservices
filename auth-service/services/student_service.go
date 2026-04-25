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
	"golang.org/x/crypto/bcrypt"
)

func CreateStudent(c context.Context, auth models.User, body map[string]any) (createdAuth models.User, responseBody json.RawMessage, err error) {
	// Getting required fields for student profile
	name := getString(body, "name")
	lastName := getString(body, "last_name")
	relation := getString(body, "relation")
	if name == "" || lastName == "" || relation == "" {
		return models.User{}, nil, errors.New("name, last_name, and relation are required for student")
	}

	parentID, err := getUUID(body, "parent_id")
	createdParent := false
	if err != nil {
		return models.User{}, nil, err
	}
	if parentID == uuid.Nil {
		parentID, err = createParentForStudent(c, body)
		if err != nil {
			return models.User{}, nil, err
		}
		createdParent = true
	}

	createdAuth, err = saveAuthUser(c, &auth)
	if err != nil {
		if createdParent {
			rollbackParentAccount(c, parentID)
		}
		return
	}
	defer func() {
		if err != nil {
			rollbackAuthUser(c, createdAuth.ID)
		}
	}()

	payload := map[string]any{
		"id":        createdAuth.ID,
		"name":      name,
		"last_name": lastName,
		"class":     getString(body, "class"),
		"parent_id": parentID,
		"relation":  relation,
	}

	responseBody, err = callAPI.CallAPI(c, "POST", "http://dev-user-service:8002/api/v1/user/create/student", payload)
	if err != nil {
		if createdParent {
			rollbackParentAccount(c, parentID)
		}
		return
	}

	return
}

func createParentForStudent(c context.Context, body map[string]any) (uuid.UUID, error) {
	parentName := getString(body, "parent_name")
	parentLastName := getString(body, "parent_last_name")
	parentEmail := getString(body, "parent_email")

	if parentName == "" || parentLastName == "" || parentEmail == "" {
		return uuid.Nil, errors.New("parent_name, parent_last_name, and parent_email are required when creating a student")
	}

	parentPassword := uuid.New().String()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(parentPassword), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, err
	}

	parentAuth := models.User{
		Email:    parentEmail,
		Password: string(hashedPassword),
		Role:     "parent",
		IsActive: true,
	}

	createdParentAuth, err := saveAuthUser(c, &parentAuth)
	if err != nil {
		return uuid.Nil, err
	}

	payload := map[string]any{
		"id":           createdParentAuth.ID,
		"name":         parentName,
		"last_name":    parentLastName,
		"phone_number": getString(body, "parent_phone_primary"),
	}

	_, err = callAPI.CallAPI(c, "POST", "http://dev-user-service:8002/api/v1/user/create/parent", payload) //nolint:bodyclose
	if err != nil {
		rollbackAuthUser(c, createdParentAuth.ID)
		return uuid.Nil, err
	}

	return createdParentAuth.ID, nil
}

func rollbackParentAccount(C context.Context, id uuid.UUID) {
	if err := rollbackParentProfile(C, id); err != nil {
		logger.C(C).Error("Failed to rollback parent profile", zap.Error(err), zap.String("parent_id", id.String()))
	}
	rollbackAuthUser(C, id)
}

func rollbackParentProfile(C context.Context, id uuid.UUID) error {
	_, err := callAPI.CallAPI(C, "DELETE", fmt.Sprintf("http://dev-user-service:8002/api/v1/user/internal/delete/profile/parent/%s", id.String()), nil)
	return err
}

func rollbackStudentProfile(C context.Context, id uuid.UUID) {
	if _, err := callAPI.CallAPI(C, "DELETE", fmt.Sprintf("http://dev-user-service:8002/api/v1/user/internal/delete/profile/student/%s", id.String()), nil); err != nil {
		logger.C(C).Error("Failed to rollback student profile", zap.Error(err), zap.String("student_id", id.String()))
	}
}
