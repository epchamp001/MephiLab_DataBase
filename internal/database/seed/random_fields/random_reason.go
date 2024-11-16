package random_fields

import (
	randData "ORM_DB/internal/database/seed/RandomData"
	"ORM_DB/models"
	"math/rand"
)

func GenerateReason(participantType models.ParticipantTypeEnum) string {
	switch participantType {
	case models.CourierParticipant:
		// Если участник - курьер, выбираем причину из списка курьеров
		return randData.ReasonChatCourier[rand.Intn(len(randData.ReasonChatCourier))]
	case models.ClientParticipant:
		// Если участник - клиент, выбираем причину из списка клиентов
		return randData.ReasonChatClient[rand.Intn(len(randData.ReasonChatClient))]
	default:
		return "Неопределённый участник."
	}
}
