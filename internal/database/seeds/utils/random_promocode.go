package utils

import (
	"fmt"
	"math/rand"
)

func GenerateRandomPromoCode() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 7)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return fmt.Sprintf("PROMO-%s", string(b))
}
