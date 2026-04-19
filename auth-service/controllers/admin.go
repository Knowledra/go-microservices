package controllers

import (
	"auth/models"
	"auth/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/logger"

	"github.com/sushantpardhi/shared/response"
	"go.uber.org/zap"
)

func CreateAdmin(c *gin.Context) {
	logger.Ctx(c).Info("CreateAdmin endpoint hit")
	// Bind json
	var input models.RegisterAdmin

	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Ctx(c).Error("Invalid Input", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "Invalid Input", err)
		return
	}

	user, admin, err := services.CreateAdminUser(c, input)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create admin", err)
		return
	}

	logger.Ctx(c).Info("Admin created successfully", zap.String("user_id", user.ID.String()))
	response.Success(c, http.StatusCreated, "Admin created", gin.H{"user": user, "admin": admin})
}
