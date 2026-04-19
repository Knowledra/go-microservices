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

func CreateStudent(ctx *gin.Context) {
	logger.Ctx(ctx).Info("Creating Student Profile")

	var input models.Student
	if err := ctx.ShouldBindJSON(&input); err != nil {
		logger.Ctx(ctx).Error("Invalid Input", zap.Error(err))
		response.Error(ctx, http.StatusBadRequest, "Invalid Input", err)
		return
	}

	student, err := services.CreateStudentProfile(ctx, input)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create student", err)
		return
	}

	logger.Ctx(ctx).Info("Student Profile created successfully", zap.String("user_id", student.ID.String()))
	response.Success(ctx, http.StatusCreated, "Student created", gin.H{"student": student})
}
