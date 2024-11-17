package migrations

import (
	"ORM_DB/internal/models"
	"gorm.io/gorm"
	"log"
)

// MigrateTables Функция для миграции таблиц
func MigrateTables(db *gorm.DB) {
	modelsToMigrate := []interface{}{
		&models.Client{},
		&models.Rate{},
		&models.SupportStaff{},
		&models.PromoCode{},
		&models.Courier{},
		&models.Chat{},
		&models.Message{},
		&models.Order{},
	}

	for _, model := range modelsToMigrate {
		if !db.Migrator().HasTable(model) {
			if err := db.Migrator().CreateTable(model); err != nil {
				log.Fatalf("failed to migrate model: %v", err)
			}
		} else {
			log.Printf("Table for model %T already exists, skipping migration\n", model)
		}
	}
}
