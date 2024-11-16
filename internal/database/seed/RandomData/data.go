package RandomData

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
)

var (
	FemaleNames            []string
	MaleNames              []string
	StreetsInMoscow        []string
	RegionsInMoscow        []string
	ItemTypes              []string
	CourierMessages        []string
	SupportMessagesCourier []string
	SupportMessagesClient  []string
	ClientMessages         []string
	RussianMaleSurnames    []string
	RussianFemaleSurnames  []string
	ReasonChatCourier      []string
	ReasonChatClient       []string
	FemalePatronymics      []string
	MalePatronymics        []string
	Rates                  []string
)

// LoadDataFromExcel загружает данные из data.xlsx и возвращает ошибку в случае сбоя
func LoadDataFromExcel(filePath string) error {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return fmt.Errorf("ошибка при открытии файла %s: %w", filePath, err)
	}
	FemaleNames, err = loadColumnData(f, "femaleName")
	MaleNames, err = loadColumnData(f, "maleName")
	StreetsInMoscow, err = loadColumnData(f, "streetsInMoscow")
	RegionsInMoscow, err = loadColumnData(f, "regionsInMoscow")
	ItemTypes, err = loadColumnData(f, "itemTypes")
	CourierMessages, err = loadColumnData(f, "courierMessages")
	SupportMessagesCourier, err = loadColumnData(f, "supportMessagesCourier")
	SupportMessagesClient, err = loadColumnData(f, "supportMessagesClient")
	ClientMessages, err = loadColumnData(f, "clientMessages")
	RussianMaleSurnames, err = loadColumnData(f, "russianMaleSurnames")
	RussianFemaleSurnames, err = loadColumnData(f, "russianFemaleSurnames")
	ReasonChatCourier, err = loadColumnData(f, "reasonChatCourier")
	ReasonChatClient, err = loadColumnData(f, "reasonChatClient")
	FemalePatronymics, err = loadColumnData(f, "femalePatronymics")
	MalePatronymics, err = loadColumnData(f, "malePatronymics")
	Rates, err = loadColumnData(f, "rates")

	log.Println("Все данные успешно загружены из файла Excel")
	return nil
}

// loadColumnData загружает данные из указанного столбца листа Excel
func loadColumnData(f *excelize.File, columnName string) ([]string, error) {
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении листа Sheet1: %w", err)
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

	return nil, fmt.Errorf("колонка %s не найдена в Sheet1", columnName)
}
