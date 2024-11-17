package migrations

import (
	"ORM_DB/internal/models"
	"gorm.io/gorm"
	"log"
)

// MigrateTables Функция для миграции таблиц
func MigrateTables(db *gorm.DB) {
	modelsToMigrate := []interface{}{
		&models.SupportStaff{},
		&models.Client{},
		&models.Courier{},
		&models.Chat{},
		&models.Message{},
		&models.Order{},
		&models.PromoCode{},
		&models.Rate{},
	}

	for _, model := range modelsToMigrate {
		if !db.Migrator().HasTable(model) {
			if err := db.AutoMigrate(model); err != nil {
				log.Fatalf("failed to migrate model: %v", err)
			}
			log.Printf("Table for model %T created successfully\n", model)
		} else {
			log.Printf("Table for model %T already exists, skipping migration\n", model)
		}
	}
}
