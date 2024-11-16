package migrations

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

// Удаление всех таблиц и ENUM полей
func DropAllTablesAndEnums(db *gorm.DB) {
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

	// Удаление ENUM полей
	enums := []string{
		"role_enum", "employment_status_enum", "transport_type_enum", "availability_status_enum",
		"participant_type_enum", "status_enum", "sender_type_enum", "urgency_enum", "current_status_enum",
		"payment_status_enum", "promo_code_type_enum", "delivery_type_enum",
	}

	for _, enum := range enums {
		query := fmt.Sprintf("DROP TYPE IF EXISTS %s CASCADE;", enum)
		err := db.Exec(query).Error
		if err != nil {
			log.Fatalf("Failed to drop enum %s: %v", enum, err)
		}
		fmt.Printf("Enum %s dropped successfully\n", enum)
	}
}
