package seed

import (
	"ORM_DB/internal/database/seed/random_fields"
	"ORM_DB/models"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
	"log"
	"math"
	"math/rand"
	"time"
)

func generateSupportStaff(tx *gorm.DB) []models.SupportStaff {
	var supportStaffs []models.SupportStaff

	for i := 0; i < 150; i++ {
		firstName, lastName, _ := random_fields.GenerateFullRussianName()
		staff := models.SupportStaff{
			FirstName: firstName,
			LastName:  lastName,
			Phone:     faker.Phonenumber(),
			Email:     faker.Email(),
			Position:  random_fields.RandomJobTitle(),
		}
		if err := tx.Create(&staff).Error; err != nil {
			tx.Rollback()
			log.Fatalf("Ошибка при создании support staff: %v", err)
		}
		supportStaffs = append(supportStaffs, staff)
	}
	fmt.Println("Сгенерированы support staff")
	return supportStaffs
}

func generateClients(tx *gorm.DB) []models.Client {
	var clients []models.Client
	usedPhones := make(map[string]bool)

	for i := 0; i < 5000; i++ {
		var phone string
		for {
			phone = faker.Phonenumber()
			if !usedPhones[phone] {
				usedPhones[phone] = true
				break
			}
		}

		firstName, lastName, middleName := random_fields.GenerateFullRussianName()

		client := models.Client{
			Role:       random_fields.RandomRole(),
			FirstName:  firstName,
			LastName:   lastName,
			MiddleName: middleName,
			Phone:      phone,
			Email:      faker.Email(),
			Address:    random_fields.GenerateMoscowAddress(),
		}
		if err := tx.Create(&client).Error; err != nil {
			tx.Rollback()
			log.Fatalf("Ошибка при создании клиента: %v", err)
		}
		clients = append(clients, client)
	}
	fmt.Println("Сгенерированы клиенты")
	return clients
}

func generateCouriers(tx *gorm.DB) []models.Courier {
	var couriers []models.Courier
	usedPhones := make(map[string]bool)

	for i := 0; i < 2000; i++ {
		var phone string
		for {
			phone = faker.Phonenumber()
			if !usedPhones[phone] {
				usedPhones[phone] = true
				break
			}
		}

		firstName, lastName, _ := random_fields.GenerateFullRussianName()

		courier := models.Courier{
			EmploymentStatus:   random_fields.RandomEmploymentStatus(),
			TransportType:      random_fields.RandomTransportType(),
			AvailabilityStatus: random_fields.RandomAvailabilityStatus(),
			FirstName:          firstName,
			LastName:           lastName,
			Phone:              phone,
			Photo:              faker.URL(),
			Passport:           random_fields.GeneratePassportNumber(),
			GPSCoordinates:     random_fields.GenerateMoscowCoordinates(),
		}
		if err := tx.Create(&courier).Error; err != nil {
			tx.Rollback()
			log.Fatalf("Ошибка при создании курьера: %v", err)
		}
		couriers = append(couriers, courier)
	}
	fmt.Println("Сгенерированы курьеры")
	return couriers
}

