package parsers

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
)

type Data struct {
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
}

// LoadDataFromExcel загружает данные из data.xlsx и возвращает ошибку в случае сбоя
func LoadDataFromExcel(filePath string) (*Data, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("ошибка при открытии файла %s: %w", filePath, err)
	}

	data := &Data{}
	data.FemaleNames, err = loadColumnData(f, "femaleName")
	data.MaleNames, err = loadColumnData(f, "maleName")
	data.StreetsInMoscow, err = loadColumnData(f, "streetsInMoscow")
	data.RegionsInMoscow, err = loadColumnData(f, "regionsInMoscow")
	data.ItemTypes, err = loadColumnData(f, "itemTypes")
	data.CourierMessages, err = loadColumnData(f, "courierMessages")
	data.SupportMessagesCourier, err = loadColumnData(f, "supportMessagesCourier")
	data.SupportMessagesClient, err = loadColumnData(f, "supportMessagesClient")
	data.ClientMessages, err = loadColumnData(f, "clientMessages")
	data.RussianMaleSurnames, err = loadColumnData(f, "russianMaleSurnames")
	data.RussianFemaleSurnames, err = loadColumnData(f, "russianFemaleSurnames")
	data.ReasonChatCourier, err = loadColumnData(f, "reasonChatCourier")
	data.ReasonChatClient, err = loadColumnData(f, "reasonChatClient")
	data.FemalePatronymics, err = loadColumnData(f, "femalePatronymics")
	data.MalePatronymics, err = loadColumnData(f, "malePatronymics")
	data.Rates, err = loadColumnData(f, "rates")

	if err != nil {
		log.Printf("Ошибка при загрузке данных: %v", err)
		return data, err
	}

	return data, nil
}
