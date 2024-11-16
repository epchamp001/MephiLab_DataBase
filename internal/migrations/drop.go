package migrations

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

// Удаление всех таблиц и ENUM полей
func DropAllTables(db *gorm.DB) {
	// Удаление всех таблиц
	tables := []string{
		"support_staffs", "clients", "couriers", "chats", "messages", "orders", "promo_codes", "rates",
	}

	for _, table := range tables {
		if db.Migrator().HasTable(table) {
			err := db.Migrator().DropTable(table)
			if err != nil {
				log.Fatalf("Failed to drop table %s: %v", table, err)
			}
			fmt.Printf("Table %s dropped successfully\n", table)
		} else {
			fmt.Printf("Table %s does not exist, skipping drop\n", table)
		}
	}
}
