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
	"github.com/sushantpardhi/shared/logger"
	"github.com/sushantpardhi/shared/response"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
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
	logger.C(c).Info("CreateUser endpoint hit")

	var body map[string]any
	if err := c.ShouldBindJSON(&body); err != nil {
		logger.C(c).Error("Invalid input", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "invalid input", err)
		return
	}
	logger.C(c).Info("JSON input bound successfully")

	// Extract auth safely
	email, _ := body["email"].(string)
	password := generateRandomPassword(10)
	body["temporary_password"] = password
	role, _ := body["role"].(string)

	role = strings.TrimSpace(strings.ToLower(role))
	email = strings.ToLower(email)

	if email == "" || password == "" || role == "" {
		logger.C(c).Error("Missing required fields")
		response.Error(c, 400, "missing required fields", nil)
		return
	}

	// Check existing user
	existing, _ := services.GetUserByEmail(c, email)
	if existing.ID != uuid.Nil {
		logger.C(c).Error("User already exists", zap.String("email", email))
		response.Error(c, 400, "user already exists", nil)
		return
	}

	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.C(c).Error("Failed to hash password", zap.Error(err))
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
		logger.C(c).Error("Failed to create user", zap.Error(err))
		response.Error(c, 500, "failed to create user", err)
		return
	}

	logger.C(c).Info("User created successfully", zap.String("user_id", createdAuth.ID.String()), zap.String("email", createdAuth.Email), zap.String("role", createdAuth.Role))
	response.Success(c, http.StatusCreated, "user created", gin.H{
		"id":                 createdAuth.ID,
		"email":              createdAuth.Email,
		"role":               createdAuth.Role,
		"is_active":          createdAuth.IsActive,
		"temporary_password": password,
	})
}

func CreateSuperAdmin(c *gin.Context) {
	logger.C(c).Info("CreateSuperAdmin endpoint hit")

	var body map[string]any
	if err := c.ShouldBindJSON(&body); err != nil {
		logger.C(c).Error("Invalid input", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "invalid input", err)
		return
	}
	logger.C(c).Info("JSON input bound successfully")

	// Extract other fields
	email, _ := body["email"].(string)
	password := generateRandomPassword(10)
	role := "super_admin"

	email = strings.TrimSpace(strings.ToLower(email))

	if email == "" {
		logger.C(c).Error("Email is required")
		response.Error(c, 400, "email is required", nil)
		return
	}

	// Check existing user
	existing, _ := services.GetUserByEmail(c, email)
	if existing.ID != uuid.Nil {
		logger.C(c).Error("User already exists", zap.String("email", email))
		response.Error(c, 400, "user already exists", nil)
		return
	}

	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.C(c).Error("Failed to hash password", zap.Error(err))
		response.Error(c, 500, "failed to hash password", err)
		return
	}

	auth := models.User{
		Email:    email,
		Password: string(hashed),
		Role:     role,
	}

	// Pass temporary password to service for email
	body["temporary_password"] = password

	createdAuth, _, err := services.CreateUserByRole(c, auth, body)
	if err != nil {
		logger.C(c).Error("Failed to create super admin", zap.Error(err))
		response.Error(c, 500, "failed to create super admin", err)
		return
	}

	logger.C(c).Info("Super admin created successfully", zap.String("user_id", createdAuth.ID.String()), zap.String("email", createdAuth.Email))
	response.Success(c, http.StatusCreated, "super admin created", gin.H{
		"id":                 createdAuth.ID,
		"email":              createdAuth.Email,
		"role":               createdAuth.Role,
		"is_active":          createdAuth.IsActive,
		"temporary_password": password,
	})
}
