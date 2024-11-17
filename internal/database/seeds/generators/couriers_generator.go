package generators

import (
	"ORM_DB/internal/database/seeds/utils"
	"ORM_DB/internal/models"
	"gorm.io/gorm"
)

func GenerateCouriers(tx *gorm.DB, count int) ([]models.Courier, error) {
	var couriers []models.Courier

	for i := 0; i < count; i++ {
		firstName, lastName, _ := utils.GenerateFullRussianName()

		courier := models.Courier{
			EmploymentStatus:   utils.RandomEmploymentStatus(),
			TransportType:      utils.RandomTransportType(),
			AvailabilityStatus: utils.RandomAvailabilityStatus(),
			FirstName:          firstName,
			LastName:           lastName,
			Phone:              utils.GenerateRussianPhoneNumber(),
			Photo:              utils.GeneratePhotoPath(),
			Passport:           utils.GeneratePassportNumber(),
			GPSCoordinates:     utils.GenerateMoscowCoordinates(),
		}

		if err := tx.Create(&courier).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		couriers = append(couriers, courier)
	}

	return couriers, nil
}
