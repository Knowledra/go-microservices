package models

import (
	"time"

	"github.com/google/uuid"
)

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterUser struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Password string `json:"password"`
	Email    string `json:"email" binding:"required,email"`
	Role     string `json:"role" binding:"required"`
}

type RegisterAdmin struct {
	Name          string `json:"name" binding:"required"`
	LastName      string `json:"last_name" binding:"required"`
	Password      string `json:"password" binding:"required,min=8"`
	Email         string `json:"email" binding:"required,email"`
	Role          string `json:"role"`
	AdminPassword string `json:"admin_password" binding:"required"`
}

type RegisterStudent struct {
	Name           string     `json:"student_name" binding:"required"`
	LastName       string     `json:"student_last_name" binding:"required"`
	Email          string     `json:"student_email" binding:"required,email"`
	Password       string     `json:"student_password"` // Generating temp Pass
	Role           string     `json:"role"`             // Always "student"
	DepartmentID   uuid.UUID  `json:"student_department_id" gorm:"type:uuid"`
	ClassID        uuid.UUID  `json:"student_class_id" gorm:"type:uuid"`
	DateOfBirth    *time.Time `json:"student_date_of_birth"`
	Gender         string     `json:"student_gender" gorm:"type:varchar(20)"`
	BloodGroup     string     `json:"student_blood_group" gorm:"type:varchar(10)"`
	MobileNumber   string     `json:"student_mobile_number" gorm:"type:varchar(20)"`
	Address        string     `json:"student_address" gorm:"type:text"`
	EnrollmentDate *time.Time `json:"enrollment_date"`

	// Parent Info
	ParentName       string `json:"parent_name" binding:"required"`
	ParentLastName   string `json:"parent_last_name" binding:"required"`
	ParentEmail      string `json:"parent_email" binding:"required,email"`
	ParentOccupation string `json:"parent_occupation"`
	ParentRelation   string `json:"parent_relation" binding:"required"`
	PhonePrimary     string `json:"parent_phone_primary" gorm:"type:varchar(20)"`
	PhoneSecondary   string `json:"parent_phone_secondary" gorm:"type:varchar(20)"`
	ParentAddress    string `json:"parent_address" gorm:"type:text"`
	ParentBloodGroup string `json:"parent_blood_group" gorm:"type:varchar(10)"`
}

type RegisterTeacher struct {
	Name                  string     `json:"teacher_name" binding:"required"`
	LastName              string     `json:"teacher_last_name" binding:"required"`
	Email                 string     `json:"teacher_email" binding:"required,email"`
	Password              string     `json:"teacher_password"` // Generating temp Pass
	Role                  string     `json:"role"`             // Always "teacher"
	DepartmentID          uuid.UUID  `json:"teacher_department_id" gorm:"type:uuid"`
	DateOfBirth           *time.Time `json:"teacher_date_of_birth"`
	Gender                string     `json:"teacher_gender" gorm:"type:varchar(20)"`
	BloodGroup            string     `json:"teacher_blood_group" gorm:"type:varchar(10)"`
	MobileNumber          string     `json:"teacher_mobile_number" gorm:"type:varchar(20)"`
	Address               string     `json:"teacher_address" gorm:"type:text"`
	Qualification         string     `json:"teacher_qualification" gorm:"type:varchar(255)" `
	JoiningDate           *time.Time `json:"teacher_joining_date"`
	SubjectSpecialization string     `gorm:"type:varchar(255)" json:"teacher_subject_specialization"`
}
