package random_fields

import (
	"ORM_DB/models"
	"math/rand"
)

func RandomRole() models.RoleEnum {
	roles := []models.RoleEnum{models.Sender, models.Receiver}
	return roles[rand.Intn(len(roles))]
}

func RandomEmploymentStatus() models.EmploymentStatusEnum {
	statuses := []models.EmploymentStatusEnum{models.SelfEmployed, models.Official}
	return statuses[rand.Intn(len(statuses))]
}

func RandomTransportType() models.TransportTypeEnum {
	types := []models.TransportTypeEnum{models.OnFoot, models.Car, models.Truck}
	return types[rand.Intn(len(types))]
}

func RandomAvailabilityStatus() models.AvailabilityStatusEnum {
	statuses := []models.AvailabilityStatusEnum{models.Available, models.Busy}
	return statuses[rand.Intn(len(statuses))]
}

func RandomParticipantType() models.ParticipantTypeEnum {
	types := []models.ParticipantTypeEnum{models.ClientParticipant, models.CourierParticipant}
	return types[rand.Intn(len(types))]
}

func RandomStatus() models.StatusEnum {
	statuses := []models.StatusEnum{models.Open, models.Closed}
	return statuses[rand.Intn(len(statuses))]
}

func RandomSenderType() models.SenderTypeEnum {
	types := []models.SenderTypeEnum{
		models.ClientSender,
		models.CourierSender,
		models.SupportStaffSender,
	}
	return types[rand.Intn(len(types))]
}

func RandomUrgency() models.UrgencyEnum {
	urgencies := []models.UrgencyEnum{models.Urgent, models.Scheduled}
	return urgencies[rand.Intn(len(urgencies))]
}

func RandomCurrentStatus() models.CurrentStatusEnum {
	statuses := []models.CurrentStatusEnum{models.WaitingForCourier, models.InTransit, models.Delivered}
	return statuses[rand.Intn(len(statuses))]
}

func RandomPaymentStatus() models.PaymentStatusEnum {
	statuses := []models.PaymentStatusEnum{models.Paid, models.Unpaid}
	return statuses[rand.Intn(len(statuses))]
}

func RandomPromoCodeType() models.PromoCodeTypeEnum {
	types := []models.PromoCodeTypeEnum{models.Discount, models.AdditionalService}
	return types[rand.Intn(len(types))]
}

func RandomDeliveryType() models.DeliveryTypeEnum {
	types := []models.DeliveryTypeEnum{models.DeliveryUrgent, models.DeliveryScheduled}
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
