package main

import (
	"ORM_BD/database/seed"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"ORM_BD/models"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not found in environment")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	fmt.Println("Database connected successfully")

	// Создание таблиц, если они отсутствуют
	migrateTables(db)
	fmt.Println("Database tables checked and migrated successfully")

	// Генерация данных
	seed.SeedData(db)
	fmt.Println("Данные успешно сгенерированы")
}

// Функция для миграции таблиц, если они отсутствуют
func migrateTables(db *gorm.DB) {
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
			fmt.Printf("Table for model %T created successfully\n", model)
		} else {
			fmt.Printf("Table for model %T already exists, skipping migration\n", model)
		}
	}
}
