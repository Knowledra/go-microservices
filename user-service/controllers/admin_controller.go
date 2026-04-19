package controllers

import (
	"net/http"
	"user/models"
	"user/services"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/logger"
	"github.com/sushantpardhi/shared/response"
	"go.uber.org/zap"
)

func CreateAdmin(ctx *gin.Context) {
	logger.Ctx(ctx).Info("Creating Admin Profile")

	var input models.Admin
	if err := ctx.ShouldBindJSON(&input); err != nil {
		logger.Ctx(ctx).Error("Invalid Input", zap.Error(err))
		response.Error(ctx, http.StatusBadRequest, "Invalid Input", err)
		return
	}

	admin, err := services.CreateAdminProfile(ctx, input)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create admin", err)
		return
	}

	logger.Ctx(ctx).Info("Admin Profile created successfully", zap.String("user_id", admin.ID.String()))
	response.Success(ctx, http.StatusCreated, "Admin created", gin.H{"admin": admin})
}

func UpdateAdmin(c *gin.Context) {
	logger.Ctx(c).Info("UpdateAdmin endpoint hit")
	// Get user_id from ctx
	userId := c.MustGet("user_id").(string)

	// Bind json
	var input models.UpdateAdmin
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Ctx(c).Error("Invalid Input", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "Invalid Input", err)
		return
	}

	// Update admin
	err := services.UpdateAdminUser(c, userId, input)
	if err != nil {
		logger.Ctx(c).Error("Failed to update admin", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "Failed to update admin", err)
		return
	}

	logger.Ctx(c).Info("Admin updated successfully", zap.String("user_id", userId))
	response.Success(c, http.StatusOK, "Admin updated", gin.H{"user": userId})
}

func DeleteAdmin(ctx *gin.Context) {
	logger.Ctx(ctx).Info("Deleting Admin Profile")
	userId := ctx.Param("user_id")
	err := services.DeleteAdminProfile(ctx, userId)
	if err != nil {
		logger.Ctx(ctx).Error("Failed to delete admin", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, "Failed to delete admin", err)
		return
	}
	logger.Ctx(ctx).Info("Admin Profile deleted successfully", zap.String("user_id", userId))
	response.Success(ctx, http.StatusOK, "Admin Profile deleted", nil)
}
