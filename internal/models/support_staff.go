package models

import "gorm.io/gorm"

type SupportStaff struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Phone     string `gorm:"not null"`
	Email     string `gorm:"not null"`
	Position  string
	Chats     []Chat `gorm:"foreignKey:SupportStaffID"` // One-to-Many связь с Chat
}
