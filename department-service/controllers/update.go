package controllers

import (
	"department/models"
	"department/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/logger"
	"github.com/sushantpardhi/shared/response"
	"go.uber.org/zap"
)

func UpdateDepartment(c *gin.Context) {
	logger.C(c).Info("Updating Department")

	departmentID := c.Param("department_id")
	if departmentID == "" {
		response.Error(c, http.StatusBadRequest, "department_id is required", errors.New("department_id is required"))
		return
	}

	var input models.UpdateDepartment
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.C(c).Error("Invalid Input", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "Invalid Input", err)
		return
	}

	updatedDepartment, err := services.UpdateDepartment(c, departmentID, input)
	if err != nil {
		switch err.Error() {
		case "department not found":
			response.Error(c, http.StatusNotFound, "Department not found", err)
		case "no fields provided to update":
			response.Error(c, http.StatusBadRequest, "No fields provided to update", err)
		default:
			response.Error(c, http.StatusInternalServerError, "Failed to update department", err)
		}
		return
	}

	logger.C(c).Info("Department updated successfully", zap.String("department_id", updatedDepartment.ID.String()))
	response.Success(c, http.StatusOK, "Department updated", gin.H{"department": updatedDepartment})

}
