package controller

import (
	"email/models"
	"email/service"

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

	err := service.Send(req)
	if err != nil {
		logger.Error("Failed to send email", zap.String("recipient", req.To), zap.Error(err))
		response.Error(c, 500, "failed to send email", nil)
		return
	}

	logger.Info("Email sent request completed", zap.String("recipient", req.To))
	response.Success(c, 200, "email sent", nil)
}
