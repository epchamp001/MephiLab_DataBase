package utils

import (
	"ORM_DB/internal/database/seeds/parsers"
	"math/rand"
)

func GetRandomRateDescription() string {
	if len(parsers.GlobalData.Rates) == 0 {
		return "Дефолтное описание"
	}
	return parsers.GlobalData.Rates[rand.Intn(len(parsers.GlobalData.Rates))]
}
