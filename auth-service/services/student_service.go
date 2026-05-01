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

	createdAuth, responseBody, err = createUserWithProfileAtomic(c, auth, func(userID uuid.UUID) (json.RawMessage, error) {
		payload := map[string]any{
			"id":        userID,
			"name":      name,
			"last_name": lastName,
			"class":     getString(body, "class"),
			"parent_id": parentID,
			"relation":  relation,
		}

		return callAPI.CallAPI(c, "POST", "http://user-service:8002/api/v1/user/create/student", payload)
	}, func(ctx context.Context, id any) {
		studentID, ok := id.(uuid.UUID)
		if !ok {
			return
		}
		rollbackStudentProfile(ctx, studentID)
	})
	if err != nil {
		if createdParent {
			rollbackParentAccount(c, parentID)
		}
		return models.User{}, nil, err
	}

	sendWelcomeEmailAsync(c, "templates/student_welcome.html", "Welcome to the Platform - Student Account Created", map[string]string{
		"FirstName":         name,
		"LastName":          lastName,
		"Email":             auth.Email,
		"TemporaryPassword": getString(body, "temporary_password"),
	})

	return createdAuth, responseBody, nil
}

func createParentForStudent(c context.Context, body map[string]any) (uuid.UUID, error) {
	parentName := getString(body, "parent_name")
	parentLastName := getString(body, "parent_last_name")
	parentEmail := getString(body, "parent_email")

	if parentName == "" || parentLastName == "" || parentEmail == "" {
		return uuid.Nil, errors.New("parent_name, parent_last_name, and parent_email are required when creating a student")
	}
	if err := ValidateEmailAddress(parentEmail); err != nil {
		return uuid.Nil, err
	}

	parentPassword := GenerateTemporaryPassword()

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

	createdParentAuth, _, err := createParentAccount(c, parentAuth, parentName, parentLastName, getString(body, "parent_phone_primary"), parentPassword)
	if err != nil {
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
	_, err := callAPI.CallAPI(C, "DELETE", fmt.Sprintf("http://user-service:8002/api/v1/user/internal/delete/profile/parent/%s", id.String()), nil)
	return err
}

func rollbackStudentProfile(C context.Context, id uuid.UUID) {
	if _, err := callAPI.CallAPI(C, "DELETE", fmt.Sprintf("http://user-service:8002/api/v1/user/internal/delete/profile/student/%s", id.String()), nil); err != nil {
		logger.C(C).Error("Failed to rollback student profile", zap.Error(err), zap.String("student_id", id.String()))
	}
}
