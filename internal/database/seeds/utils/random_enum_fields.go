package utils

import (
	models2 "ORM_DB/internal/models"
	"ORM_DB/models"
	"math/rand"
)

func RandomRole() models.RoleEnum {
	roles := []models.RoleEnum{models2.Sender, models2.Receiver}
	return roles[rand.Intn(len(roles))]
}

func RandomEmploymentStatus() models.EmploymentStatusEnum {
	statuses := []models.EmploymentStatusEnum{models2.SelfEmployed, models2.Official}
	return statuses[rand.Intn(len(statuses))]
}

func RandomTransportType() models.TransportTypeEnum {
	types := []models.TransportTypeEnum{models2.OnFoot, models2.Car, models2.Truck}
	return types[rand.Intn(len(types))]
}

func RandomAvailabilityStatus() models.AvailabilityStatusEnum {
	statuses := []models.AvailabilityStatusEnum{models2.Available, models2.Busy}
	return statuses[rand.Intn(len(statuses))]
}

func RandomParticipantType() models.ParticipantTypeEnum {
	types := []models.ParticipantTypeEnum{models2.ClientParticipant, models2.CourierParticipant}
	return types[rand.Intn(len(types))]
}

func RandomStatus() models.StatusEnum {
	statuses := []models.StatusEnum{models2.Open, models2.Closed}
	return statuses[rand.Intn(len(statuses))]
}

func RandomSenderType() models.SenderTypeEnum {
	types := []models.SenderTypeEnum{
		models2.ClientSender,
		models2.CourierSender,
		models2.SupportStaffSender,
	}
	return types[rand.Intn(len(types))]
}

func RandomUrgency() models.UrgencyEnum {
	urgencies := []models.UrgencyEnum{models2.Urgent, models2.Scheduled}
	return urgencies[rand.Intn(len(urgencies))]
}

func RandomCurrentStatus() models.CurrentStatusEnum {
	statuses := []models.CurrentStatusEnum{models2.WaitingForCourier, models2.InTransit, models2.Delivered}
	return statuses[rand.Intn(len(statuses))]
}

func RandomPaymentStatus() models.PaymentStatusEnum {
	statuses := []models.PaymentStatusEnum{models2.Paid, models2.Unpaid}
	return statuses[rand.Intn(len(statuses))]
}

func RandomPromoCodeType() models.PromoCodeTypeEnum {
	types := []models.PromoCodeTypeEnum{models2.Discount, models2.AdditionalService}
	return types[rand.Intn(len(types))]
}

func RandomDeliveryType() models.DeliveryTypeEnum {
	types := []models.DeliveryTypeEnum{models2.DeliveryUrgent, models2.DeliveryScheduled}
	return types[rand.Intn(len(types))]
}

func RandomJobTitle() string {
	jobTitles := []string{"Manager", "Engineer", "Developer", "Consultant", "Analyst", "Specialist", "Coordinator"}
	return jobTitles[rand.Intn(len(jobTitles))]
}

func RandomPaymentMethod() string {
	methods := []string{"credit card", "cash"}
	return methods[rand.Intn(len(methods))]
}
