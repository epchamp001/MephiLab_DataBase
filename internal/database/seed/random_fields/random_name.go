package random_fields

import (
	randData "ORM_DB/internal/database/seed/RandomData"
	"math/rand"
)

func GenerateRussianFirstName() (string, string) {
	// Возвращаем имя и пол
	if rand.Intn(2) == 0 { // 50% вероятность выбора мужского имени
		return randData.MaleNames[rand.Intn(len(randData.MaleNames))], "male"
	}
	return randData.FemaleNames[rand.Intn(len(randData.FemaleNames))], "female"
}

func GenerateRussianLastName(gender string) string {
	// Возвращаем фамилию в зависимости от пола
	if gender == "male" {
		return randData.RussianMaleSurnames[rand.Intn(len(randData.RussianMaleSurnames))]
	}
	return randData.RussianFemaleSurnames[rand.Intn(len(randData.RussianFemaleSurnames))]
}

func GenerateMiddleNameByGender(gender string) string {
	// Возвращаем отчество в зависимости от пола
	if gender == "male" {
		return randData.MalePatronymics[rand.Intn(len(randData.MalePatronymics))]
	}
	return randData.FemalePatronymics[rand.Intn(len(randData.FemalePatronymics))]
}

func GenerateFullRussianName() (string, string, string) {
	firstName, gender := GenerateRussianFirstName()  // Генерируем имя и определяем пол
	lastName := GenerateRussianLastName(gender)      // Генерируем фамилию на основе пола
	middleName := GenerateMiddleNameByGender(gender) // Генерируем отчество на основе пола
	return firstName, lastName, middleName           // Возвращаем ФИО
}
