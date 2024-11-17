package models

import "gorm.io/gorm"

type Rate struct {
	gorm.Model
	DeliveryType  DeliveryType  `gorm:"not null"`
	TransportType TransportType `gorm:"not null"`
	Name          string        `gorm:"not null"`
	Price         float64       `gorm:"not null"`
	Description   string
}
