package seed

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

// ClearDatabase очищает таблицы базы данных
func ClearDatabase(db *gorm.DB) {
	query := `
		TRUNCATE TABLE messages, chats, orders, rates, promo_codes, couriers, clients, support_staffs 
		RESTART IDENTITY CASCADE;
	`

	if err := db.Exec(query).Error; err != nil {
		log.Fatalf("Ошибка при очистке базы данных: %v", err)
	}

	fmt.Println("База данных очищена")
}
