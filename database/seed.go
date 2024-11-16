package database

import (
	"ORM_BD/models"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

func randomRole() models.RoleEnum {
	roles := []models.RoleEnum{models.Sender, models.Receiver}
	return roles[rand.Intn(len(roles))]
}

func randomEmploymentStatus() models.EmploymentStatusEnum {
	statuses := []models.EmploymentStatusEnum{models.SelfEmployed, models.Official}
	return statuses[rand.Intn(len(statuses))]
}

func randomTransportType() models.TransportTypeEnum {
	types := []models.TransportTypeEnum{models.OnFoot, models.Car, models.Truck}
	return types[rand.Intn(len(types))]
}

func randomAvailabilityStatus() models.AvailabilityStatusEnum {
	statuses := []models.AvailabilityStatusEnum{models.Available, models.Busy}
	return statuses[rand.Intn(len(statuses))]
}

func randomParticipantType() models.ParticipantTypeEnum {
	types := []models.ParticipantTypeEnum{models.ClientParticipant, models.CourierParticipant}
	return types[rand.Intn(len(types))]
}

func randomStatus() models.StatusEnum {
	statuses := []models.StatusEnum{models.Open, models.Closed}
	return statuses[rand.Intn(len(statuses))]
}

func randomSenderType() models.SenderTypeEnum {
	types := []models.SenderTypeEnum{models.ClientSender, models.CourierSender, models.SupportStaffSender}
	return types[rand.Intn(len(types))]
}

func randomUrgency() models.UrgencyEnum {
	urgencies := []models.UrgencyEnum{models.Urgent, models.Scheduled}
	return urgencies[rand.Intn(len(urgencies))]
}

func randomCurrentStatus() models.CurrentStatusEnum {
	statuses := []models.CurrentStatusEnum{models.WaitingForCourier, models.InTransit, models.Delivered}
	return statuses[rand.Intn(len(statuses))]
}

func randomPaymentStatus() models.PaymentStatusEnum {
	statuses := []models.PaymentStatusEnum{models.Paid, models.Unpaid}
	return statuses[rand.Intn(len(statuses))]
}

func randomPromoCodeType() models.PromoCodeTypeEnum {
	types := []models.PromoCodeTypeEnum{models.Discount, models.AdditionalService}
	return types[rand.Intn(len(types))]
}

func randomDeliveryType() models.DeliveryTypeEnum {
	types := []models.DeliveryTypeEnum{models.DeliveryUrgent, models.DeliveryScheduled}
	return types[rand.Intn(len(types))]
}

func randomJobTitle() string {
	jobTitles := []string{"Manager", "Engineer", "Developer", "Consultant", "Analyst", "Specialist", "Coordinator"}
	return jobTitles[rand.Intn(len(jobTitles))]
}

func randomPaymentMethod() string {
	methods := []string{"credit card", "cash"}
	return methods[rand.Intn(len(methods))]
}

func SeedData(db *gorm.DB) {
	rand.Seed(time.Now().UnixNano())

	// Начинаем транзакцию
	tx := db.Begin()

	// Обработка ошибок транзакции
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			fmt.Println("Транзакция откатилась из-за ошибки")
		}
	}()

	// Генерация данных
	supportStaffs := generateSupportStaff(tx)
	clients := generateClients(tx)
	couriers := generateCouriers(tx)
	promoCodes := generatePromoCodes(tx, clients)
	rates := generateRates(tx)
	generateOrders(tx, clients, couriers, promoCodes, rates)
	chats := generateChats(tx, supportStaffs, clients, couriers)
	generateMessages(tx, chats, clients)

	// Завершаем транзакцию
	if err := tx.Commit().Error; err != nil {
		log.Fatalf("Ошибка при коммите транзакции: %v", err)
	}
	fmt.Println("Транзакция завершена успешно")
}

