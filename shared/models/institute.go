package models

type Institute struct {
	Base
	Name      string `gorm:"not null"`
	Subdomain string `gorm:"uniqueIndex;not null"`
	Status    string
}
