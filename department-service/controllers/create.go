package controllers

import (
	"department/models"
	"department/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/logger"
	"github.com/sushantpardhi/shared/response"
	"go.uber.org/zap"
)

func CreateDepartment(c *gin.Context) {
	logger.C(c).Info("Creating Department")

	var input models.Department
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.C(c).Error("Invalid Input", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "Invalid Input", err)
		return
	}

	department, err := services.CreateDepartment(c, input)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create department", err)
		return
	}

	logger.C(c).Info("Department created successfully", zap.String("department_id", department.ID.String()))
	response.Success(c, http.StatusCreated, "Department created", gin.H{"department": department})
}
