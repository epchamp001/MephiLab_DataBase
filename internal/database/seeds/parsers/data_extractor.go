package parsers

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
)

// loadColumnData загружает данные из указанного столбца листа Excel
func loadColumnData(f *excelize.File, columnName string) ([]string, error) {
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Printf("ошибка при чтении листа Sheet1: %v", err)
		return nil, err // Возвращаем исходную ошибку без изменений
	}

	var data []string
	for i, header := range rows[0] { // Предполагаем, что первая строка — заголовок
		if header == columnName {
			for _, row := range rows[1:] { // Пропускаем заголовок
				if len(row) > i && row[i] != "" { // Проверяем на пустую строку
					data = append(data, row[i])
				}
			}
			return data, nil
		}
	}

	errMsg := fmt.Sprintf("колонка %s не найдена в Sheet1", columnName)
	log.Println(errMsg)
	return nil, fmt.Errorf(errMsg) // Возвращаем ошибку с подробным описанием
}
