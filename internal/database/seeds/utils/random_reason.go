package utils

import (
	randData "ORM_DB/internal/database/seeds/parsers"
	models2 "ORM_DB/internal/models"
	"ORM_DB/models"
	"math/rand"
)

func GenerateReason(participantType models.ParticipantTypeEnum) string {
	switch participantType {
	case models2.CourierParticipant:
		// Если участник - курьер, выбираем причину из списка курьеров
		return randData.ReasonChatCourier[rand.Intn(len(randData.ReasonChatCourier))]
	case models2.ClientParticipant:
		// Если участник - клиент, выбираем причину из списка клиентов
		return randData.ReasonChatClient[rand.Intn(len(randData.ReasonChatClient))]
	default:
		return "Неопределённый участник."
	}
}
