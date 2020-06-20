package main

import (
	"fmt"
)

func gameReadyStart(nation string) {
	_, _ = initAll()
	fmt.Println(nation)
}

type VIEW int

const (
	DefaultView    VIEW = 0
	FinanceView    VIEW = 1
	LawView        VIEW = 2
	PorductionView VIEW = 3
	MilitaryView   VIEW = 4
	TechView       VIEW = 5
)

func showCountryIntroductionPopup(pickedNation country) {
	fmt.Println("The nation of ", countryCodes[(int)(pickedNation.code)])
	showCountryStats(pickedNation)
	fmt.Println("May you be successful in your endevours!")
}
func writeOutput(currentView VIEW, isFirstTime bool, pickedNation country) {
	if isFirstTime {
		showCountryIntroductionPopup(pickedNation)
		isFirstTime = false
	}
	writeOverview(currentView, pickedNation)
}
func writeFinances(pickedNation country) {
	fmt.Println("Treasury: ", pickedNation.money)
	fmt.Println("Debt: ", pickedNation.debt)
	fmt.Println("Income: ", pickedNation.income)
	fmt.Println("Expenses: ", pickedNation.expenses)
	fmt.Println("Interest: ", pickedNation.interest)
	fmt.Println("Balance: ", int32(pickedNation.income)-int32(pickedNation.expenses))
}
func writeOverview(currentView VIEW, pickedNation country) {
	switch currentView {
	case DefaultView:
		fmt.Println("type 'finances' or 'fin' for a detailed view of your finances")
		fmt.Println("type 'tech' or 'technology' for a detailed view of your technological progress")
		fmt.Println("type 'army' or 'mil' for a detailed view of your army")
		fmt.Println("type 'prod' or 'production' for a detailed view of your economic production")
		fmt.Println("type 'law' or 'gov' or 'government' for a detailed view of your realm's governing structure")
		break
	case FinanceView:
		writeFinances(pickedNation)
		break
	case MilitaryView:
		writeMilitary()
		break
	case PorductionView:
		writeProduction()
		break
	case LawView:
		writeLaw(pickedNation.laws)
		break
	case TechView:
		writeTech()
		break
	}
}

func writeTech() {
	//writes the technology interface on the console
	//TODO: IMPLEMENT MILITARY INTERFACE
}

func writeMilitary() {
	//writes the military interface on the console
	//TODO: IMPLEMENT MILITARY INTERFACE
}

func writeProduction() {
	//writes the production interface on the console
	//TODO: IMPLEMENT PRODUCTION INTERFACE
}

func writeLaw(laws []law) {
	//writes the law interface on the console
	fmt.Println("Laws Enacted:")
	for i, s := range laws {
		fmt.Println(i, s.name)
	}
	fmt.Println("To propose new legistlation please type 'propose'")
	fmt.Println("To abolish legislation" +
		" please type 'abolish' and the number of the law")
	fmt.Println("-> ")
}
func abolishLaw (i int, laws []law) []law {
	var laws2 = append(laws[:i], laws[i+1:]...);
	return laws2
}
func executeLogic() {
	executePlayerActions();
	executeAIActions();
	runNextStepInSimulation();
}

func runNextStepInSimulation() {
	executeUpdateEconomics();
	executeUpdateCountries();
	executeUpdateTroops();
}

func executeUpdateCountries() {

}

func executeUpdateEconomics() {
	
}

func executeUpdateTroops() {

}

func executeAIActions() {

}

func executePlayerActions() {

}
func handleGameInput() {
	fmt.Println("type help")
}
func Tick(currentView VIEW, isFirstTime bool, pickedNation country) {
	writeOutput(currentView, isFirstTime, pickedNation)
	executeLogic()
	handleGameInput()
}
