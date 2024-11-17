package generators

import (
	"ORM_DB/internal/database/seeds/utils"
	"ORM_DB/internal/models"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
)

func GenerateChats(tx *gorm.DB, clients []models.Client, couriers []models.Courier, supportStaff []models.SupportStaff, count int) ([]models.Chat, error) {
	var chats []models.Chat

	if len(clients) == 0 || len(couriers) == 0 || len(supportStaff) == 0 {
		return nil, fmt.Errorf("нет данных для генерации чатов")
	}

	for i := 0; i < count; i++ {
		participantType := utils.RandomParticipantType()

		var participantID uint
		if participantType == models.ClientParticipant && len(clients) > 0 {
			client := clients[rand.Intn(len(clients))]
			participantID = client.ID
		} else if participantType == models.CourierParticipant && len(couriers) > 0 {
			courier := couriers[rand.Intn(len(couriers))]
			participantID = courier.ID
		} else {
			continue
		}

		support := supportStaff[rand.Intn(len(supportStaff))]

		chat := models.Chat{
			ParticipantID:   participantID,
			ParticipantType: participantType, // Используем сгенерированный тип участника
			SupportStaffID:  &support.ID,
			Status:          utils.RandomStatus(),
			CreationDate:    utils.GenerateRandomDate2024(),
			Reason:          utils.GenerateReason(participantType),
		}

		if err := tx.Create(&chat).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		chats = append(chats, chat)
	}

	return chats, nil
}
