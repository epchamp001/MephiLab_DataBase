package seed

import (
	"ORM_DB/internal/database/seed/RandomData"
	"fmt"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

func SeedData(db *gorm.DB) {
	rand.Seed(time.Now().UnixNano())

	// Загружаем данные из Excel
	err := RandomData.LoadDataFromExcel("internal/database/seed/RandomData/data.xlsx")
	if err != nil {
		log.Fatalf("Ошибка загрузки данных из файла Excel: %v", err)
	}
	fmt.Println("Данные из Excel загружены успешно")

	// Начинаем транзакцию
	tx := db.Begin()

	// Обработка ошибок транзакции
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			fmt.Printf("Транзакция откатилась из-за ошибки: %v\n", r)
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
	generateMessages(tx, chats, clients, couriers, supportStaffs)

	// Завершаем транзакцию
	if err := tx.Commit().Error; err != nil {
		log.Fatalf("Ошибка при коммите транзакции: %v", err)
	}
	fmt.Println("Транзакция завершена успешно")
}
