package models

type SupportStaff struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Phone     string `gorm:"not null"`
	Email     string `gorm:"not null; unique"`
	Position  string
	Chats     []Chat `gorm:"foreignKey:SupportStaffID"` // One-to-Many связь с Chat
}
