package seeder

import (
	"ORM_DB/internal/config"
	"ORM_DB/internal/database/seeds/generators"
	"ORM_DB/internal/database/seeds/parsers"
	"fmt"
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

	fmt.Println(rates[1], supportStaffs[1], clients[1], couriers[1], promoCodes[1], orders[1], chats[1])
	//generators.GenerateMessages(tx, chats, clients, couriers, supportStaffs, countGen.Messages)

	// Завершение транзакции
	if err := tx.Commit().Error; err != nil {
		log.Fatalf("Ошибка при коммите транзакции: %v", err)
	}
	log.Println("Транзакция завершена успешно")
}