package db

import (
	"os"
	"time"

	"github.com/sushantpardhi/shared/logger"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Get Database URL from environment variables
	db_url := os.Getenv("DATABASE_URL")

	logger.Info("Connecting to database...")

	// Connect to Database
	db, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})
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

	// Configure connection pool for efficient reuse
	sqlDB.SetMaxIdleConns(10)                 // Keep up to 10 idle connections ready
	sqlDB.SetMaxOpenConns(100)                // Limit total open connections to 100
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // Recycle connections every 5 minutes

	// Set DB instance
	DB = db
	logger.Info("Database connected successfully")
}
