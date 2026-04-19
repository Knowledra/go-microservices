package models

import (
	"time"

	"github.com/google/uuid"
)

type Admin struct {
	ID        uuid.UUID `gorm:"type:uuid;not null;unique" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	LastName  string    `gorm:"not null" json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Teacher struct {
	ID             uuid.UUID `gorm:"type:uuid;not null;unique" json:"id"`
	Name           string    `gorm:"not null" json:"name"`
	LastName       string    `gorm:"not null" json:"last_name"`
	Specialization string    `gorm:"type:varchar(255)" json:"specialization"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Student struct {
	ID        uuid.UUID `gorm:"type:uuid;not null;unique" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	LastName  string    `gorm:"not null" json:"last_name"`
	Class     string    `gorm:"type:varchar(100)" json:"class"`
	ParentID  uuid.UUID `gorm:"type:uuid;index" json:"parent_id"`
	Parent    *Parent   `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Parent struct {
	ID          uuid.UUID `gorm:"type:uuid;not null;unique" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	LastName    string    `gorm:"not null" json:"last_name"`
	PhoneNumber string    `gorm:"type:varchar(20)" json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateAdmin struct {
	Name     string `json:"name" binding:"omitempty"`
	LastName string `json:"last_name" binding:"omitempty"`
}
