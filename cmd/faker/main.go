package main

import (
	"ORM_BD/internal/config"
	"ORM_BD/internal/database"
	"ORM_BD/internal/database/seed"
	"ORM_BD/internal/migrations"
	"fmt"
)

func main() {
	// Загружаем переменные окружения
	config.LoadEnv()

	// Устанавливаем соединение с базой данных
	db := database.Connect()

	// Миграция таблиц
	migrations.MigrateTables(db)
	fmt.Println("Миграция таблиц завершена успешно")

	// Генерация данных
	seed.ClearDatabase(db)
	seed.SeedData(db)
	fmt.Println("Генерация данных завершена успешно")

	fmt.Println("Приложение завершено успешно")
}
