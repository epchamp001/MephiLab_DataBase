package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Role       Role   `gorm:"not null"` // Role: sender, receiver
	FirstName  string `gorm:"not null"`
	LastName   string `gorm:"not null"`
	MiddleName string
	Phone      string `gorm:"not null; unique"`
	Email      string `gorm:"not null"`
	Address    string
	Chats      []Chat      `gorm:"foreignKey:ClientID"`
	OrdersSent []Order     `gorm:"foreignKey:SenderID"`
	OrdersRecv []Order     `gorm:"foreignKey:RecipientID"`
	PromoCodes []PromoCode `gorm:"foreignKey:ClientID"`
}
