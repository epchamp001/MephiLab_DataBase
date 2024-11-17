package utils

import (
	"ORM_DB/internal/database/seeds/parsers"
	"fmt"
	"math/rand"
)

// GenerateMoscowAddress генерирует случайный адрес в Москве
func GenerateMoscowAddress() string {
	if len(parsers.GlobalData.StreetsInMoscow) == 0 || len(parsers.GlobalData.RegionsInMoscow) == 0 {
		return "Данные для адресов не загружены"
	}

	street := parsers.GlobalData.StreetsInMoscow[rand.Intn(len(parsers.GlobalData.StreetsInMoscow))]
	district := parsers.GlobalData.RegionsInMoscow[rand.Intn(len(parsers.GlobalData.RegionsInMoscow))]
	houseNumber := rand.Intn(100) + 1     // Генерируем номер дома от 1 до 100
	apartmentNumber := rand.Intn(200) + 1 // Генерируем номер квартиры от 1 до 200

	return fmt.Sprintf("%s, дом %d, кв. %d, район %s, Москва", street, houseNumber, apartmentNumber, district)
}
