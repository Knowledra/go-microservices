package services

import (
	"auth/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sushantpardhi/shared/db"
	"github.com/sushantpardhi/shared/logger"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupLoginTestDB(t *testing.T) {
	logger.Init("test")

	database, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = database.AutoMigrate(&models.User{})
	assert.NoError(t, err)

	db.DB = database
}

func TestLogin_Success(t *testing.T) {
	setupLoginTestDB(t)

	hashed, _ := bcrypt.GenerateFromPassword([]byte("password123"), 10)

	user := models.User{
		Email:    "john@example.com",
		Password: string(hashed),
		Role:     "admin",
	}

	db.DB.Create(&user)

	token, resultUser, err := Login("john@example.com", "password123")

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.Equal(t, user.Email, resultUser.Email)
}

func TestLogin_UserNotFound(t *testing.T) {
	setupLoginTestDB(t)

	token, user, err := Login("notfound@example.com", "password123")

	assert.Error(t, err)
	assert.Empty(t, token)
	assert.Empty(t, user.ID)
}

func TestLogin_WrongPassword(t *testing.T) {
	setupLoginTestDB(t)

	hashed, _ := bcrypt.GenerateFromPassword([]byte("password123"), 10)

	db.DB.Create(&models.User{
		Email:    "john@example.com",
		Password: string(hashed),
		Role:     "admin",
	})

	token, _, err := Login("john@example.com", "wrongpass")

	assert.Error(t, err)
	assert.Empty(t, token)
}
