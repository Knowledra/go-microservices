package models

import (
	"time"

	"github.com/google/uuid"
)

type Quiz struct {
	Base
	Title string

	ClassID      uuid.UUID `gorm:"type:uuid"`
	DepartmaneID uuid.UUID `gorm:"type:uuid"`
	SubjectID    uuid.UUID `gorm:"type:uuid"`
	TeacherID    uuid.UUID `gorm:"type:uuid"`

	TotalMarks int
	StartTime  time.Time
	EndTime    time.Time
}

type QuizQuestion struct {
	Base
	QuizID       uuid.UUID `gorm:"type:uuid;index"`
	QuestionText string
	QuestionType string // mcq / paragraph
	Marks        int
}

type QuizOption struct {
	Base
	QuestionID uuid.UUID `gorm:"type:uuid;index"`
	Text       string
	IsCorrect  bool
}

type QuizSubmission struct {
	Base
	QuizID    uuid.UUID `gorm:"type:uuid"`
	StudentID uuid.UUID `gorm:"type:uuid"`

	SubmittedAt time.Time
	TotalScore  float64
}

type QuizAnswer struct {
	Base
	SubmissionID uuid.UUID `gorm:"type:uuid"`
	QuestionID   uuid.UUID `gorm:"type:uuid"`

	SelectedOptionID *uuid.UUID `gorm:"type:uuid"`
	AnswerText       *string

	MarksAwarded float64
}
