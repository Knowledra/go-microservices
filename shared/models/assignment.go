package models

import (
	"time"

	"github.com/google/uuid"
)

// Assignment Tables
type Assignment struct {
	Base
	Title       string
	Description string

	ClassID      uuid.UUID `gorm:"type:uuid"`
	DepartmaneID uuid.UUID `gorm:"type:uuid"`
	SubjectID    uuid.UUID `gorm:"type:uuid"`
	TeacherID    uuid.UUID `gorm:"type:uuid"`

	DueDate time.Time
}

// Join Table
type AssignmentSubmission struct {
	Base
	AssignmentID uuid.UUID `gorm:"type:uuid"`
	StudentID    uuid.UUID `gorm:"type:uuid"`

	FileURL     string
	SubmittedAt time.Time

	Grade    float64
	Feedback string
}
