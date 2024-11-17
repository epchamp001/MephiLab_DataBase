package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"unicode"
)

func Transliterate(input string) string {
	translitMap := map[rune]string{
		'а': "a", 'б': "b", 'в': "v", 'г': "g", 'д': "d", 'е': "e", 'ё': "e", 'ж': "zh",
		'з': "z", 'и': "i", 'й': "y", 'к': "k", 'л': "l", 'м': "m", 'н': "n", 'о': "o",
		'п': "p", 'р': "r", 'с': "s", 'т': "t", 'у': "u", 'ф': "f", 'х': "h", 'ц': "ts",
		'ч': "ch", 'ш': "sh", 'щ': "shch", 'ы': "y", 'э': "e", 'ю': "yu", 'я': "ya",
		'ъ': "", 'ь': "",
	}

	result := strings.Builder{}
	for _, c := range input {
		if val, ok := translitMap[c]; ok {
			result.WriteString(val)
		} else if unicode.IsLetter(c) || unicode.IsNumber(c) {
			result.WriteRune(c)
		}
	}
	return result.String()
}

func GenerateEmail(firstName, lastName string) string {
	// Нижний регистр и удаление пробелов
	cleanFirstName := strings.Join(strings.Fields(strings.ToLower(firstName)), "")
	cleanLastName := strings.Join(strings.Fields(strings.ToLower(lastName)), "")

	translitFirstName := Transliterate(cleanFirstName)
	translitLastName := Transliterate(cleanLastName)

	domains := []string{"@gmail.com", "@mail.ru"}
	domain := domains[rand.Intn(len(domains))]

	email := fmt.Sprintf("%s.%s%s", translitFirstName, translitLastName, domain)

	return email
}
