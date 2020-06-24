package main

import (
	"strconv"
	"strings"
)
func loadCountriesFromString(str string) []Country{
	var countries = make([]Country,30)
	var cnt = 0
	var separatedStrings = strings.Split(str,";")
	for cnt=0;cnt<len(separatedStrings);cnt++{
		countries[cnt] = loadCountryFromString(separatedStrings[cnt])
	}
	return countries
}
func loadCountryFromString(str string) Country {
	//[TESTED]
	var splits = strings.Split(str,":")
	var newGDP, _ = strconv.Atoi(splits[0])
	var newMoney, _ = strconv.Atoi(splits[1])
	var newDebt, _ = strconv.Atoi(splits[2])
	var newPop, _ = strconv.Atoi(splits[3])
	var newInfrastructureScore, _ = strconv.Atoi(splits[4])
	var newMilitaryScore, _ = strconv.Atoi(splits[5])
	var newCultureScore, _ = strconv.Atoi(splits[6])
	var newCode, _ = strconv.Atoi(splits[7])
	var newIncome, _ = strconv.Atoi(splits[8])
	var newExpenses, _= strconv.Atoi(splits[9])
	return Country{
		Gdp:                 uint32(newGDP),
		Money:               uint32(newMoney),
		Debt:                uint32(newDebt),
		Population:          uint32(newPop),
		InfrastructureScore: uint32(newInfrastructureScore),
		MilitaryScore:       uint32(newMilitaryScore),
		CultureScore:        uint32(newCultureScore),
		Code:                uint32(newCode),
		Income:              uint32(newIncome),
		Expenses:            uint32(newExpenses),
	}
}
