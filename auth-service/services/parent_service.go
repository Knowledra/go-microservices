package services

import (
	"auth/models"
	"context"
	"encoding/json"
	"errors"

	"github.com/sushantpardhi/shared/callAPI"
)

func CreateParent(ctx context.Context, auth models.User, body map[string]any) (createdAuth models.User, responseBody json.RawMessage, err error) {
	name := getString(body, "name")
	lastName := getString(body, "last_name")
	if name == "" || lastName == "" {
		return models.User{}, nil, errors.New("name and last_name are required for parent")
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
		"id":           createdAuth.ID,
		"name":         name,
		"last_name":    lastName,
		"phone_number": getString(body, "phone_number"),
	}

	responseBody, err = callAPI.CallAPI(ctx, "POST", "http://user-service:8002/api/v1/profile/create/parent", payload)
	return
}
