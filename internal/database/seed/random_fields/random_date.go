package random_fields

import (
	"math/rand"
	"time"
)

// Генерация случайной даты в пределах 2024 года
func GenerateRandomDate2024() time.Time {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)

	// Разница между началом и концом в секундах
	delta := end.Unix() - start.Unix()
	seconds := rand.Int63n(delta)

	return start.Add(time.Duration(seconds) * time.Second)
}
