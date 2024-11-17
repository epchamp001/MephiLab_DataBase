package generators

import (
	"ORM_DB/internal/database/seeds/utils"
	"ORM_DB/internal/models"
	"gorm.io/gorm"
	"math/rand"
)

func GenerateOrders(tx *gorm.DB, clients []models.Client, couriers []models.Courier, promoCodes []models.PromoCode, rates []models.Rate, count int) ([]models.Order, error) {
	var orders []models.Order

	clientPromoCodes := make(map[uint]*models.PromoCode)
	for _, promoCode := range promoCodes {
		if promoCode.ClientID != nil {
			clientPromoCodes[*promoCode.ClientID] = &promoCode
		}
	}

	for i := 0; i < count; i++ {
		senderIndex := rand.Intn(len(clients))
		sender := clients[senderIndex]

		var recipient models.Client
		for {
			recipientIndex := rand.Intn(len(clients))
			recipient = clients[recipientIndex]
			if sender.ID != recipient.ID {
				break
			}
		}

		var promoCodeID *uint
		if promo, ok := clientPromoCodes[sender.ID]; ok {
			promoCodeID = &promo.ID
		}

		courier := couriers[rand.Intn(len(couriers))]
		rate := rates[rand.Intn(len(rates))]

		order := models.Order{
			Urgency:            utils.RandomUrgency(),
			SenderID:           sender.ID,
			RecipientID:        recipient.ID,
			CourierID:          courier.ID,
			PromoCodeID:        promoCodeID,
			RateID:             rate.ID,
			CurrentStatus:      utils.RandomCurrentStatus(),
			PaymentStatus:      utils.RandomPaymentStatus(),
			CreationDate:       utils.GenerateRandomDate2024(),
			ItemType:           utils.RandomItemType(),
			ItemValue:          float64(rand.Intn(1000)), // Значение предмета до 1000
			Weight:             float64(rand.Intn(50)),   // Вес до 50 кг
			DiscountSurcharges: float64(rand.Intn(50)),   // Скидки/надбавки до 50
			PaymentMethod:      utils.RandomPaymentMethod(),
		}

		if err := tx.Create(&order).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}
