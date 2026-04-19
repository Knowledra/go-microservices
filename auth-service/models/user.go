package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (u *AuthUser) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

type AuthUser struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Email     string         `json:"email" binding:"required,email"`
	Password  string         `json:"password" binding:"required,min=8"`
	Role      string         `json:"role" binding:"required"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	IsDeleted bool           `json:"is_deleted" gorm:"default:false"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type RegisterAdmin struct {
	AuthUser
	Name          string `json:"name" binding:"required"`
	LastName      string `json:"last_name" binding:"required"`
	AdminPassword string `json:"admin_password" binding:"required"`
}

type RegisterTeacher struct {
	AuthUser
	Name           string `json:"name" binding:"required"`
	LastName       string `json:"last_name" binding:"required"`
	Specialization string `json:"specialization" binding:"required"`
}

type RegisterStudent struct {
	AuthUser
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Class    string `json:"class" binding:"required"`
}

type RegisterParent struct {
	AuthUser
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
}
