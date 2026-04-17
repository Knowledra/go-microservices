package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *Base) BeforeCreate(tx *gorm.DB) error {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return nil
}

type Admin struct {
	Base
	UserID uuid.UUID `gorm:"type:uuid;not null;unique"`
}

type Teacher struct {
	Base
	UserID                uuid.UUID  `gorm:"type:uuid;not null;unique" json:"user_id"`
	EmployeeID            string     `gorm:"type:varchar(50);uniqueIndex" json:"teacher_employee_id"`
	DepartmentID          uuid.UUID  `gorm:"type:uuid" json:"teacher_department_id"`
	DateOfBirth           *time.Time `gorm:"type:timestamp" json:"teacher_date_of_birth"`
	Gender                string     `json:"teacher_gender" gorm:"type:varchar(20)"`
	SubjectSpecialization string     `gorm:"type:varchar(255)" json:"teacher_subject_specialization"`
	BloodGroup            string     `gorm:"type:varchar(10)" json:"teacher_blood_group"`
	Qualification         string     `gorm:"type:varchar(255)" json:"teacher_qualification"`
	MobileNumber          string     `json:"teacher_mobile_number" gorm:"type:varchar(20)"`
	Address               string     `json:"teacher_address" gorm:"type:text"`
	JoiningDate           *time.Time `gorm:"type:timestamp" json:"teacher_joining_date"`
	EmploymentStatus      string     `gorm:"type:varchar(50);default:'active'" json:"teacher_employment_status"` // Full-time, Part-time, Contract
}

type Student struct {
	Base
	UserID         uuid.UUID       `json:"user_id" gorm:"type:uuid;not null;unique" `
	PRN            string          `json:"prn" gorm:"type:varchar(50);uniqueIndex" `
	MobileNumber   string          `json:"student_mobile_number" gorm:"type:varchar(20)" `
	DepartmentID   uuid.UUID       `json:"department_id" gorm:"type:uuid" `
	ClassID        uuid.UUID       `json:"student_class_id" gorm:"type:uuid" `
	SectionID      uuid.UUID       `json:"student_section_id" gorm:"type:uuid" `
	DateOfBirth    *time.Time      `json:"student_date_of_birth"`
	Gender         string          `json:"student_gender" gorm:"type:varchar(20)" `
	BloodGroup     string          `json:"student_blood_group" gorm:"type:varchar(10)" `
	EnrollmentDate *time.Time      `json:"student_enrollment_date"`
	StudentStatus  string          `json:"student_status" gorm:"type:varchar(50);default:'active'" ` // Active, Graduated
	Address        string          `json:"student_address" gorm:"type:text" `
	Parents        []Parent        `gorm:"many2many:student_parents;" json:"parents"`
	StudentParents []StudentParent `json:"student_parents" gorm:"constraint:OnDelete:CASCADE;"`
}

type Parent struct {
	Base
	UserID         uuid.UUID `json:"user_id" gorm:"type:uuid;not null;unique" `
	Occupation     string    `json:"parent_occupation" gorm:"type:varchar(100)" `
	PhonePrimary   string    `json:"parent_phone_primary" gorm:"type:varchar(20)" `
	PhoneSecondary string    `json:"parent_phone_secondary" gorm:"type:varchar(20)" `
	Address        string    `json:"parent_address" gorm:"type:text" `
	BloodGroup     string    `json:"parent_blood_group" gorm:"type:varchar(10)" `

	Students       []Student       `gorm:"many2many:student_parents;" json:"students"`
	StudentParents []StudentParent `json:"student_parents"`
}

type StudentParent struct {
	Base
	StudentID uuid.UUID `gorm:"type:uuid;index" json:"student_id"`
	ParentID  uuid.UUID `gorm:"type:uuid;index" json:"parent_id"`
	Relation  string    `gorm:"type:varchar(50)" json:"relation"` // Father, Mother, Guardian
	Student   *Student  `json:"student,omitempty"`
	Parent    *Parent   `json:"parent,omitempty"`
}
