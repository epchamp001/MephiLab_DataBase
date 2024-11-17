package models

import (
	"gorm.io/gorm"
	"time"
)

type PromoCode struct {
	gorm.Model
	Type           PromoCodeType `gorm:"not null"`
	ClientID       *uint         `gorm:"constraint:OnDelete:SET NULL"`
	Code           string        `gorm:"not null; unique"`
	DiscountAmount float64       `gorm:"not null"`
	ValidUntil     time.Time
	Personalized   bool
	Client         Client `gorm:"foreignKey:ClientID"`
}
