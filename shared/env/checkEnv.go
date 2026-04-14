package env

import (
	"fmt"
	"os"

	"github.com/sushantpardhi/shared/logger"
	"go.uber.org/zap"
)

func CheckEnv() error {
	logger.Info("Checking required environment variables")

	// Check if PORT is set
	if os.Getenv("PORT") == "" {
		logger.Error("Required environment variable not set", zap.String("variable", "PORT"))
		return fmt.Errorf("PORT is not set")
	}

	// Check if DATABASE_URL is set
	if os.Getenv("DATABASE_URL") == "" {
		logger.Error("Required environment variable not set", zap.String("variable", "DATABASE_URL"))
		return fmt.Errorf("DATABASE_URL is not set")
	}

	// Check if ADMIN_PASS is set
	if os.Getenv("ADMIN_PASS") == "" {
		logger.Error("Required environment variable not set", zap.String("variable", "ADMIN_PASS"))
		return fmt.Errorf("ADMIN_PASS is not set")
	}

	// Check if ACCESS_TOKEN_SECRET is set
	if os.Getenv("ACCESS_TOKEN_SECRET") == "" {
		logger.Error("Required environment variable not set", zap.String("variable", "ACCESS_TOKEN_SECRET"))
		return fmt.Errorf("ACCESS_TOKEN_SECRET is not set")
	}

	logger.Info("All required environment variables are set")
	return nil
}
