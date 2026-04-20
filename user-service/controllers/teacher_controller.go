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

func CreateTeacher(C *gin.Context) {
	logger.C(C).Info("Creating Teacher Profile")

	var input models.Teacher
	if err := C.ShouldBindJSON(&input); err != nil {
		logger.C(C).Error("Invalid Input", zap.Error(err))
		response.Error(C, http.StatusBadRequest, "Invalid Input", err)
		return
	}

	teacher, err := services.CreateTeacherProfile(C, input)
	if err != nil {
		response.Error(C, http.StatusInternalServerError, "Failed to create teacher", err)
		return
	}

	logger.C(C).Info("Teacher Profile created successfully", zap.String("user_id", teacher.ID.String()))
	response.Success(C, http.StatusCreated, "Teacher created", gin.H{"teacher": teacher})
}

func UpdateTeacher(c *gin.Context, userId string) {
	logger.C(c).Info("UpdateTeacher endpoint hit")

	// Bind json
	var input models.UpdateTeacher
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.C(c).Error("Invalid Input", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "Invalid Input", err)
		return
	}

	// Update teacher
	err := services.UpdateTeacherUser(c, userId, input)
	if err != nil {
		logger.C(c).Error("Failed to update teacher", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "Failed to update teacher", err)
		return
	}

	logger.C(c).Info("Teacher updated successfully", zap.String("user_id", userId))
}
