package main

import (
	"ORM_DB/internal/config"
	"ORM_DB/internal/database/connections"
	"ORM_DB/migrations"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func formatJSONFile(inputFile, outputFile string) error {
	// Чтение данных из входного файла
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("ошибка чтения файла %s: %v", inputFile, err)
	}

	// Создание временной структуры для проверки валидности JSON
	var jsonData interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return fmt.Errorf("ошибка парсинга JSON: %v", err)
	}

	// Форматирование JSON с отступами
	formattedJSON, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка форматирования JSON: %v", err)
	}

	// Запись отформатированного JSON в выходной файл
	err = ioutil.WriteFile(outputFile, formattedJSON, 0644)
	if err != nil {
		return fmt.Errorf("ошибка записи в файл %s: %v", outputFile, err)
	}

	fmt.Printf("JSON из файла %s успешно отформатирован и сохранен в файл %s\n", inputFile, outputFile)
	return nil
}

func main() {
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Устанавливаем соединение с базой данных
	db := connections.Connect(cfg.Database.URL)

	// Удаляем все таблицы
	migrations.DropAllTables(db)

	// Миграция таблиц
	migrations.MigrateTables(db)
	log.Println("Миграция таблиц завершена успешно")

	//Генерация данных
	//seeds.ClearDatabase(db) // очищаю во избежание дублирование одинаковых рандомных id
	//seeds.SeedData(db, cfg.DataGeneration)
	//log.Println("Генерация данных завершена успешно")

	log.Println("Приложение завершено успешно")
}
