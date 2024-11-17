package models

import (
	"gorm.io/gorm"
	"time"
)

type Chat struct {
	gorm.Model
	ParticipantID   uint            `gorm:"not null"` // ID клиента или курьера
	ParticipantType ParticipantType `gorm:"not null"` // Тип участника: "клиент" или "курьер"
	SupportStaffID  *uint           `gorm:"constraint:OnDelete:SET NULL"`
	Status          Status          `gorm:"not null"`
	CreationDate    time.Time
	Reason          string
	SupportStaff    SupportStaff `gorm:"foreignKey:SupportStaffID"`
}
