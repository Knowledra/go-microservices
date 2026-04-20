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

func CreateParent(C *gin.Context) {
	logger.C(C).Info("Creating Parent Profile")

	var input models.Parent
	if err := C.ShouldBindJSON(&input); err != nil {
		logger.C(C).Error("Invalid Input", zap.Error(err))
		response.Error(C, http.StatusBadRequest, "Invalid Input", err)
		return
	}

	parent, err := services.CreateParentProfile(C, input)
	if err != nil {
		response.Error(C, http.StatusInternalServerError, "Failed to create parent", err)
		return
	}

	logger.C(C).Info("Parent Profile created successfully", zap.String("user_id", parent.ID.String()))
	response.Success(C, http.StatusCreated, "Parent created", gin.H{"parent": parent})
}

func UpdateParent(c *gin.Context, parentId string) {
	logger.C(c).Info("UpdateParent endpoint hit")

	// Bind json
	var input models.UpdateParent
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.C(c).Error("Invalid Input", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "Invalid Input", err)
		return
	}

	// Update parent
	err := services.UpdateParentUser(c, parentId, input)
	if err != nil {
		logger.C(c).Error("Failed to update parent", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "Failed to update parent", err)
		return
	}

	logger.C(c).Info("Parent updated successfully", zap.String("user_id", parentId))
}
