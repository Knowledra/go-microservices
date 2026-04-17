package env

import (
	"fmt"
	"os"

	"github.com/sushantpardhi/shared/logger"
	"go.uber.org/zap"
)

func CheckEnv(envs ...string) error {
	logger.Info("Checking required environment variables")

	for _, key := range envs {
		if os.Getenv(key) == "" {
			logger.Error("Required environment variable not set", zap.String("variable", key))
			return fmt.Errorf("%s is not set", key)
		}
	}

	logger.Info("All required environment variables are set")
	return nil
}
