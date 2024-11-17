package models

type Rate struct {
	ID            uint          `gorm:"primaryKey"`
	DeliveryType  DeliveryType  `gorm:"not null"`
	TransportType TransportType `gorm:"not null"`
	Name          string        `gorm:"not null"`
	Price         float64       `gorm:"not null"`
	Description   string
}
