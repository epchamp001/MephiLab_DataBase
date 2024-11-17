package main

import (
	"ORM_DB/internal/config"
	"ORM_DB/internal/database/connections"
	"ORM_DB/internal/database/seeds/seeder"
	"ORM_DB/migrations"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db := connections.Connect(cfg.Database.URL)

	migrations.DropAllTables(db)

	// Миграция таблиц
	migrations.MigrateTables(db)
	log.Println("Миграция таблиц завершена успешно")

	seeder.SeedData(db, cfg.DataGeneration)

	log.Println("Генерация данных завершена успешно")
	log.Println("Приложение завершено успешно")
}
