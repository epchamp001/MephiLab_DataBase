package models

import "time"

type PromoCode struct {
	ID             uint          `gorm:"primaryKey"`
	Type           PromoCodeType `gorm:"not null"`
	ClientID       *uint         `gorm:"constraint:OnDelete:SET NULL"`
	Code           string        `gorm:"not null; unique"`
	DiscountAmount float64       `gorm:"not null"`
	ValidUntil     time.Time
	Personalized   bool
	Client         Client `gorm:"foreignKey:ClientID"`
}
