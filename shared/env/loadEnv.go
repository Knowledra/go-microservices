package env

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sushantpardhi/shared/logger"
	"go.uber.org/zap"
)

var Environment string

func LoadEnv() error {
	Environment = os.Getenv("APP_ENV")
	if Environment == "" {
		Environment = "development"
	}

	envFile := ".env." + Environment

	logger.Info("Loading environment file", zap.String("file", envFile))
	err := godotenv.Load(envFile)
	if err != nil {
		logger.Warn("Failed to load environment file. Relying on system environment variables.", zap.Error(err))
	} else {
		logger.Info("Environment loaded successfully", zap.String("env", Environment))
	}

	return nil
}
