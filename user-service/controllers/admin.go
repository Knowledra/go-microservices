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

func DeleteAdmin(ctx *gin.Context) {
	logger.Ctx(ctx).Info("Deleting Admin Profile")
	// Get user_id from ctx
	userId := ctx.Param("user_id")
	// Delete admin
	err := services.DeleteAdminProfile(ctx, userId)
	if err != nil {
		logger.Ctx(ctx).Error("Failed to delete admin", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, "Failed to delete admin", err)
		return
	}
	logger.Ctx(ctx).Info("Admin Profile deleted successfully", zap.String("user_id", userId))
	response.Success(ctx, http.StatusOK, "Admin Profile deleted", nil)
}
