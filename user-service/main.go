package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/env"
	"github.com/sushantpardhi/shared/logger"
	"github.com/sushantpardhi/shared/server"
	"go.uber.org/zap"
)

func main() {
	logger.Init("user-service")
	logger.Info("User service startup initiated")

	// Load Environment Variables
	logger.Info("Loading environment variables")
	if err := env.LoadEnv(""); err != nil {
		logger.Fatal("Failed to load environment variables", zap.Error(err))
	}

	// Check Environment Variables
	logger.Info("Validating environment variables")
	if err := env.CheckEnv(); err != nil {
		logger.Fatal("Environment validation failed", zap.Error(err))
	}

	// Set Gin Mode
	// if env.Environment == "production" {
	// 	gin.SetMode(gin.ReleaseMode)
	// 	logger.Info("Gin mode set to release")
	// } else {
	// 	logger.Info("Gin mode set to debug")
	// }

	app := gin.Default()

	db.ConnectDB()

	// routes.Init(app)

	port := os.Getenv("PORT")
	srv := server.New(app, port)
	srv.Start()
}
