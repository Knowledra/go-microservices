package services

import (
	"auth/models"
	"context"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/sushantpardhi/shared/callAPI"
	"golang.org/x/crypto/bcrypt"
)

func CreateStudent(ctx context.Context, auth models.User, body map[string]any) (createdAuth models.User, responseBody json.RawMessage, err error) {
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
		parentID, err = createParentForStudent(ctx, body)
		if err != nil {
			return models.User{}, nil, err
		}
		createdParent = true
	}

	createdAuth, err = saveAuthUser(ctx, &auth)
	if err != nil {
		if createdParent {
			rollbackParentAccount(ctx, parentID)
		}
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
		"class":     getString(body, "class"),
		"parent_id": parentID,
		"relation":  relation,
	}

	responseBody, err = callAPI.CallAPI(ctx, "POST", "http://user-service:8002/api/v1/profile/create/student", payload)
	if err != nil {
		if createdParent {
			rollbackParentAccount(ctx, parentID)
		}
		return
	}

	return
}

func createParentForStudent(ctx context.Context, body map[string]any) (uuid.UUID, error) {
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
