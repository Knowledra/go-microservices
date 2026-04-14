package env

import (
	"github.com/joho/godotenv"
	"github.com/sushantpardhi/shared/logger"
	"go.uber.org/zap"
)

var Environment string

func LoadEnv(env string) error {
	// Load Local env
	if env == "local" {
		logger.Info("Loading local environment file", zap.String("file", ".env.local"))
		err := godotenv.Load(".env.local")
		Environment = "local"
		if err != nil {
			logger.Error("Failed to load local environment file", zap.Error(err))
			return err
		}
		logger.Info("Local environment loaded successfully")
	}

	// Load Development env
	if env == "development" {
		logger.Info("Loading development environment file", zap.String("file", ".env.development"))
		err := godotenv.Load(".env.development")
		Environment = "development"
		if err != nil {
			logger.Error("Failed to load development environment file", zap.Error(err))
			return err
		}
		logger.Info("Development environment loaded successfully")
	}

	// Load Production env
	if env == "production" || env == "" {
		logger.Info("Loading production environment file", zap.String("file", ".env"))
		err := godotenv.Load(".env")
		Environment = "production"
		if err != nil {
			logger.Error("Failed to load production environment file", zap.Error(err))
			return err
		}
		logger.Info("Production environment loaded successfully")
	}

	return nil
}
