package utils

import (
	"ORM_DB/internal/database/seeds/parsers"
	"ORM_DB/internal/models"
	"math/rand"
)

func GenerateReason(participantType models.ParticipantType) string {
	switch participantType {
	case models.CourierParticipant:
		// Если участник - курьер, выбираем причину из списка курьеров
		if len(parsers.GlobalData.ReasonChatCourier) > 0 {
			return parsers.GlobalData.ReasonChatCourier[rand.Intn(len(parsers.GlobalData.ReasonChatCourier))]
		}
		return "Данные для причин чатов с курьерами не загружены"
	case models.ClientParticipant:
		// Если участник - клиент, выбираем причину из списка клиентов
		if len(parsers.GlobalData.ReasonChatClient) > 0 {
			return parsers.GlobalData.ReasonChatClient[rand.Intn(len(parsers.GlobalData.ReasonChatClient))]
		}
		return "Данные для причин чатов с клиентами не загружены"
	default:
		return "Неопределённый участник."
	}
}
