package services

import (
	"bytes"
	"context"
	"text/template"

	"github.com/sushantpardhi/shared/email"
	"github.com/sushantpardhi/shared/logger"
	sharedModels "github.com/sushantpardhi/shared/models"
	"go.uber.org/zap"
)

func sendWelcomeEmailAsync(c context.Context, templatePath string, subject string, data map[string]string) {
	go func() {
		tmpl, err := template.ParseFiles(templatePath)
		if err != nil {
			logger.C(c).Error("Failed to parse welcome template", zap.String("template", templatePath), zap.Error(err))
			return
		}

		var htmlBody bytes.Buffer
		if err := tmpl.Execute(&htmlBody, data); err != nil {
			logger.C(c).Error("Failed to render welcome template", zap.String("template", templatePath), zap.Error(err))
			return
		}

		emailReq := sharedModels.EmailRequest{
			To:      data["Email"],
			Subject: subject,
			Body:    htmlBody.String(),
			IsHTML:  true,
		}

		if err := email.SendEmail(c, emailReq); err != nil {
			logger.C(c).Error("Failed to send welcome email", zap.String("email", data["Email"]), zap.Error(err))
			return
		}

		logger.C(c).Info("Welcome email sent", zap.String("email", data["Email"]), zap.String("template", templatePath))
	}()
}
