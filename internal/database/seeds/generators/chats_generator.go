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

	// Проверяем, есть ли данные для генерации чатов
	if len(clients) == 0 || len(couriers) == 0 || len(supportStaff) == 0 {
		return nil, fmt.Errorf("нет данных для генерации чатов")
	}

	for i := 0; i < count; i++ {
		// Случайно выбираем тип участника
		participantType := utils.RandomParticipantType()

		var clientID, courierID *uint
		if participantType == models.ClientParticipant && len(clients) > 0 {
			// Если участник - клиент, выбираем случайного клиента
			client := clients[rand.Intn(len(clients))]
			clientID = &client.ID
		} else if participantType == models.CourierParticipant && len(couriers) > 0 {
			// Если участник - курьер, выбираем случайного курьера
			courier := couriers[rand.Intn(len(couriers))]
			courierID = &courier.ID
		} else {
			// Пропускаем, если не удалось выбрать участника
			continue
		}

		// Выбираем случайного сотрудника поддержки
		support := supportStaff[rand.Intn(len(supportStaff))]

		// Создаем чат с заполнением соответствующих полей
		chat := models.Chat{
			ParticipantType: participantType,
			ClientID:        clientID,
			CourierID:       courierID,
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
