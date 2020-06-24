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
	fmt.Println("##############################")
	fmt.Println("Type 'pick' to pick a nation")
	fmt.Println("Type 'play NATION' to play that nation")
	fmt.Println("For example: play FRANCE")
	fmt.Println("Type 'back' to go back to the main menu")
	fmt.Println("Type 'load' to load a game")
	fmt.Println("Type 'random' for a random nation pick")
	fmt.Println("Type 'exit' to exit the game")
	fmt.Println("Type 'clear' to clear the screen")
	fmt.Println("For more information, type 'help' ")
	fmt.Println("##############################")
}
func VIEWShowCountryFinances(pickedNation country) {
	fmt.Println("Treasury: ", pickedNation.money)
	fmt.Println("Debt: ", pickedNation.debt)
	fmt.Println("Income: ", pickedNation.income)
	fmt.Println("Expenses: ", pickedNation.expenses)
	fmt.Println("Interest: ", pickedNation.interest)
	fmt.Println("Balance: ", int32(pickedNation.income)-int32(pickedNation.expenses))
	fmt.Println("To raise taxes please type 'raise tax'")
	fmt.Println("To lower taxes please type 'lower tax'")
}

func VIEWWRITECountryTechs(pickedNation country) {
	//writes the technology interface on the console
	//TODO: IMPLEMENT MILITARY INTERFACE
}

func VIEWShowCountryMilitary(pickedNation country) {
	//writes the military interface on the console
	//TODO: IMPLEMENT MILITARY INTERFACE
}

func VIEWWRITECountryProduction(pickedNation country) {
	//writes the production interface on the console
	//TODO: IMPLEMENT PRODUCTION INTERFACE
}

func VIEWWRITECountryLaws(pickedNation country) {
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

func VIEWShowCountryIntroductionPopup(pickedNation country) {
	fmt.Println("The nation of ", countryCodes[(int)(pickedNation.code)])
	showCountryStats(pickedNation)
	fmt.Println("May you be successful in your endevours!")
}
