package models

type Courier struct {
	ID                 uint               `gorm:"primaryKey"`
	EmploymentStatus   EmploymentStatus   `gorm:"not null"`
	TransportType      TransportType      `gorm:"not null"`
	AvailabilityStatus AvailabilityStatus `gorm:"not null"`
	FirstName          string             `gorm:"not null"`
	LastName           string             `gorm:"not null"`
	Phone              string             `gorm:"not null; unique"`
	Photo              string
	Passport           string
	GPSCoordinates     string
	Chats              []Chat  `gorm:"foreignKey:ParticipantID"`
	Orders             []Order `gorm:"foreignKey:CourierID"`
}
