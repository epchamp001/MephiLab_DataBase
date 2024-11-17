package generators

import (
	"ORM_DB/internal/database/seeds/utils"
	"ORM_DB/internal/models"
	"gorm.io/gorm"
)

func GenerateClients(tx *gorm.DB, count int) ([]models.Client, error) {
	var clients []models.Client

	for i := 0; i < count; i++ {
		firstName, lastName, middleName := utils.GenerateFullRussianName()
		email := utils.GenerateEmail(firstName, lastName)

		client := models.Client{
			Role:       utils.RandomRole(), // Случайно выбираем роль клиента
			FirstName:  firstName,
			LastName:   lastName,
			MiddleName: middleName,
			Phone:      utils.GenerateRussianPhoneNumber(),
			Email:      email,
			Address:    utils.GenerateMoscowAddress(),
		}

		if err := tx.Create(&client).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		clients = append(clients, client)
	}

	return clients, nil
}
