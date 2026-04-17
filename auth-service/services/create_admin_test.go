package services

import (
	"auth/models"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/logger"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setup test DB
func setupTestDB(t *testing.T) {
	logger.Init("test")

	database, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// migrate schema
	err = database.AutoMigrate(&models.User{})
	assert.NoError(t, err)

	db.DB = database
}

func TestCreateAdmin_Success(t *testing.T) {
	setupTestDB(t)

	os.Setenv("ADMIN_PASS", "secret")

	input := models.RegisterAdmin{
		Name:          "John",
		LastName:      "Doe",
		Email:         "john@example.com",
		Password:      "password123",
		AdminPassword: "secret",
	}

	user, err := CreateAdminUser(input)

	assert.NoError(t, err)
	assert.Equal(t, "admin", user.Role)
	assert.Equal(t, input.Email, user.Email)
}

func TestCreateAdmin_InvalidAdminPassword(t *testing.T) {
	setupTestDB(t)

	os.Setenv("ADMIN_PASS", "secret")

	input := models.RegisterAdmin{
		Email:         "john@example.com",
		Password:      "password123",
		AdminPassword: "wrong",
	}

	_, err := CreateAdminUser(input)

	assert.Error(t, err)
	assert.Equal(t, "invalid admin credentials", err.Error())
}

func TestCreateAdmin_UserAlreadyExists(t *testing.T) {
	setupTestDB(t)

	os.Setenv("ADMIN_PASS", "secret")

	// insert existing user
	db.DB.Create(&models.User{
		Email: "john@example.com",
	})

	input := models.RegisterAdmin{
		Email:         "john@example.com",
		Password:      "password123",
		AdminPassword: "secret",
	}

	_, err := CreateAdminUser(input)

	assert.Error(t, err)
	assert.Equal(t, "user already exists", err.Error())
}

func TestCreateAdmin_DBFailure(t *testing.T) {
	setupTestDB(t)

	os.Setenv("ADMIN_PASS", "secret")

	// close DB to simulate failure
	sqlDB, _ := db.DB.DB()
	sqlDB.Close()

	input := models.RegisterAdmin{
		Email:         "john@example.com",
		Password:      "password123",
		AdminPassword: "secret",
	}

	_, err := CreateAdminUser(input)

	assert.Error(t, err)
}