func generatePromoCodes(tx *gorm.DB, clients []models.Client) []models.PromoCode {
	var promoCodes []models.PromoCode
	usedPromoCodes := make(map[string]bool) // Хранилище для проверки уникальности кодов

	for _, client := range clients {
		// Проверяем, сколько промокодов уже выдано пользователю
		var existingPromoCodesCount int64
		if err := tx.Model(&models.PromoCode{}).Where("client_id = ?", client.ID).Count(&existingPromoCodesCount).Error; err != nil {
			tx.Rollback()
			log.Fatalf("Ошибка при проверке существующих промокодов: %v", err)
		}

		// Генерируем промокоды только если у пользователя их меньше 5
		if existingPromoCodesCount >= 5 {
			continue
		}

		// Генерация до 5 промокодов на клиента
		for i := int(existingPromoCodesCount); i < 5; i++ {
			var promoCode string
			for {
				promoCode = fmt.Sprintf("%s-%d", faker.Word(), rand.Intn(100000))
				if !usedPromoCodes[promoCode] {
					usedPromoCodes[promoCode] = true // Уникальный код
					break
				}
			}

			promo := models.PromoCode{
				Type:           random_fields.RandomPromoCodeType(),
				ClientID:       &client.ID,
				Code:           promoCode,
				DiscountAmount: math.Round(rand.Float64() * 20),                     // Случайная скидка
				ValidUntil:     time.Now().AddDate(0, rand.Intn(12), rand.Intn(30)), // Случайная дата
				Personalized:   rand.Intn(2) == 0,
			}

			// Сохраняем в БД
			if err := tx.Create(&promo).Error; err != nil {
				tx.Rollback()
				log.Fatalf("Ошибка при создании промокода: %v", err)
			}
			promoCodes = append(promoCodes, promo)
		}
	}

	fmt.Println("Сгенерированы промокоды")
	return promoCodes
}

func generateRates(tx *gorm.DB) []models.Rate {
	var rates []models.Rate
	for i := 0; i < 50; i++ {
		rate := models.Rate{
			DeliveryType:  random_fields.RandomDeliveryType(),
			TransportType: random_fields.RandomTransportType(),
			Name:          fmt.Sprintf("Rate %d", i+1),
			Price:         math.Round(rand.Float64() * 100),
			Description:   faker.Sentence(),
		}
		if err := tx.Create(&rate).Error; err != nil {
			tx.Rollback()
			log.Fatalf("Ошибка при создании тарифа: %v", err)
		}
		rates = append(rates, rate)
	}
	fmt.Println("Сгенерированы тарифы")
	return rates
}

func generateOrders(tx *gorm.DB, clients []models.Client, couriers []models.Courier, promoCodes []models.PromoCode, rates []models.Rate) {
	// Создаем мапу промокодов для быстрого поиска
	clientPromoCodes := make(map[uint][]models.PromoCode)
	for _, promo := range promoCodes {
		if promo.ClientID != nil {
			clientPromoCodes[*promo.ClientID] = append(clientPromoCodes[*promo.ClientID], promo)
		}
	}

	for i := 0; i < 50000; i++ {
		sender := clients[rand.Intn(len(clients))]

		var recipient models.Client
		for {
			recipient = clients[rand.Intn(len(clients))]
			if recipient.ID != sender.ID {
				break
			}
		}

		var promoCodeID *uint
		if promos, exists := clientPromoCodes[sender.ID]; exists && len(promos) > 0 {
			randomPromo := promos[rand.Intn(len(promos))]
			promoCodeID = &randomPromo.ID
		} else if promos, exists := clientPromoCodes[recipient.ID]; exists && len(promos) > 0 {
			randomPromo := promos[rand.Intn(len(promos))]
			promoCodeID = &randomPromo.ID
		}

		// Генерация заказа
		order := models.Order{
			Urgency:            random_fields.RandomUrgency(),
			SenderID:           sender.ID,
			RecipientID:        recipient.ID,
			CourierID:          couriers[rand.Intn(len(couriers))].ID,
			CurrentStatus:      random_fields.RandomCurrentStatus(),
			PaymentStatus:      random_fields.RandomPaymentStatus(),
			RateID:             rates[rand.Intn(len(rates))].ID,
			PromoCodeID:        promoCodeID, // Только если промокод связан с клиентом
			CreationDate:       random_fields.GenerateRandomDate2024(),
			ItemType:           random_fields.RandomItemType(),
			ItemValue:          math.Round(rand.Float64() * 100),
			Weight:             math.Round(rand.Float64() * 10),
			DiscountSurcharges: math.Round(rand.Float64() * 10),
			PaymentMethod:      random_fields.RandomPaymentMethod(),
		}

		// Сохранение заказа в БД
		if err := tx.Create(&order).Error; err != nil {
			tx.Rollback()
			log.Fatalf("Ошибка при создании заказа: %v", err)
		}
	}
	fmt.Println("Сгенерированы заказы")
}

