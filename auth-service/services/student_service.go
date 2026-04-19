package services

import (
	"auth/models"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/sushantpardhi/shared/callAPI"
	"golang.org/x/crypto/bcrypt"
)

func CreateStudent(ctx context.Context, auth models.AuthUser, body map[string]any) (models.AuthUser, json.RawMessage, error) {
	name := getString(body, "name")
	lastName := getString(body, "last_name")
	if name == "" || lastName == "" {
		return models.AuthUser{}, nil, errors.New("name and last_name are required for student")
	}

	parentID, err := getUUID(body, "parent_id")
	createdParent := false
	if err != nil {
		return models.AuthUser{}, nil, err
	}
	if parentID == uuid.Nil {
		parentID, err = createParentForStudent(ctx, body)
		if err != nil {
			return models.AuthUser{}, nil, err
		}
		createdParent = true
	}

	createdAuth, err := saveAuthUser(ctx, &auth)
	if err != nil {
		if createdParent {
			rollbackParentAccount(ctx, parentID)
		}
		return models.AuthUser{}, nil, err
	}

	payload := map[string]any{
		"id":        createdAuth.ID,
		"name":      name,
		"last_name": lastName,
		"class":     getString(body, "class"),
		"parent_id": parentID,
	}

	responseBody, err := callAPI.CallAPI(ctx, "POST", "http://user-service:8002/api/v1/profile/create/student", payload)
	if err != nil {
		rollbackAuthUser(ctx, createdAuth.ID)
		if createdParent {
			rollbackParentAccount(ctx, parentID)
		}
		return models.AuthUser{}, nil, err
	}

	return createdAuth, responseBody, nil
}

func createParentForStudent(ctx context.Context, body map[string]any) (uuid.UUID, error) {
	parentName := getString(body, "parent_name")
	parentLastName := getString(body, "parent_last_name")
	parentEmail := getString(body, "parent_email")
	if parentName == "" || parentLastName == "" || parentEmail == "" {
		return uuid.Nil, errors.New("parent_name, parent_last_name, and parent_email are required when creating a student")
	}

	parentPassword := getString(body, "parent_password")
	if parentPassword == "" {
		parentPassword = generateRandomPassword(12)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(parentPassword), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, err
	}

	parentAuth := models.AuthUser{
		Email:    parentEmail,
		Password: string(hashedPassword),
		Role:     "parent",
		IsActive: true,
	}

	createdParentAuth, err := saveAuthUser(ctx, &parentAuth)
	if err != nil {
		return uuid.Nil, err
	}

	payload := map[string]any{
		"id":           createdParentAuth.ID,
		"name":         parentName,
		"last_name":    parentLastName,
		"phone_number": getString(body, "parent_phone_primary"),
	}

	_, err = callAPI.CallAPI(ctx, "POST", "http://user-service:8002/api/v1/profile/create/parent", payload)
	if err != nil {
		rollbackAuthUser(ctx, createdParentAuth.ID)
		return uuid.Nil, err
	}

	return createdParentAuth.ID, nil
}

func rollbackParentAccount(ctx context.Context, id uuid.UUID) {
	rollbackAuthUser(ctx, id)
}

func generateRandomPassword(length int) string {
	buf := make([]byte, (length+1)/2)
	if _, err := rand.Read(buf); err != nil {
		return "TempPass123!"
	}
	pw := hex.EncodeToString(buf)
	return pw[:length]
}
