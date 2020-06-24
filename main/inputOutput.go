package main

import (
	"strconv"
	"strings"
)
func loadCountriesFromString(str string) []country{
	var countries = make([]country,30)
	var cnt = 0
	var separatedStrings = strings.Split(str,";")
	for cnt=0;cnt<len(separatedStrings);cnt++{
		countries[cnt] = loadCountryFromString(separatedStrings[cnt])
	}
	return countries
}
func loadCountryFromString(str string) country {
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
	return country{
		gdp: 				 uint32(newGDP),
		money:               uint32(newMoney),
		debt:                uint32(newDebt),
		population:          uint32(newPop),
		infrastructureScore: uint32(newInfrastructureScore),
		militaryScore:       uint32(newMilitaryScore),
		cultureScore:        uint32(newCultureScore),
		code:                uint32(newCode),
		income:              uint32(newIncome),
		expenses:            uint32(newExpenses),
	}
}
