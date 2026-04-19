package controllers

import (
	"auth/models"
	"auth/services"
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/sushantpardhi/shared/response"
)

func generateRandomPassword(length int) string {
	buf := make([]byte, (length+1)/2)
	if _, err := rand.Read(buf); err != nil {
		return "TempPass123!"
	}
	pw := hex.EncodeToString(buf)
	return pw[:length]
}

func CreateUser(c *gin.Context) {
	var body map[string]any
	if err := c.ShouldBindJSON(&body); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid input", err)
		return
	}

	// Extract auth safely
	email, _ := body["email"].(string)
	password := generateRandomPassword(10)
	role, _ := body["role"].(string)

	role = strings.TrimSpace(strings.ToLower(role))
	email = strings.ToLower(email)

	if email == "" || password == "" || role == "" {
		response.Error(c, 400, "missing required fields", nil)
		return
	}

	// Check existing user
	existing, _ := services.GetUserByEmail(c, email)
	if existing.ID != uuid.Nil {
		response.Error(c, 400, "user already exists", nil)
		return
	}

	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Error(c, 500, "failed to hash password", err)
		return
	}

	auth := models.User{
		Email:    email,
		Password: string(hashed),
		Role:     strings.ToLower(role),
	}

	createdAuth, _, err := services.CreateUserByRole(c, auth, body)
	if err != nil {
		response.Error(c, 500, "failed to create user", err)
		return
	}

	response.Success(c, http.StatusCreated, "user created", gin.H{
		"id":                 createdAuth.ID,
		"email":              createdAuth.Email,
		"role":               createdAuth.Role,
		"is_active":          createdAuth.IsActive,
		"temporary_password": password,
	})
}
