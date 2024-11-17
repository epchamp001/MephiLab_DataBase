package generators

import (
	"ORM_DB/internal/database/seeds/utils"
	"ORM_DB/internal/models"
	"gorm.io/gorm"
	"math"
	"math/rand"
	"time"
)

func GeneratePromoCodes(tx *gorm.DB, clients []models.Client, count int) ([]models.PromoCode, error) {
	var promoCodes []models.PromoCode
	clientPromoCounts := make(map[uint]int) // Карта для отслеживания количества промокодов на клиента

	for _, client := range clients {
		clientPromoCounts[client.ID] = 0 // Инициализация карты счетчиком промокодов
	}

	for i := 0; i < count; i++ {
		promoCode := models.PromoCode{
			Type:           utils.RandomPromoCodeType(),
			Code:           utils.GenerateRandomPromoCode(),
			DiscountAmount: math.Round(rand.Float64() * 20),
			ValidUntil:     time.Now().AddDate(0, 1, 0),
			Personalized:   rand.Intn(2) == 0,
		}

		if promoCode.Personalized && len(clients) > 0 {
			client := clients[rand.Intn(len(clients))]
			promoCode.ClientID = &client.ID
		}

		if err := tx.Create(&promoCode).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		promoCodes = append(promoCodes, promoCode)
	}

	return promoCodes, nil
}
