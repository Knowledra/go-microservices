package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/env"
	"github.com/sushantpardhi/shared/logger"
	"github.com/sushantpardhi/shared/middleware"
	"github.com/sushantpardhi/shared/server"
)

func main() {
	logger.Init("quiz-service")
	logger.Info("Quiz service startup initiated")

	// Load Environment Variables
	logger.Info("Loading environment variables")
	if err := env.LoadEnv(); err != nil {
		logger.Fatal("Failed to load environment variables", zap.Error(err))
	}

	// Check Environment Variables
	logger.Info("Validating environment variables")
	if err := env.CheckEnv("PORT", "DATABASE_URL"); err != nil {
		logger.Fatal("Environment validation failed", zap.Error(err))
	}

	// Set Gin Mode
	Environment := os.Getenv("APP_ENV")
	if Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
		logger.Info("Gin mode set to release")
	} else {
		gin.SetMode(gin.DebugMode)
		logger.Info("Gin mode set to debug")
	}

	app := gin.Default()
	app.Use(middleware.RequestID())

	db.ConnectDB()

	// if err := models.Migrate(db.DB); err != nil {
	// 	log.Fatal("Migration failed:", err)
	// }

	// routes.Init(app)

	port := os.Getenv("PORT")
	srv := server.New(app, port)
	srv.Start()
}
