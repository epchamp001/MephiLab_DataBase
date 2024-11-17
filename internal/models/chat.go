package models

import "time"

type Chat struct {
	ID              uint            `gorm:"primaryKey"`
	ParticipantType ParticipantType `gorm:"not null"`
	ParticipantID   uint            `gorm:"not null"`
	SupportStaffID  *uint           `gorm:"constraint:OnDelete:SET NULL"`
	Status          Status          `gorm:"not null"`
	CreationDate    time.Time
	Reason          string
	SupportStaff    SupportStaff `gorm:"foreignKey:SupportStaffID"`
}
