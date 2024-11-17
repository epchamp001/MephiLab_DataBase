package models

import (
	"gorm.io/gorm"
	"time"
)

type Chat struct {
	gorm.Model
	ParticipantType ParticipantType `gorm:"not null"`
	ClientID        *uint           // NULL, если участник - курьер
	CourierID       *uint           // NULL, если участник - клиент
	SupportStaffID  *uint           `gorm:"constraint:OnDelete:SET NULL"`
	Status          Status          `gorm:"not null"`
	CreationDate    time.Time
	Reason          string
	SupportStaff    SupportStaff `gorm:"foreignKey:SupportStaffID"`
}
