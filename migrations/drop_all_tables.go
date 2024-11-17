package migrations

import (
	"gorm.io/gorm"
	"log"
)

// DropAllTables Удаление всех таблиц и ENUM полей
func DropAllTables(db *gorm.DB) {
	tables, err := db.Migrator().GetTables()
	if err != nil {
		log.Fatalf("Failed to retrieve tables: %v", err)
	}

	for _, table := range tables {
		err := db.Migrator().DropTable(table)
		if err != nil {
			log.Fatalf("Failed to drop table %s: %v", table, err)
		}
	}
}
