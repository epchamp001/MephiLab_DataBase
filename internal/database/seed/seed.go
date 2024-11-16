package seed

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

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
