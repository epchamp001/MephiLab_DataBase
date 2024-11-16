package random_fields

import (
	randData "ORM_DB/internal/database/seed/RandomData"
	"ORM_DB/models"
	"math/rand"
)

func GenerateMessage(senderType, recipientType models.SenderTypeEnum) string {
	switch senderType {
	case models.ClientSender:
		if recipientType == models.SupportStaffSender {
			return randData.SupportMessagesClient[rand.Intn(len(randData.SupportMessagesClient))]
		} else if recipientType == models.CourierSender {
			return randData.ClientMessages[rand.Intn(len(randData.ClientMessages))]
		}
	case models.CourierSender:
		if recipientType == models.SupportStaffSender {
			return randData.SupportMessagesCourier[rand.Intn(len(randData.SupportMessagesCourier))]
		} else if recipientType == models.ClientSender {
			return randData.CourierMessages[rand.Intn(len(randData.CourierMessages))]
		}
	case models.SupportStaffSender:
		if recipientType == models.ClientSender {
			return randData.SupportMessagesClient[rand.Intn(len(randData.SupportMessagesClient))]
		} else if recipientType == models.CourierSender {
			return randData.SupportMessagesCourier[rand.Intn(len(randData.SupportMessagesCourier))]
		}
	default:
		return "Неопределённый отправитель."
	}
	return "Некорректное сочетание отправителя и получателя."
}
