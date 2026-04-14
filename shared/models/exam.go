package models

import (
	"time"

	"github.com/google/uuid"
)

type Exam struct {
	Base
	Name      string
	StartDate time.Time
}

type ExamSubject struct {
	Base
	ExamID    uuid.UUID `gorm:"type:uuid"`
	SubjectID uuid.UUID `gorm:"type:uuid"`
	MaxMarks  int
}
