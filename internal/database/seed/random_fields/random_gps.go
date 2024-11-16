package random_fields

import (
	"fmt"
	"math/rand"
)

// Генерация координат по Москве
func GenerateMoscowCoordinates() string {
	// Диапазон широты
	minLat := 55.55
	maxLat := 55.95

	// Диапазон долготы
	minLon := 37.36
	maxLon := 37.84

	// Случайное значение широты и долготы
	latitude := minLat + rand.Float64()*(maxLat-minLat)
	longitude := minLon + rand.Float64()*(maxLon-minLon)

	// Возврат координат в формате "широта,долгота"
	return fmt.Sprintf("%.6f,%.6f", latitude, longitude)
}
