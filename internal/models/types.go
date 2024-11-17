package models

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