func generateSupportStaff(tx *gorm.DB) []models.SupportStaff {
	var supportStaffs []models.SupportStaff
	for i := 0; i < 150; i++ {
		staff := models.SupportStaff{
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Phone:     faker.Phonenumber(),
			Email:     faker.Email(),
			Position:  randomJobTitle(),
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
		client := models.Client{
			Role:       randomRole(),
			FirstName:  faker.FirstName(),
			LastName:   faker.LastName(),
			MiddleName: faker.Name(),
			Phone:      phone,
			Email:      faker.Email(),
			Address:    fmt.Sprintf("%s, %s, %s, %s", randomdata.Street(), randomdata.City(), randomdata.State(randomdata.Small), randomdata.PostalCode("RU")),
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
		courier := models.Courier{
			EmploymentStatus:   randomEmploymentStatus(),
			TransportType:      randomTransportType(),
			AvailabilityStatus: randomAvailabilityStatus(),
			FirstName:          faker.FirstName(),
			LastName:           faker.LastName(),
			Phone:              phone,
			Photo:              faker.URL(),
			Passport:           faker.Word(),
			GPSCoordinates:     fmt.Sprintf("%f,%f", faker.Latitude(), faker.Longitude()),
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
	for i := 0; i < 500; i++ {
		clientID := clients[rand.Intn(len(clients))].ID
		promo := models.PromoCode{
			Type:           randomPromoCodeType(),
			ClientID:       &clientID,
			Code:           fmt.Sprintf("%s-%d", faker.Word(), i),
			DiscountAmount: rand.Float64() * 20,
			ValidUntil:     time.Now().AddDate(0, rand.Intn(12), rand.Intn(30)),
			Personalized:   rand.Intn(2) == 0,
		}
		if err := tx.Create(&promo).Error; err != nil {
			tx.Rollback()
			log.Fatalf("Ошибка при создании промокода: %v", err)
		}
		promoCodes = append(promoCodes, promo)
	}
	fmt.Println("Сгенерированы промокоды")
	return promoCodes
}

func generateRates(tx *gorm.DB) []models.Rate {
	var rates []models.Rate
	for i := 0; i < 50; i++ {
		rate := models.Rate{
			DeliveryType:  randomDeliveryType(),
			TransportType: randomTransportType(),
			Name:          fmt.Sprintf("Rate %d", i+1),
			Price:         rand.Float64() * 100,
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
	for i := 0; i < 50000; i++ {
		order := models.Order{
			Urgency:            randomUrgency(),
			SenderID:           clients[rand.Intn(len(clients))].ID,
			RecipientID:        clients[rand.Intn(len(clients))].ID,
			CourierID:          couriers[rand.Intn(len(couriers))].ID,
			CurrentStatus:      randomCurrentStatus(),
			PaymentStatus:      randomPaymentStatus(),
			RateID:             rates[rand.Intn(len(rates))].ID,
			PromoCodeID:        &promoCodes[rand.Intn(len(promoCodes))].ID,
			CreationDate:       time.Now(),
			ItemType:           faker.Word(),
			ItemValue:          rand.Float64() * 100,
			Weight:             rand.Float64() * 10,
			DiscountSurcharges: rand.Float64() * 10,
			PaymentMethod:      randomPaymentMethod(),
		}
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
		participantType := randomParticipantType()
		var participantID uint

		// Определяем участника (client или courier)
		if participantType == models.ClientParticipant && len(clients) > 0 {
			participantID = clients[rand.Intn(len(clients))].ID
		} else if participantType == models.CourierParticipant && len(couriers) > 0 {
			participantID = couriers[rand.Intn(len(couriers))].ID
		} else {
			continue
		}

		supportStaffID := supportStaffs[rand.Intn(len(supportStaffs))].ID
		chat := models.Chat{
			ParticipantType: participantType,
			ParticipantID:   participantID,
			SupportStaffID:  &supportStaffID,
			Status:          randomStatus(),
			CreationDate:    time.Now(),
			Reason:          faker.Sentence(),
		}

		if err := tx.Create(&chat).Error; err != nil {
			tx.Rollback()
			log.Fatalf("Ошибка при создании чата: %v", err)
		}
		chats = append(chats, chat)
	}
	fmt.Println("Сгенерированы чаты")
	return chats
}

func generateMessages(tx *gorm.DB, chats []models.Chat, clients []models.Client) {
	for i := 0; i < 8000; i++ {
		chatID := chats[rand.Intn(len(chats))].ID
		senderID := clients[rand.Intn(len(clients))].ID

		message := models.Message{
			ChatID:     chatID,
			SenderType: randomSenderType(),
			SenderID:   senderID,
			Timestamp:  time.Now(),
			Text:       faker.Sentence(),
		}

		if err := tx.Create(&message).Error; err != nil {
			tx.Rollback()
			log.Fatalf("Ошибка при создании сообщения: %v", err)
		}
	}
	fmt.Println("Сгенерированы сообщения")
}
