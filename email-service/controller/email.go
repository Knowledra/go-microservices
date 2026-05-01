package controller

import (
	"email/models"
	"email/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/logger"
	"github.com/sushantpardhi/shared/response"
	"go.uber.org/zap"
)

func SendEmail(c *gin.Context) {
	var req models.EmailRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("Invalid send email request", zap.Error(err))
		response.Error(c, 400, "invalid request payload", nil)
		return
	}

	go func(request models.EmailRequest) {
		if err := service.Send(request); err != nil {
			logger.Error("Background email send failed", zap.String("recipient", request.To), zap.Error(err))
			return
		}

		logger.Info("Background email send succeeded", zap.String("recipient", request.To))
	}(req)

	logger.Info("Email queued for background sending", zap.String("recipient", req.To))
	c.JSON(http.StatusAccepted, gin.H{
		"success": true,
		"code":    http.StatusAccepted,
		"message": "email queued for sending",
		"data":    nil,
	})
}