func generateChats(tx *gorm.DB, supportStaffs []models.SupportStaff, clients []models.Client, couriers []models.Courier) []models.Chat {
	var chats []models.Chat

	for i := 0; i < 1000; i++ {
		// Определяем случайный тип участника (client или courier)
		participantType := random_fields.RandomParticipantType()
		var participantID uint

		// Выбираем ID участника в зависимости от его типа
		if participantType == models.ClientParticipant && len(clients) > 0 {
			participantID = clients[rand.Intn(len(clients))].ID
		} else if participantType == models.CourierParticipant && len(couriers) > 0 {
			participantID = couriers[rand.Intn(len(couriers))].ID
		} else {
			continue // Если подходящий участник не найден, пропускаем итерацию
		}

		// Выбираем случайного сотрудника поддержки
		supportStaffID := supportStaffs[rand.Intn(len(supportStaffs))].ID

		// Генерируем причину чата в зависимости от типа участника
		reason := random_fields.GenerateReason(participantType)

		// Создаём чат
		chat := models.Chat{
			ParticipantType: participantType,
			ParticipantID:   participantID,
			SupportStaffID:  &supportStaffID,
			Status:          random_fields.RandomStatus(),
			CreationDate:    random_fields.GenerateRandomDate2024(),
			Reason:          reason,
		}

		// Сохраняем чат в БД
		if err := tx.Create(&chat).Error; err != nil {
			tx.Rollback()
			log.Fatalf("Ошибка при создании чата: %v", err)
		}

		chats = append(chats, chat)
	}

	fmt.Println("Сгенерированы чаты")
	return chats
}

func generateMessages(tx *gorm.DB, chats []models.Chat, clients []models.Client, couriers []models.Courier, supportStaffs []models.SupportStaff) {
	for i := 0; i < 8000; i++ {
		// Выбираем случайный чат
		chat := chats[rand.Intn(len(chats))]

		// Определяем случайный тип отправителя
		senderType := random_fields.RandomSenderType()

		var senderID uint
		var recipientType models.SenderTypeEnum

		switch senderType {
		case models.ClientSender:
			// Отправитель - клиент
			sender := clients[rand.Intn(len(clients))]
			senderID = sender.ID

			// Получатель - либо курьер, либо поддержка
			if rand.Intn(2) == 0 {
				recipientType = models.CourierSender
			} else {
				recipientType = models.SupportStaffSender
			}

		case models.CourierSender:
			// Отправитель - курьер
			sender := couriers[rand.Intn(len(couriers))]
			senderID = sender.ID

			// Получатель - либо клиент, либо поддержка
			if rand.Intn(2) == 0 {
				recipientType = models.ClientSender
			} else {
				recipientType = models.SupportStaffSender
			}

		case models.SupportStaffSender:
			// Отправитель - поддержка
			sender := supportStaffs[rand.Intn(len(supportStaffs))]
			senderID = sender.ID

			// Получатель - либо клиент, либо курьер
			if rand.Intn(2) == 0 {
				recipientType = models.ClientSender
			} else {
				recipientType = models.CourierSender
			}
		}

		// Генерируем сообщение на основе отправителя и получателя
		message := models.Message{
			ChatID:     chat.ID,
			SenderType: senderType,
			SenderID:   senderID,
			Timestamp:  random_fields.GenerateRandomDate2024(),
			Text:       random_fields.GenerateMessage(senderType, recipientType), // Передаем оба параметра
		}

		// Сохраняем сообщение в БД
		if err := tx.Create(&message).Error; err != nil {
			tx.Rollback()
			log.Fatalf("Ошибка при создании сообщения: %v", err)
		}
	}

	fmt.Println("Сгенерированы сообщения")
}
