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
		// Генерируем случайный тип участника
		participantType := utils.RandomParticipantType()

		var participantID uint
		if participantType == models.ClientParticipant && len(clients) > 0 {
			// Выбираем случайного клиента
			client := clients[rand.Intn(len(clients))]
			participantID = client.ID
		} else if participantType == models.CourierParticipant && len(couriers) > 0 {
			// Выбираем случайного курьера
			courier := couriers[rand.Intn(len(couriers))]
			participantID = courier.ID
		} else {
			// Если подходящего участника нет, пропускаем итерацию
			continue
		}

		// Выбираем случайного сотрудника поддержки
		support := supportStaff[rand.Intn(len(supportStaff))]

		// Создаем объект чата
		chat := models.Chat{
			ParticipantID:   participantID,
			ParticipantType: participantType, // Используем сгенерированный тип участника
			SupportStaffID:  &support.ID,
			Status:          utils.RandomStatus(),
			CreationDate:    utils.GenerateRandomDate2024(),
			Reason:          utils.GenerateReason(participantType),
		}

		// Сохраняем чат в базе данных
		if err := tx.Create(&chat).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		// Добавляем чат в список
		chats = append(chats, chat)
	}

	return chats, nil
}
