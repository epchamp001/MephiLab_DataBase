package seeder

import (
	"ORM_DB/internal/config"
	"ORM_DB/internal/database/seeds/generators"
	"ORM_DB/internal/database/seeds/parsers"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

func SeedData(db *gorm.DB, countGen config.DataGenerationConfig) {
	rand.Seed(time.Now().UnixNano())

	err := parsers.LoadDataFromJSON("internal/database/seeds/parsers/data_for_generation.json")
	if err != nil {
		log.Println(err)
		return
	}

	// Начинаем транзакцию
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Транзакция откатилась из-за ошибки: %v\n", r)
		}
	}()

	rates, _ := generators.GenerateRates(tx, countGen.Rate)
	supportStaffs, _ := generators.GenerateSupportStaff(tx, countGen.SupportStaff)
	clients, _ := generators.GenerateClients(tx, countGen.Clients)
	couriers, _ := generators.GenerateCouriers(tx, countGen.Courier)
	promoCodes, _ := generators.GeneratePromoCodes(tx, clients, countGen.PromoCodes)
	orders, _ := generators.GenerateOrders(tx, clients, couriers, promoCodes, rates, countGen.Orders)
	chats, _ := generators.GenerateChats(tx, clients, couriers, supportStaffs, countGen.Chat)
	messages, _ := generators.GenerateMessages(tx, chats, clients, couriers, supportStaffs, countGen.Messages)

	// Заглушки, возможно могут понадобится эти массивы
	_ = orders
	_ = messages

	// Завершение транзакции
	if err := tx.Commit().Error; err != nil {
		log.Fatalf("Ошибка при коммите транзакции: %v", err)
	}
	log.Println("Транзакция завершена успешно")
}
