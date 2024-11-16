package models

import (
	"time"
)

// Enum types
type RoleEnum string

const (
	Sender   RoleEnum = "sender"
	Receiver RoleEnum = "receiver"
)

type EmploymentStatusEnum string

const (
	SelfEmployed EmploymentStatusEnum = "self-employed"
	Official     EmploymentStatusEnum = "official"
)

type TransportTypeEnum string

const (
	OnFoot TransportTypeEnum = "on foot"
	Car    TransportTypeEnum = "car"
	Truck  TransportTypeEnum = "truck"
)

type AvailabilityStatusEnum string

const (
	Available AvailabilityStatusEnum = "available"
	Busy      AvailabilityStatusEnum = "busy"
)

type ParticipantTypeEnum string

const (
	ClientParticipant  ParticipantTypeEnum = "client"
	CourierParticipant ParticipantTypeEnum = "courier"
)

type StatusEnum string

const (
	Open   StatusEnum = "open"
	Closed StatusEnum = "closed"
)

type SenderTypeEnum string

const (
	ClientSender       SenderTypeEnum = "client"
	CourierSender      SenderTypeEnum = "courier"
	SupportStaffSender SenderTypeEnum = "support staff"
)

type UrgencyEnum string

const (
	Urgent    UrgencyEnum = "urgent"
	Scheduled UrgencyEnum = "scheduled"
)

type CurrentStatusEnum string

const (
	WaitingForCourier CurrentStatusEnum = "waiting for courier"
	InTransit         CurrentStatusEnum = "in transit"
	Delivered         CurrentStatusEnum = "delivered"
)

type PaymentStatusEnum string

const (
	Paid   PaymentStatusEnum = "paid"
	Unpaid PaymentStatusEnum = "unpaid"
)

type PromoCodeTypeEnum string

const (
	Discount          PromoCodeTypeEnum = "discount"
	AdditionalService PromoCodeTypeEnum = "additional service"
)

type DeliveryTypeEnum string

const (
	DeliveryUrgent    DeliveryTypeEnum = "urgent"
	DeliveryScheduled DeliveryTypeEnum = "scheduled"
)

// Entity structs with relationships

type SupportStaff struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Phone     string `gorm:"not null"`
	Email     string `gorm:"not null; unique"`
	Position  string
}

type Client struct {
	ID         uint     `gorm:"primaryKey"`
	Role       RoleEnum `gorm:"type:role_enum; not null"`
	FirstName  string   `gorm:"not null"`
	LastName   string   `gorm:"not null"`
	MiddleName string
	Phone      string `gorm:"not null; unique"`
	Email      string `gorm:"not null; unique"`
	Address    string
}

type Courier struct {
	ID                 uint                   `gorm:"primaryKey"`
	EmploymentStatus   EmploymentStatusEnum   `gorm:"type:employment_status_enum; not null"`
	TransportType      TransportTypeEnum      `gorm:"type:transport_type_enum; not null"`
	AvailabilityStatus AvailabilityStatusEnum `gorm:"type:availability_status_enum; not null"`
	FirstName          string                 `gorm:"not null"`
	LastName           string                 `gorm:"not null"`
	Phone              string                 `gorm:"not null; unique"`
	Photo              string
	Passport           string
	GPSCoordinates     string
}

type Chat struct {
	ID              uint                `gorm:"primaryKey"`
	ParticipantType ParticipantTypeEnum `gorm:"type:participant_type_enum; not null"` // client или courier
	ParticipantID   uint                `gorm:"not null"`
	SupportStaffID  *uint               `gorm:"constraint:OnDelete:SET NULL"`
	Status          StatusEnum          `gorm:"type:status_enum; not null"`
	CreationDate    time.Time
	Reason          string
}

type Message struct {
	ID         uint           `gorm:"primaryKey"`
	ChatID     uint           `gorm:"not null;constraint:OnDelete:CASCADE"`
	SenderType SenderTypeEnum `gorm:"type:sender_type_enum; not null"`
	SenderID   uint           `gorm:"not null"`
	Timestamp  time.Time      `gorm:"not null"`
	Text       string         `gorm:"not null"`
}

type Order struct {
	ID                 uint              `gorm:"primaryKey"`
	Urgency            UrgencyEnum       `gorm:"type:urgency_enum; not null"`
	SenderID           uint              `gorm:"not null;constraint:OnDelete:CASCADE"`
	RecipientID        uint              `gorm:"not null;constraint:OnDelete:CASCADE"`
	CourierID          uint              `gorm:"not null;constraint:OnDelete:SET NULL"`
	CurrentStatus      CurrentStatusEnum `gorm:"type:current_status_enum; not null"`
	PromoCodeID        *uint             `gorm:"constraint:OnDelete:SET NULL"`
	PaymentStatus      PaymentStatusEnum `gorm:"type:payment_status_enum; not null"`
	RateID             uint              `gorm:"not null;constraint:OnDelete:SET NULL"`
	CreationDate       time.Time
	ItemType           string
	ItemValue          float64
	Weight             float64
	DiscountSurcharges float64
	PaymentMethod      string
}

type PromoCode struct {
	ID             uint              `gorm:"primaryKey"`
	Type           PromoCodeTypeEnum `gorm:"type:promo_code_type_enum; not null"`
	ClientID       *uint             `gorm:"constraint:OnDelete:SET NULL"`
	Code           string            `gorm:"not null; unique"`
	DiscountAmount float64           `gorm:"not null"`
	ValidUntil     time.Time
	Personalized   bool
}

type Rate struct {
	ID            uint              `gorm:"primaryKey"`
	DeliveryType  DeliveryTypeEnum  `gorm:"type:delivery_type_enum; not null"`
	TransportType TransportTypeEnum `gorm:"type:transport_type_enum; not null"`
	Name          string            `gorm:"not null"`
	Price         float64           `gorm:"not null"`
	Description   string
}
