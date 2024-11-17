package cleanup

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

// ClearDatabase очищает таблицы от сгенерированных данных
func ClearDatabase(db *gorm.DB) error {
	tables, err := db.Migrator().GetTables()
	if err != nil {
		log.Printf("Ошибка при получении списка таблиц: %v", err)
		return fmt.Errorf("Ошибка при получении списка таблиц: %v", err)
	}

	for _, table := range tables {
		if err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE;", table)).Error; err != nil {
			log.Printf("Ошибка при очистке таблицы %s: %v", table, err)
			return fmt.Errorf("ошибка при очистке таблицы %s: %v", table, err)
		}
	}
	log.Println("База данных очищена")
	return nil
}
