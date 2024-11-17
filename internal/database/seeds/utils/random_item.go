package utils

import (
	randData "ORM_DB/internal/database/seeds/parsers"
	"math/rand"
)

func RandomItemType() string {
	if len(randData.ItemTypes) == 0 {
		return "Данные для типов товаров не загружены"
	}
	return randData.ItemTypes[rand.Intn(len(randData.ItemTypes))]
}
