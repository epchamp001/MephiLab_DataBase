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

const (
	// Role
	Sender   Role = "отправитель"
	Receiver Role = "получатель"

	// EmploymentStatus
	SelfEmployed EmploymentStatus = "самозанятый"
	Official     EmploymentStatus = "официально трудоустроенный"

	// TransportType
	OnFoot TransportType = "пешком"
	Car    TransportType = "автомобиль"
	Truck  TransportType = "грузовик"

	// AvailabilityStatus
	Available AvailabilityStatus = "доступен"
	Busy      AvailabilityStatus = "занят"

	// ParticipantType
	ClientParticipant  ParticipantType = "клиент"
	CourierParticipant ParticipantType = "курьер"

	// Status
	Open   Status = "открыт"
	Closed Status = "закрыт"

	// SenderType
	ClientSender       SenderType = "клиент"
	CourierSender      SenderType = "курьер"
	SupportStaffSender SenderType = "сотрудник поддержки"

	// Urgency
	Urgent    Urgency = "срочно"
	Scheduled Urgency = "запланировано"

	// CurrentStatus
	WaitingForCourier CurrentStatus = "ожидание курьера"
	InTransit         CurrentStatus = "в пути"
	Delivered         CurrentStatus = "доставлено"

	// PaymentStatus
	Paid   PaymentStatus = "оплачено"
	Unpaid PaymentStatus = "не оплачено"

	// PromoCodeType
	Discount          PromoCodeType = "скидка"
	AdditionalService PromoCodeType = "дополнительная услуга"

	// DeliveryType
	DeliveryUrgent    DeliveryType = "срочная доставка"
	DeliveryScheduled DeliveryType = "запланированная доставка"
)
