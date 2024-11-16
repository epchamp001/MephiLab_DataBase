package models

import (
	"time"
)

// Типы данных для строковых констант
type Role string
type EmploymentStatus string
type TransportType string
type AvailabilityStatus string
type ParticipantType string
type Status string
type SenderType string
type Urgency string
type CurrentStatus string
type PaymentStatus string
type PromoCodeType string
type DeliveryType string

// Константы
const (
	// Role
	Sender   Role = "sender"
	Receiver Role = "receiver"

	// EmploymentStatus
	SelfEmployed EmploymentStatus = "self-employed"
	Official     EmploymentStatus = "official"

	// TransportType
	OnFoot TransportType = "on foot"
	Car    TransportType = "car"
	Truck  TransportType = "truck"

	// AvailabilityStatus
	Available AvailabilityStatus = "available"
	Busy      AvailabilityStatus = "busy"

	// ParticipantType
	ClientParticipant  ParticipantType = "client"
	CourierParticipant ParticipantType = "courier"

	// Status
	Open   Status = "open"
	Closed Status = "closed"

	// SenderType
	ClientSender       SenderType = "client"
	CourierSender      SenderType = "courier"
	SupportStaffSender SenderType = "support staff"

	// Urgency
	Urgent    Urgency = "urgent"
	Scheduled Urgency = "scheduled"

	// CurrentStatus
	WaitingForCourier CurrentStatus = "waiting for courier"
	InTransit         CurrentStatus = "in transit"
	Delivered         CurrentStatus = "delivered"

	// PaymentStatus
	Paid   PaymentStatus = "paid"
	Unpaid PaymentStatus = "unpaid"

	// PromoCodeType
	Discount          PromoCodeType = "discount"
	AdditionalService PromoCodeType = "additional service"

	// DeliveryType
	DeliveryUrgent    DeliveryType = "urgent"
	DeliveryScheduled DeliveryType = "scheduled"
)

// Entity structs with relationships

type SupportStaff struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Phone     string `gorm:"not null"`
	Email     string `gorm:"not null; unique"`
	Position  string
	Chats     []Chat `gorm:"foreignKey:SupportStaffID"` // One-to-Many связь с Chat
}

type Client struct {
	ID         uint   `gorm:"primaryKey"`
	Role       Role   `gorm:"not null"` // Role: sender, receiver
	FirstName  string `gorm:"not null"`
	LastName   string `gorm:"not null"`
	MiddleName string
	Phone      string `gorm:"not null; unique"`
	Email      string `gorm:"not null; unique"`
	Address    string
	Chats      []Chat      `gorm:"foreignKey:ParticipantID"`
	OrdersSent []Order     `gorm:"foreignKey:SenderID"`
	OrdersRecv []Order     `gorm:"foreignKey:RecipientID"`
	PromoCodes []PromoCode `gorm:"foreignKey:ClientID"`
}

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

type Message struct {
	ID         uint       `gorm:"primaryKey"`
	ChatID     uint       `gorm:"not null;constraint:OnDelete:CASCADE"`
	SenderType SenderType `gorm:"not null"`
	SenderID   uint       `gorm:"not null"`
	Timestamp  time.Time
	Text       string `gorm:"not null"`
	Chat       Chat   `gorm:"foreignKey:ChatID"`
}

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

type Rate struct {
	ID            uint          `gorm:"primaryKey"`
	DeliveryType  DeliveryType  `gorm:"not null"`
	TransportType TransportType `gorm:"not null"`
	Name          string        `gorm:"not null"`
	Price         float64       `gorm:"not null"`
	Description   string
}
