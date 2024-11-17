package utils

import (
	"ORM_DB/internal/models"
	"math/rand"
)

func RandomRole() models.Role {
	roles := []models.Role{models.Sender, models.Receiver}
	return roles[rand.Intn(len(roles))]
}

func RandomEmploymentStatus() models.EmploymentStatus {
	statuses := []models.EmploymentStatus{models.SelfEmployed, models.Official}
	return statuses[rand.Intn(len(statuses))]
}

func RandomTransportType() models.TransportType {
	types := []models.TransportType{models.OnFoot, models.Car, models.Truck}
	return types[rand.Intn(len(types))]
}

func RandomAvailabilityStatus() models.AvailabilityStatus {
	statuses := []models.AvailabilityStatus{models.Available, models.Busy}
	return statuses[rand.Intn(len(statuses))]
}

func RandomParticipantType() models.ParticipantType {
	types := []models.ParticipantType{models.ClientParticipant, models.CourierParticipant}
	return types[rand.Intn(len(types))]
}

func RandomStatus() models.Status {
	statuses := []models.Status{models.Open, models.Closed}
	return statuses[rand.Intn(len(statuses))]
}

func RandomSenderType() models.SenderType {
	types := []models.SenderType{models.ClientSender, models.CourierSender, models.SupportStaffSender}
	return types[rand.Intn(len(types))]
}

func RandomUrgency() models.Urgency {
	urgencies := []models.Urgency{models.Urgent, models.Scheduled}
	return urgencies[rand.Intn(len(urgencies))]
}

func RandomCurrentStatus() models.CurrentStatus {
	statuses := []models.CurrentStatus{models.WaitingForCourier, models.InTransit, models.Delivered}
	return statuses[rand.Intn(len(statuses))]
}

func RandomPaymentStatus() models.PaymentStatus {
	statuses := []models.PaymentStatus{models.Paid, models.Unpaid}
	return statuses[rand.Intn(len(statuses))]
}

func RandomPromoCodeType() models.PromoCodeType {
	types := []models.PromoCodeType{models.Discount, models.AdditionalService}
	return types[rand.Intn(len(types))]
}

func RandomDeliveryType() models.DeliveryType {
	types := []models.DeliveryType{models.DeliveryUrgent, models.DeliveryScheduled}
	return types[rand.Intn(len(types))]
}

func RandomJobTitle() string {
	jobTitles := []string{
		"Оператор", "Координатор", "Супервайзер", "Модератор", "Аналитик",
		"Администратор", "Менеджер", "Диспетчер", "Специалист",
	}
	return jobTitles[rand.Intn(len(jobTitles))]
}

func RandomPaymentMethod() string {
	methods := []string{"карта", "наличные"}
	return methods[rand.Intn(len(methods))]
}
