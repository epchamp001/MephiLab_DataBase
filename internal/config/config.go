package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

// Загружаем переменные окружения
func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Ошибка при загрузке .env файла: %v", err)
	}
	fmt.Println("Переменные окружения загружены")
}
