package main

import (
	"ORM_DB/internal/config"
	"ORM_DB/internal/database/seeds/generators"
	"ORM_DB/internal/database/seeds/parsers"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

func SeedData(db *gorm.DB, dataGen config.DataGenerationConfig) {
	rand.Seed(time.Now().UnixNano())

	// Загружаем данные из Excel
	data, err := parsers.LoadDataFromExcel("internal/database/seeds/parsers/data.xlsx")
	if err != nil {
		log.Fatalf("Ошибка загрузки данных из файла Excel: %v", err)
	}
	log.Println("Данные из Excel загружены успешно")

	// Начинаем транзакцию
	tx := db.Begin()

	// Обработка ошибок транзакции
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Транзакция откатилась из-за ошибки: %v\n", r)
		}
	}()

	// Генерация данных
	supportStaffs := generators.GenerateSupportStaff(tx, dataGen.SupportStaff)
	clients := generators.GenerateClients(tx, dataGen.Clients)
	couriers := generators.GenerateCouriers(tx, dataGen.Courier)
	promoCodes := generators.GeneratePromoCodes(tx, clients)
	rates := generators.GenerateRates(tx, dataGen.Rate)
	generators.GenerateOrders(tx, clients, couriers, promoCodes, rates, dataGen.Orders)
	chats := generators.GenerateChats(tx, supportStaffs, clients, couriers, dataGen.Chat)
	generators.GenerateMessages(tx, chats, clients, couriers, supportStaffs, dataGen.Messages)

	// Завершаем транзакцию
	if err := tx.Commit().Error; err != nil {
		log.Fatalf("Ошибка при коммите транзакции: %v", err)
	}
	log.Println("Транзакция завершена успешно")
}
