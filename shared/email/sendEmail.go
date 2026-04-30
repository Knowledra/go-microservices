package email

import (
	"context"
	"fmt"
	"os"

	"github.com/sushantpardhi/shared/callAPI"
	"github.com/sushantpardhi/shared/logger"
	"github.com/sushantpardhi/shared/models"
	"go.uber.org/zap"
)

// SendEmail sends an email by calling the email-service API.
// It accepts an EmailRequest struct with To, Subject, Body, and IsHTML fields.
// This method is reusable across all services.
// It reads EMAIL_SERVICE_URL from environment variables, with fallback to localhost for development.
func SendEmail(ctx context.Context, req models.EmailRequest) error {
	emailServiceURL := os.Getenv("EMAIL_SERVICE_URL")
	if emailServiceURL == "" {
		emailServiceURL = "http://localhost:8082"
		logger.Warn("EMAIL_SERVICE_URL not set, using default localhost URL")
	}

	endpoint := emailServiceURL + "/send-email"

	logger.Info("Sending email via email-service",
		zap.String("recipient", req.To),
		zap.String("subject", req.Subject),
		zap.Bool("isHtml", req.IsHTML),
		zap.String("endpoint", endpoint),
	)

	_, err := callAPI.CallAPI(ctx, "POST", endpoint, req)
	if err != nil {
		logger.Error("Failed to send email",
			zap.String("recipient", req.To),
			zap.Error(err),
		)
		return fmt.Errorf("failed to send email to %s: %w", req.To, err)
	}

	logger.Info("Email sent successfully",
		zap.String("recipient", req.To),
	)
	return nil
}
