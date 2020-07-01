package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CONSOLEWRITEOverview(currentView *VIEW, pickedNation Country, globalEconomicQueue queue) {
	switch *currentView {
	case DefaultView:
		VIEWShowDEFAULTVIEW(pickedNation)
		break
	case NationPickerView:
		CONSOLEWRITENationListPicker()
		break
	case FinanceView:
		VIEWShowCountryFinances(pickedNation)
		break
	case MilitaryView:
		VIEWShowCountryMilitary(pickedNation)
		break
	case ProductionView:
		VIEWWRITECountryProduction(pickedNation)
		break
	case LawView:
		VIEWWRITECountryLaws(pickedNation)
		break
	case TechView:
		VIEWWRITECountryTechs(pickedNation)
		break
	}
}

func VIEWShowDEFAULTVIEW(pickedNation Country) {
	fmt.Println("____________________________")
	fmt.Print(pickedNation.Name, "||")
	fmt.Print("Money: ", pickedNation.Money)
	fmt.Print("||Balance: ", pickedNation.Income-pickedNation.Expenses)
	fmt.Print("||Pops: ", pickedNation.Population)
	fmt.Println("____________________________")
	fmt.Print("type 'econ' or 'economy' to go to the economics screen")
	fmt.Print("type 'save' to save your game")
	fmt.Print("type 'diplo' or 'diplomacy' to go to the diplomacy screen")
	fmt.Print("type 'exit' to exit")
}
func CONSOLEWRITENationListPicker() {
	fmt.Print(readFromFile("nationListings.json"))
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	if //noinspection SpellCheckingInspection
	strings.Contains("KOREACHINAJAPANTEUTONICPAPALAUSTRIABYZANTIUMNAPLESOTTOMANSHUNGARYPLCPOLANDLITHUANIAMUSCOVYBRANDENBURGFRANCEENGLANDSPAINSCOTLAND", text) == true {
		gameReadyStart(text)
	}
}
