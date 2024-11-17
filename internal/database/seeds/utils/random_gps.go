package utils

import (
	"fmt"
	"math/rand"
)

// Генерация координат по Москве
func GenerateMoscowCoordinates() string {
	minLat := 55.55
	maxLat := 55.95

	minLon := 37.36
	maxLon := 37.84

	latitude := minLat + rand.Float64()*(maxLat-minLat)
	longitude := minLon + rand.Float64()*(maxLon-minLon)

	return fmt.Sprintf("%.6f,%.6f", latitude, longitude)
}
