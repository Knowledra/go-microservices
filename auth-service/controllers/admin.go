package controllers

import (
	"auth/models"
	"auth/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/logger"
	"golang.org/x/crypto/bcrypt"

	"github.com/sushantpardhi/shared/response"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	var body map[string]any
	if err := c.ShouldBindJSON(&body); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid input", err)
		return
	}

	// Extract auth safely
	email, _ := body["email"].(string)
	password, _ := body["password"].(string)
	role, _ := body["role"].(string)

	email = strings.ToLower(email)

	if email == "" || password == "" || role == "" {
		response.Error(c, 400, "missing required fields", nil)
		return
	}

	// Check existing user
	existing, _ := services.GetUserByEmail(c, email)
	if existing.ID.String() != "" {
		response.Error(c, 400, "user already exists", nil)
		return
	}

	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Error(c, 500, "failed to hash password", err)
		return
	}

	auth := models.AuthUser{
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
		"id":        createdAuth.ID,
		"email":     createdAuth.Email,
		"role":      createdAuth.Role,
		"is_active": createdAuth.IsActive,
	})
}

func DeleteUser(c *gin.Context) {
	logger.Ctx(c).Info("DeleteAdmin endpoint hit")
	// Get user_id from ctx
	userId := c.MustGet("user_id").(string)

	// Delete admin
	err := services.DeleteAdminUser(c, userId)
	if err != nil {
		logger.Ctx(c).Error("Failed to delete admin", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "Failed to delete admin", err)
		return
	}

	// Remove cookies for deleted user
	logger.Info("Removing cookies")
	c.SetCookie("access_token", "", -1, "/", "", true, true)
	logger.Info("Cookies removed successfully")

	logger.Ctx(c).Info("Admin deleted successfully", zap.String("user_id", userId))
	response.Success(c, http.StatusOK, "Admin deleted", gin.H{"user": userId})
}
