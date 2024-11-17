package utils

import (
	"ORM_DB/internal/database/seeds/parsers"
	"math/rand"
)

func GenerateRussianFirstName() (string, string) {
	if rand.Intn(2) == 0 {
		return parsers.GlobalData.MaleNames[rand.Intn(len(parsers.GlobalData.MaleNames))], "male"
	}
	return parsers.GlobalData.FemaleNames[rand.Intn(len(parsers.GlobalData.FemaleNames))], "female"
}

func GenerateRussianLastName(gender string) string {
	if gender == "male" {
		return parsers.GlobalData.RussianMaleSurnames[rand.Intn(len(parsers.GlobalData.RussianMaleSurnames))]
	}
	return parsers.GlobalData.RussianFemaleSurnames[rand.Intn(len(parsers.GlobalData.RussianFemaleSurnames))]
}

func GenerateMiddleNameByGender(gender string) string {
	if gender == "male" {
		return parsers.GlobalData.MalePatronymics[rand.Intn(len(parsers.GlobalData.MalePatronymics))]
	}
	return parsers.GlobalData.FemalePatronymics[rand.Intn(len(parsers.GlobalData.FemalePatronymics))]
}

func GenerateFullRussianName() (string, string, string) {
	firstName, gender := GenerateRussianFirstName()
	lastName := GenerateRussianLastName(gender)
	middleName := GenerateMiddleNameByGender(gender)
	return firstName, lastName, middleName
}
