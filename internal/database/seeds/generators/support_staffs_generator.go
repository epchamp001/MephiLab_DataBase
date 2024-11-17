package generators

import (
	"ORM_DB/internal/database/seeds/utils"
	"ORM_DB/internal/models"
	"gorm.io/gorm"
)

func GenerateSupportStaff(tx *gorm.DB, count int) ([]models.SupportStaff, error) {
	var supportStaffs []models.SupportStaff

	for i := 0; i < count; i++ {
		firstName, lastName, _ := utils.GenerateFullRussianName()
		email := utils.GenerateEmail(firstName, lastName)

		ss := models.SupportStaff{
			FirstName: firstName,
			LastName:  lastName,
			Phone:     utils.GenerateRussianPhoneNumber(),
			Email:     email,
			Position:  utils.RandomJobTitle(),
		}

		if err := tx.Create(&ss).Error; err != nil {
			return nil, err
		}
		supportStaffs = append(supportStaffs, ss)
	}

	return supportStaffs, nil
}
