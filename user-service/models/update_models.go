package models

import (
	"time"

	"github.com/google/uuid"
)

type UpdateAdmin struct {
	Name     string `json:"name" binding:"omitempty"`
	LastName string `json:"last_name" binding:"omitempty"`
}

type UpdateTeacher struct {
	Name           string `json:"name" binding:"omitempty"`
	LastName       string `json:"last_name" binding:"omitempty"`
	Specialization string `gorm:"type:varchar(255)" json:"specialization" binding:"omitempty"`
}

type UpdateStudent struct {
	Name      string    `json:"name" binding:"omitempty"`
	LastName  string    `json:"last_name" binding:"omitempty"`
	Class     string    `gorm:"type:varchar(100)" json:"class" binding:"omitempty"`
	Parent    *Parent   `gorm:"foreignKey:ParentID" json:"parent,omitempty" binding:"omitempty"`
	ParentID  uuid.UUID `gorm:"type:uuid;index" json:"parent_id" binding:"omitempty"`
	Relation  string    `gorm:"type:varchar(100)" json:"relation" binding:"omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateParent struct {
	Name     string `json:"name" binding:"omitempty"`
	LastName string `json:"last_name" binding:"omitempty"`
}
