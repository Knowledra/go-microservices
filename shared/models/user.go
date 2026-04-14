package models

type User struct {
	Base
	Name     string `gorm:"not null" json:"name"`
	LastName string `gorm:"not null" json:"last_name"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
	Role     string `gorm:"not null" json:"role"` // admin, teacher, student, parent
	IsActive bool   `gorm:"default:true" json:"is_active"`

	Admin   *Admin   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"admin,omitempty"`
	Teacher *Teacher `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"teacher,omitempty"`
	Student *Student `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"student,omitempty"`
	Parent  *Parent  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"parent,omitempty"`
}
