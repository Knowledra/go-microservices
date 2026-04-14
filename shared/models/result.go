package models

import (
	"github.com/google/uuid"
)

type Result struct {
	Base
	StudentID uuid.UUID `gorm:"type:uuid"`
	ExamID    uuid.UUID `gorm:"type:uuid"`
	SubjectID uuid.UUID `gorm:"type:uuid"`
	Marks     float64
}
