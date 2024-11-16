package seed

import (
	"ORM_BD/models"
	"math/rand"
)

func randomRole() models.RoleEnum {
	roles := []models.RoleEnum{models.Sender, models.Receiver}
	return roles[rand.Intn(len(roles))]
}

func randomEmploymentStatus() models.EmploymentStatusEnum {
	statuses := []models.EmploymentStatusEnum{models.SelfEmployed, models.Official}
	return statuses[rand.Intn(len(statuses))]
}

func randomTransportType() models.TransportTypeEnum {
	types := []models.TransportTypeEnum{models.OnFoot, models.Car, models.Truck}
	return types[rand.Intn(len(types))]
}

func randomAvailabilityStatus() models.AvailabilityStatusEnum {
	statuses := []models.AvailabilityStatusEnum{models.Available, models.Busy}
	return statuses[rand.Intn(len(statuses))]
}

func randomParticipantType() models.ParticipantTypeEnum {
	types := []models.ParticipantTypeEnum{models.ClientParticipant, models.CourierParticipant}
	return types[rand.Intn(len(types))]
}

func randomStatus() models.StatusEnum {
	statuses := []models.StatusEnum{models.Open, models.Closed}
	return statuses[rand.Intn(len(statuses))]
}

func randomSenderType() models.SenderTypeEnum {
	types := []models.SenderTypeEnum{models.ClientSender, models.CourierSender, models.SupportStaffSender}
	return types[rand.Intn(len(types))]
}

func randomUrgency() models.UrgencyEnum {
	urgencies := []models.UrgencyEnum{models.Urgent, models.Scheduled}
	return urgencies[rand.Intn(len(urgencies))]
}

func randomCurrentStatus() models.CurrentStatusEnum {
	statuses := []models.CurrentStatusEnum{models.WaitingForCourier, models.InTransit, models.Delivered}
	return statuses[rand.Intn(len(statuses))]
}

func randomPaymentStatus() models.PaymentStatusEnum {
	statuses := []models.PaymentStatusEnum{models.Paid, models.Unpaid}
	return statuses[rand.Intn(len(statuses))]
}

func randomPromoCodeType() models.PromoCodeTypeEnum {
	types := []models.PromoCodeTypeEnum{models.Discount, models.AdditionalService}
	return types[rand.Intn(len(types))]
}

func randomDeliveryType() models.DeliveryTypeEnum {
	types := []models.DeliveryTypeEnum{models.DeliveryUrgent, models.DeliveryScheduled}
	return types[rand.Intn(len(types))]
}

func randomJobTitle() string {
	jobTitles := []string{"Manager", "Engineer", "Developer", "Consultant", "Analyst", "Specialist", "Coordinator"}
	return jobTitles[rand.Intn(len(jobTitles))]
}

func randomPaymentMethod() string {
	methods := []string{"credit card", "cash"}
	return methods[rand.Intn(len(methods))]
}
