package generators

import (
	"ORM_DB/internal/database/seeds/utils"
	"ORM_DB/internal/models"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
)

func RandomRecipientType(exclude models.SenderType) models.SenderType {
	types := []models.SenderType{models.ClientSender, models.CourierSender, models.SupportStaffSender}
	for {
		t := types[rand.Intn(len(types))]
		if t != exclude {
			return t
		}
	}
}

func GenerateMessages(tx *gorm.DB, chats []models.Chat, clients []models.Client, couriers []models.Courier, supportStaffs []models.SupportStaff, count int) ([]models.Message, error) {
	var messages []models.Message

	if len(chats) == 0 || (len(clients) == 0 && len(couriers) == 0 && len(supportStaffs) == 0) {
		return nil, fmt.Errorf("Нет данных для генерации сообщений")
	}

	for i := 0; i < count; i++ {
		chat := chats[rand.Intn(len(chats))]
		senderType := utils.RandomSenderType()

		var senderID uint
		var recipientType models.SenderType

		switch senderType {
		case models.ClientSender:
			if len(clients) > 0 {
				sender := clients[rand.Intn(len(clients))]
				senderID = sender.ID
				recipientType = RandomRecipientType(senderType)
			}

		case models.CourierSender:
			if len(couriers) > 0 {
				sender := couriers[rand.Intn(len(couriers))]
				senderID = sender.ID
				recipientType = RandomRecipientType(senderType)
			}

		case models.SupportStaffSender:
			if len(supportStaffs) > 0 {
				sender := supportStaffs[rand.Intn(len(supportStaffs))]
				senderID = sender.ID
				recipientType = RandomRecipientType(senderType)
			}
		}

		if senderID == 0 {
			continue
		}

		message := models.Message{
			ChatID:     chat.ID,
			SenderType: senderType,
			SenderID:   senderID,
			Timestamp:  utils.GenerateRandomDate2024(),
			Text:       utils.GenerateMessage(senderType, recipientType),
		}

		if err := tx.Create(&message).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		messages = append(messages, message)
	}

	return messages, nil
}
