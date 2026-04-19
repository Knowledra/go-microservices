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

func CreateParent(ctx *gin.Context) {
	logger.Ctx(ctx).Info("Creating Parent Profile")

	var input models.Parent
	if err := ctx.ShouldBindJSON(&input); err != nil {
		logger.Ctx(ctx).Error("Invalid Input", zap.Error(err))
		response.Error(ctx, http.StatusBadRequest, "Invalid Input", err)
		return
	}

	parent, err := services.CreateParentProfile(ctx, input)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create parent", err)
		return
	}

	logger.Ctx(ctx).Info("Parent Profile created successfully", zap.String("user_id", parent.ID.String()))
	response.Success(ctx, http.StatusCreated, "Parent created", gin.H{"parent": parent})
}
