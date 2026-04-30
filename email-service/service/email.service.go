package service

import (
	"crypto/tls"
	"email/models"
	"fmt"
	"net/smtp"
	"os"

	"github.com/sushantpardhi/shared/logger"
	"go.uber.org/zap"
)

func Send(req models.EmailRequest) error {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")

	if req.To == "" || req.Subject == "" || req.Body == "" {
		err := fmt.Errorf("missing required email fields")
		logger.Warn("Email request validation failed",
			zap.Bool("has_to", req.To != ""),
			zap.Bool("has_subject", req.Subject != ""),
			zap.Bool("has_body", req.Body != ""),
		)
		return err
	}

	if host == "" || port == "" || user == "" || pass == "" {
		err := fmt.Errorf("missing SMTP configuration")
		logger.Error("SMTP configuration is incomplete",
			zap.Bool("has_host", host != ""),
			zap.Bool("has_port", port != ""),
			zap.Bool("has_user", user != ""),
			zap.Bool("has_pass", pass != ""),
		)
		return err
	}

	addr := host + ":" + port
	logger.Info("Sending email over SMTP TLS", zap.String("recipient", req.To), zap.String("smtp_addr", addr))

	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		logger.Error("SMTP TLS dial failed", zap.String("smtp_addr", addr), zap.Error(err))
		return fmt.Errorf("smtp tls dial failed: %w", err)
	}
	defer func() {
		if closeErr := conn.Close(); closeErr != nil {
			logger.Warn("Failed to close SMTP TLS connection", zap.Error(closeErr))
		}
	}()

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		logger.Error("Failed to create SMTP client", zap.Error(err))
		return fmt.Errorf("create smtp client failed: %w", err)
	}
	defer func() {
		if quitErr := client.Quit(); quitErr != nil {
			logger.Warn("SMTP quit returned error", zap.Error(quitErr))
		}
	}()

	// Auth
	auth := smtp.PlainAuth("", user, pass, host)
	if err = client.Auth(auth); err != nil {
		logger.Error("SMTP authentication failed", zap.String("smtp_host", host), zap.Error(err))
		return fmt.Errorf("smtp auth failed: %w", err)
	}

	// Set sender
	if err = client.Mail(user); err != nil {
		logger.Error("Failed to set sender", zap.String("sender", user), zap.Error(err))
		return fmt.Errorf("smtp set sender failed: %w", err)
	}

	// Set recipient
	if err = client.Rcpt(req.To); err != nil {
		logger.Error("Failed to set recipient", zap.String("recipient", req.To), zap.Error(err))
		return fmt.Errorf("smtp set recipient failed: %w", err)
	}

	// Write email
	w, err := client.Data()
	if err != nil {
		logger.Error("Failed to initialize SMTP data writer", zap.Error(err))
		return fmt.Errorf("smtp data command failed: %w", err)
	}

	// Set Content-Type based on whether body is HTML
	contentType := "text/plain; charset=utf-8"
	if req.IsHTML {
		contentType = "text/html; charset=utf-8"
		logger.Info("Sending email as HTML", zap.String("recipient", req.To))
	}

	msg := fmt.Sprintf(
		"To: %s\r\nSubject: %s\r\nContent-Type: %s\r\n\r\n%s",
		req.To, req.Subject, contentType, req.Body,
	)

	_, err = w.Write([]byte(msg))
	if err != nil {
		logger.Error("Failed to write email message", zap.String("recipient", req.To), zap.Error(err))
		return fmt.Errorf("smtp write message failed: %w", err)
	}

	err = w.Close()
	if err != nil {
		logger.Error("Failed to finalize email message", zap.String("recipient", req.To), zap.Error(err))
		return fmt.Errorf("smtp close message writer failed: %w", err)
	}

	logger.Info("Email sent successfully over SMTP TLS", zap.String("recipient", req.To))
	return nil
}
