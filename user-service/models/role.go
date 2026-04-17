package models

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"` // this should be same in the Auth DB and User DB
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Admin struct {
	Base
	Name     string    `gorm:"not null" json:"name"`
	LastName string    `gorm:"not null" json:"last_name"`
	UserID   uuid.UUID `gorm:"type:uuid;not null;unique"`
}
