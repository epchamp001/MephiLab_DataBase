package utils

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
)

var (
	phoneNumbers = make(map[string]bool) // Для уникальных номеров
)

func GenerateRussianPhoneNumber() string {
	var phoneNumber string
	for {
		rawNumber := faker.Phonenumber()
		phoneNumber = fmt.Sprintf("+7-%s", rawNumber)

		if _, exists := phoneNumbers[phoneNumber]; !exists {
			phoneNumbers[phoneNumber] = true
			break
		}
	}
	return phoneNumber
}
