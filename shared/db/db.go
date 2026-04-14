package db

import (
	"os"

	"github.com/sushantpardhi/shared/logger"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Get Database URL from environment variables
	dsn := os.Getenv("DATABASE_URL")

	logger.Info("Connecting to database...")

	// Connect to Database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("DB connection failed", zap.Error(err))
	}

	// Get DB instance
	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatal("Failed to get DB instance", zap.Error(err))
	}

	// Check if DB is reachable
	if err := sqlDB.Ping(); err != nil {
		logger.Fatal("DB not reachable", zap.Error(err))
	}

	// Set DB instance
	DB = db
	logger.Info("Database connected successfully")
}
