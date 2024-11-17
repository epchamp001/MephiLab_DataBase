package parsers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Data struct {
	FemaleNames            []string `json:"femaleNames"`
	MaleNames              []string `json:"maleNames"`
	StreetsInMoscow        []string `json:"streetsInMoscow"`
	RegionsInMoscow        []string `json:"regionsInMoscow"`
	ItemTypes              []string `json:"itemTypes"`
	CourierMessages        []string `json:"courierMessages"`
	SupportMessagesCourier []string `json:"supportMessagesCourier"`
	SupportMessagesClient  []string `json:"supportMessagesClient"`
	ClientMessages         []string `json:"clientMessages"`
	RussianMaleSurnames    []string `json:"russianMaleSurnames"`
	RussianFemaleSurnames  []string `json:"russianFemaleSurnames"`
	ReasonChatCourier      []string `json:"reasonChatCourier"`
	ReasonChatClient       []string `json:"reasonChatClient"`
	FemalePatronymics      []string `json:"femalePatronymics"`
	MalePatronymics        []string `json:"malePatronymics"`
	Rates                  []string `json:"rates"`
}

var GlobalData Data

// LoadDataFromJSON загружает данные из JSON-файла в структуру Data.
func LoadDataFromJSON(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Ошибка при открытии файла '%s': %v", filename, err)
		return fmt.Errorf("ошибка при открытии файла: %w", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&GlobalData)
	if err != nil {
		log.Printf("Ошибка при декодировании JSON из файла '%s': %v", filename, err)
		return fmt.Errorf("ошибка при декодировании JSON: %w", err)
	}

	return nil
}
