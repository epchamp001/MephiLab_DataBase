package utils

import (
	"ORM_DB/internal/database/seeds/parsers"
	"math/rand"
)

func RandomItemType() string {
	if len(parsers.GlobalData.ItemTypes) == 0 {
		return "Данные для типов товаров не загружены"
	}
	return parsers.GlobalData.ItemTypes[rand.Intn(len(parsers.GlobalData.ItemTypes))]
}
