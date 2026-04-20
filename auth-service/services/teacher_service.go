package services

import (
	"auth/models"
	"context"
	"encoding/json"
	"errors"

	"github.com/sushantpardhi/shared/callAPI"
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

	responseBody, err = callAPI.CallAPI(c, "POST", "http://user-service:8002/api/v1/profile/create/teacher", payload)
	return
}
