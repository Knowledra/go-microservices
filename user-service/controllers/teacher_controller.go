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

func UpdateTeacher(c *gin.Context) {
	logger.Ctx(c).Info("UpdateTeacher endpoint hit")
	// Get user_id from ctx
	userId := c.MustGet("user_id").(string)

	// Bind json
	var input models.UpdateTeacher
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Ctx(c).Error("Invalid Input", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "Invalid Input", err)
		return
	}

	// Update teacher
	err := services.UpdateTeacherUser(c, userId, input)
	if err != nil {
		logger.Ctx(c).Error("Failed to update teacher", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "Failed to update teacher", err)
		return
	}

	logger.Ctx(c).Info("Teacher updated successfully", zap.String("user_id", userId))
	response.Success(c, http.StatusOK, "Teacher updated", gin.H{"user": userId})
}
