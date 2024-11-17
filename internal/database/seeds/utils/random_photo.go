package utils

import (
	"fmt"
	"math/rand"
)

func GeneratePhotoPath() string {
	photoNumber := rand.Intn(100) + 1 // Генерируем номер фото от 1 до 100
	extensions := []string{"png", "jpg"}
	extension := extensions[rand.Intn(len(extensions))]

	return fmt.Sprintf("photo%d.%s", photoNumber, extension)
}
