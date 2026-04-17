package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"not null" json:"name"`
	LastName  string    `gorm:"not null" json:"last_name"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	Role      string    `gorm:"not null" json:"role"` // admin, teacher, student, parent
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

type RegisterAdmin struct {
	Name          string `json:"name" binding:"required"`
	LastName      string `json:"last_name" binding:"required"`
	Password      string `json:"password" binding:"required,min=8"`
	Email         string `json:"email" binding:"required,email"`
	Role          string `json:"role"`
	AdminPassword string `json:"admin_password" binding:"required"`
}
