package models

import "github.com/google/uuid"

// Academic Tables
type Department struct {
	Base

	Name     string     `json:"name" gorm:"not null" binding:"required"`    // e.g. Computer Science, Electronics and Telecommunication
	Code     string     `json:"code" gorm:"uniqueIndex" binding:"required"` // e.g. CS, ENTC
	HOD      *uuid.UUID `json:"hod" gorm:"type:uuid"`                       // teacher/admin as HOD
	IsActive bool       `json:"is_active" gorm:"default:true"`

	Subjects []Subject `json:"subjects" gorm:"foreignKey:DepartmentID;constraint:OnDelete:CASCADE"`
	Classes  []Class   `json:"classes" gorm:"foreignKey:DepartmentID;constraint:OnDelete:CASCADE"`
}

type Class struct {
	Base
	Name         string    `json:"name" gorm:"not null" binding:"required"`    // e.g. First Year or 10th Grade
	Code         string    `json:"code" gorm:"uniqueIndex" binding:"required"` // e.g. FY or 10th
	DepartmentID uuid.UUID `json:"department_id" gorm:"type:uuid" binding:"required"`

	Sections []Section `json:"sections" gorm:"foreignKey:ClassID;constraint:OnDelete:CASCADE"`
}

type Section struct {
	Base
	Name    string    `json:"name" gorm:"not null"` // e.g. A, B, C
	ClassID uuid.UUID `json:"class_id" gorm:"type:uuid"`
}

type Subject struct {
	Base
	Name         string    `json:"name" gorm:"not null"`    // e.g. Data Structures, Operating Systems
	Code         string    `json:"code" gorm:"uniqueIndex"` // e.g. CS201, CS202
	Description  string    `json:"description"`
	DepartmentID uuid.UUID `json:"department_id" gorm:"type:uuid" binding:"required"`
}

// DepartmentSummary for listing all departments with counts
type DepartmentSummary struct {
	ID           uuid.UUID  `json:"ID"`
	CreatedAt    string     `json:"CreatedAt"`
	UpdatedAt    string     `json:"UpdatedAt"`
	DeletedAt    *string    `json:"DeletedAt"`
	Name         string     `json:"name"`
	Code         string     `json:"code"`
	HOD          *uuid.UUID `json:"hod"`
	IsActive     bool       `json:"is_active"`
	SubjectCount int        `json:"subject_count"`
	ClassCount   int        `json:"class_count"`
}

// Join Tables
// type DepartmentClass struct {
// 	Base
// 	DepartmentID uuid.UUID `json:"department_id" gorm:"type:uuid"`
// 	ClassID      uuid.UUID `json:"class_id" gorm:"type:uuid"`
// }

// type ClassSection struct {
// 	Base
// 	ClassID   uuid.UUID  `json:"class_id" gorm:"type:uuid"`
// 	SectionID uuid.UUID  `json:"section_id" gorm:"type:uuid"`
// 	MentorID  *uuid.UUID `json:"mentor_id" gorm:"type:uuid"` // teacher as mentor for the class
// }

// type ClassSubject struct {
// 	Base
// 	ClassID   uuid.UUID `json:"class_id" gorm:"type:uuid"`
// 	SubjectID uuid.UUID `json:"subject_id" gorm:"type:uuid"`
// 	TeacherID uuid.UUID `json:"teacher_id" gorm:"type:uuid"`
// }

type TeacherSubject struct {
	Base
	TeacherID uuid.UUID `json:"teacher_id" gorm:"type:uuid"`
	SubjectID uuid.UUID `json:"subject_id" gorm:"type:uuid"` // dept will be taken from subject
	ClassID   uuid.UUID `json:"class_id" gorm:"type:uuid"`   // for teachers teaching specific classes
	SectionID uuid.UUID `json:"section_id" gorm:"type:uuid"` // for teachers teaching specific sections
}

// type DepartmentSubject struct {
// 	Base
// 	DepartmentID uuid.UUID `json:"department_id" gorm:"type:uuid"`
// 	SubjectID    uuid.UUID `json:"subject_id" gorm:"type:uuid"`
// }

type ClassTeacher struct {
	Base
	ClassID   uuid.UUID `json:"class_id" gorm:"type:uuid"`
	TeacherID uuid.UUID `json:"teacher_id" gorm:"type:uuid"`
}
