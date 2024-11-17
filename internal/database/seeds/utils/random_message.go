package utils

import (
	randData "ORM_DB/internal/database/seeds/parsers"
	models2 "ORM_DB/internal/models"
	"ORM_DB/models"
	"math/rand"
)

func GenerateMessage(senderType, recipientType models.SenderTypeEnum) string {
	switch senderType {
	case models2.ClientSender:
		if recipientType == models2.SupportStaffSender {
			return randData.SupportMessagesClient[rand.Intn(len(randData.SupportMessagesClient))]
		} else if recipientType == models2.CourierSender {
			return randData.ClientMessages[rand.Intn(len(randData.ClientMessages))]
		}
	case models2.CourierSender:
		if recipientType == models2.SupportStaffSender {
			return randData.SupportMessagesCourier[rand.Intn(len(randData.SupportMessagesCourier))]
		} else if recipientType == models2.ClientSender {
			return randData.CourierMessages[rand.Intn(len(randData.CourierMessages))]
		}
	case models2.SupportStaffSender:
		if recipientType == models2.ClientSender {
			return randData.SupportMessagesClient[rand.Intn(len(randData.SupportMessagesClient))]
		} else if recipientType == models2.CourierSender {
			return randData.SupportMessagesCourier[rand.Intn(len(randData.SupportMessagesCourier))]
		}
	default:
		return "Неопределённый отправитель."
	}
	return "Некорректное сочетание отправителя и получателя."
}
