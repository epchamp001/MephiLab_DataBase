package models

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	gorm.Model
	ChatID     uint       `gorm:"not null;constraint:OnDelete:CASCADE"`
	SenderType SenderType `gorm:"not null"`
	SenderID   uint       `gorm:"not null"`
	Timestamp  time.Time
	Text       string `gorm:"not null"`
	Chat       Chat   `gorm:"foreignKey:ChatID"`
}
