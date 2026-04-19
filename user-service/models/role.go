package models

import (
	"time"

	"github.com/google/uuid"
)

type Admin struct {
	ID        uuid.UUID `gorm:"type:uuid;not null;unique" json:"user_id"` // will not be generated, instead taken from the auth service
	Name      string    `gorm:"not null" json:"name"`
	LastName  string    `gorm:"not null" json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateAdmin struct {
	Name     string `json:"name" binding:"omitempty"`
	LastName string `json:"last_name" binding:"omitempty"`
}
