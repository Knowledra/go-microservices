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

func CreateStudent(C *gin.Context) {
	logger.C(C).Info("Creating Student Profile")

	var input models.Student
	if err := C.ShouldBindJSON(&input); err != nil {
		logger.C(C).Error("Invalid Input", zap.Error(err))
		response.Error(C, http.StatusBadRequest, "Invalid Input", err)
		return
	}

	student, err := services.CreateStudentProfile(C, input)
	if err != nil {
		response.Error(C, http.StatusInternalServerError, "Failed to create student", err)
		return
	}

	logger.C(C).Info("Student Profile created successfully", zap.String("user_id", student.ID.String()))
	response.Success(C, http.StatusCreated, "Student created", gin.H{"student": student})
}

func UpdateStudent(C *gin.Context, userId string) {
	logger.C(C).Info("Updating Student Profile")

	var input models.UpdateStudent
	if err := C.ShouldBindJSON(&input); err != nil {
		logger.C(C).Error("Invalid Input", zap.Error(err))
		response.Error(C, http.StatusBadRequest, "Invalid Input", err)
		return
	}

	err := services.UpdateStudentUser(C, userId, input)
	if err != nil {
		response.Error(C, http.StatusInternalServerError, "Failed to update student", err)
		return
	}

	logger.C(C).Info("Student Profile updated successfully", zap.String("user_id", userId))
}
