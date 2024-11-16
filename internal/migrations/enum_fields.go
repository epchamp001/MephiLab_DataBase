package migrations

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

// Создание ENUM полей
func CreateEnums(db *gorm.DB) {
	enumQueries := []string{
		`CREATE TYPE role_enum AS ENUM ('sender', 'receiver');`,
		`CREATE TYPE employment_status_enum AS ENUM ('self-employed', 'official');`,
		`CREATE TYPE transport_type_enum AS ENUM ('on foot', 'car', 'truck');`,
		`CREATE TYPE availability_status_enum AS ENUM ('available', 'busy');`,
		`CREATE TYPE participant_type_enum AS ENUM ('client', 'courier');`,
		`CREATE TYPE status_enum AS ENUM ('open', 'closed');`,
		`CREATE TYPE sender_type_enum AS ENUM ('client', 'courier', 'support staff');`,
		`CREATE TYPE urgency_enum AS ENUM ('urgent', 'scheduled');`,
		`CREATE TYPE current_status_enum AS ENUM ('waiting for courier', 'in transit', 'delivered');`,
		`CREATE TYPE payment_status_enum AS ENUM ('paid', 'unpaid');`,
		`CREATE TYPE promo_code_type_enum AS ENUM ('discount', 'additional service');`,
		`CREATE TYPE delivery_type_enum AS ENUM ('urgent', 'scheduled');`,
	}

	for _, query := range enumQueries {
		err := db.Exec(query).Error
		if err != nil {
			log.Fatalf("Failed to create enum: %v", err)
		}
		fmt.Println("Enum created successfully")
	}
}
