package models

import (
	"time"

	"github.com/google/uuid"
)

type Attendance struct {
	Base
	StudentID uuid.UUID `gorm:"type:uuid;index"`
	ClassID   uuid.UUID `gorm:"type:uuid"`
	Date      time.Time `gorm:"index"`
	Status    string    `gorm:"type:varchar(20)"` // present/absent
}
