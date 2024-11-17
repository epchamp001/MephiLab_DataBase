package generators

import (
	"ORM_DB/internal/database/seeds/utils"
	"ORM_DB/internal/models"
	"fmt"
	"gorm.io/gorm"
	"log"
	"math"
	"math/rand"
)

func GenerateRates(tx *gorm.DB, count int) ([]models.Rate, error) {
	var rates []models.Rate

	for i := 0; i < count; i++ {
		rate := models.Rate{
			DeliveryType:  utils.RandomDeliveryType(),
			TransportType: utils.RandomTransportType(),
			Name:          fmt.Sprintf("Тариф %d", i+1),
			Price:         math.Round(rand.Float64() * 100),
			Description:   utils.GetRandomRateDescription(),
		}

		// Пытаемся создать тариф в базе данных
		if err := tx.Create(&rate).Error; err != nil {
			tx.Rollback() // Откатываем транзакцию при ошибке
			log.Fatalf("Ошибка при создании тарифа: %v", err)
		}
		rates = append(rates, rate)
	}

	return rates, nil
}
