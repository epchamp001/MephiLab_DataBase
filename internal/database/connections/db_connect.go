package connections

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// Устанавливаем соединение с базой данных
func Connect(connectionString string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	return db
}
