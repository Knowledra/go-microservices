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

func CreateTeacher(ctx *gin.Context) {
	logger.Ctx(ctx).Info("Creating Teacher Profile")

	var input models.Teacher
	if err := ctx.ShouldBindJSON(&input); err != nil {
		logger.Ctx(ctx).Error("Invalid Input", zap.Error(err))
		response.Error(ctx, http.StatusBadRequest, "Invalid Input", err)
		return
	}

	teacher, err := services.CreateTeacherProfile(ctx, input)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create teacher", err)
		return
	}

	logger.Ctx(ctx).Info("Teacher Profile created successfully", zap.String("user_id", teacher.ID.String()))
	response.Success(ctx, http.StatusCreated, "Teacher created", gin.H{"teacher": teacher})
}
