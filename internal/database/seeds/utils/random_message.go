package utils

import (
	"ORM_DB/internal/database/seeds/parsers"
	"ORM_DB/internal/models"
	"math/rand"
)

func GenerateMessage(senderType, recipientType models.SenderType) string {
	switch senderType {
	case models.ClientSender:
		if recipientType == models.SupportStaffSender {
			return parsers.GlobalData.SupportMessagesClient[rand.Intn(len(parsers.GlobalData.SupportMessagesClient))]
		} else if recipientType == models.CourierSender {
			return parsers.GlobalData.ClientMessages[rand.Intn(len(parsers.GlobalData.ClientMessages))]
		}
	case models.CourierSender:
		if recipientType == models.SupportStaffSender {
			return parsers.GlobalData.SupportMessagesCourier[rand.Intn(len(parsers.GlobalData.SupportMessagesCourier))]
		} else if recipientType == models.ClientSender {
			return parsers.GlobalData.CourierMessages[rand.Intn(len(parsers.GlobalData.CourierMessages))]
		}
	case models.SupportStaffSender:
		if recipientType == models.ClientSender {
			return parsers.GlobalData.SupportMessagesClient[rand.Intn(len(parsers.GlobalData.SupportMessagesClient))]
		} else if recipientType == models.CourierSender {
			return parsers.GlobalData.SupportMessagesCourier[rand.Intn(len(parsers.GlobalData.SupportMessagesCourier))]
		}
	default:
		return "Неопределённый отправитель."
	}
	return "Некорректное сочетание отправителя и получателя."
}
