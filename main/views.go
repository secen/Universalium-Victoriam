package main

import "fmt"

type VIEW int

const (
	DefaultView    VIEW = 0
	FinanceView    VIEW = 1
	LawView        VIEW = 2
	PorductionView VIEW = 3
	MilitaryView   VIEW = 4
	TechView       VIEW = 5
	NationPicker VIEW = 6
)

func VIEWShowNationPickerHelpData() {
	fmt.Println(readFromFile("\\consoleData\\nationPickerHelpData.txt"))
}
func VIEWShowCountryFinances(pickedNation Country) {
	fmt.Println("Treasury: ", pickedNation.Money)
	fmt.Println("Debt: ", pickedNation.Debt)
	fmt.Println("Income: ", pickedNation.Income)
	fmt.Println("Expenses: ", pickedNation.Expenses)
	fmt.Println("Interest: ", pickedNation.interest)
	fmt.Println("Balance: ", int32(pickedNation.Income)-int32(pickedNation.Expenses))
	fmt.Println("To raise taxes please type 'raise tax'")
	fmt.Println("To lower taxes please type 'lower tax'")
}

func VIEWWRITECountryTechs(pickedNation Country) {
	//writes the technology interface on the console
	//TODO: IMPLEMENT MILITARY INTERFACE
}

func VIEWShowCountryMilitary(pickedNation Country) {
	//writes the military interface on the console
	//TODO: IMPLEMENT MILITARY INTERFACE
}

func VIEWWRITECountryProduction(pickedNation Country) {
	//writes the production interface on the console
	//TODO: IMPLEMENT PRODUCTION INTERFACE
}

func VIEWWRITECountryLaws(pickedNation Country) {
	//writes the law interface on the console
	fmt.Println("Laws Enacted:")
	for i, s := range pickedNation.laws {
		fmt.Println(i, s.name)
	}
	fmt.Println("To propose new legistlation please type 'propose'")
	fmt.Println("To abolish legislation" +
		" please type 'abolish' and the number of the law")
	fmt.Println("-> ")
}

func VIEWShowCountryIntroductionPopup(pickedNation Country) {
	fmt.Println("The nation of ", countryCodes[(int)(pickedNation.Code)])
	showCountryStats(pickedNation)
	fmt.Println("May you be successful in your endevours!")
}
