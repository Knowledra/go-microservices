package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Department struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name string    `gorm:"not null" json:"name"`
	Code string    `gorm:"not null;unique" json:"code"`
}

func (d *Department) BeforeCreate(tx *gorm.DB) error {
	if d.ID == uuid.Nil {
		d.ID = uuid.New()
	}
	return nil
}
