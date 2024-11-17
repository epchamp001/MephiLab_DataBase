package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomDate2024() time.Time {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)

	delta := end.Unix() - start.Unix()
	seconds := rand.Int63n(delta)

	return start.Add(time.Duration(seconds) * time.Second)
}
