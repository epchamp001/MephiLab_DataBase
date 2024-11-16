package random_fields

import (
	"fmt"
	"math/rand"
	"time"
)

// Генерация номера паспорта (10 цифр)
func GeneratePassportNumber() string {
	rand.Seed(time.Now().UnixNano())
	series := rand.Intn(10000)                      // Генерация серии (4 цифры)
	number := rand.Intn(1000000)                    // Генерация номера (6 цифр)
	return fmt.Sprintf("%04d %06d", series, number) // Формат "Серия Номер"
}
