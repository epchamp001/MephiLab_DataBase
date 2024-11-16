package main

import (
	"ORM_DB/internal/config"
	"ORM_DB/internal/database"
	"ORM_DB/internal/migrations"
	"fmt"
)

func main() {
	// Загружаем переменные окружения
	config.LoadEnv()

	// Устанавливаем соединение с базой данных
	db := database.Connect()

	//// Удаляем все таблицы и ENUM поля
	migrations.DropAllTables(db)

	// Миграция таблиц
	migrations.MigrateTables(db)
	fmt.Println("Миграция таблиц завершена успешно")

	//// Генерация данных
	//seed.ClearDatabase(db) // очищаю во избежание дублирование одинаковых рандомных id
	//seed.SeedData(db)
	//fmt.Println("Генерация данных завершена успешно")

	fmt.Println("Приложение завершено успешно")
}
