package models

import "time"

type Order struct {
	ID                 uint          `gorm:"primaryKey"`
	Urgency            Urgency       `gorm:"not null"`
	SenderID           uint          `gorm:"not null;constraint:OnDelete:CASCADE"`
	RecipientID        uint          `gorm:"not null;constraint:OnDelete:CASCADE"`
	CourierID          uint          `gorm:"not null;constraint:OnDelete:SET NULL"`
	CurrentStatus      CurrentStatus `gorm:"not null"`
	PromoCodeID        *uint         `gorm:"constraint:OnDelete:SET NULL"`
	PaymentStatus      PaymentStatus `gorm:"not null"`
	RateID             uint          `gorm:"not null;constraint:OnDelete:SET NULL"`
	CreationDate       time.Time
	ItemType           string
	ItemValue          float64
	Weight             float64
	DiscountSurcharges float64
	PaymentMethod      string
	Sender             Client    `gorm:"foreignKey:SenderID"`
	Recipient          Client    `gorm:"foreignKey:RecipientID"`
	Courier            Courier   `gorm:"foreignKey:CourierID"`
	PromoCode          PromoCode `gorm:"foreignKey:PromoCodeID"`
	Rate               Rate      `gorm:"foreignKey:RateID"`
}
