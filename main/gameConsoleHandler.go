package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CONSOLEWRITEOverview(currentView *VIEW, pickedNation Country,globalEconomicQueue queue) {
	switch *currentView {
	case DefaultView:
		*currentView= NationPicker
		break
	case NationPicker:
		CONSOLEWRITENationListPicker()
		break
	case FinanceView:
		VIEWShowCountryFinances(pickedNation)
		handleFinancesInput(pickedNation, globalEconomicQueue)
		break
	case MilitaryView:
		VIEWShowCountryMilitary(pickedNation)
		handleMilitaryInput(pickedNation)
		break
	case PorductionView:
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
func CONSOLEWRITENationListPicker() {
	fmt.Print(readFromFile("nationListings.json"))
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	if //noinspection SpellCheckingInspection
	strings.Contains("KOREACHINAJAPANTEUTONICPAPALAUSTRIABYZANTIUMNAPLESOTTOMANSHUNGARYPLCPOLANDLITHUANIAMUSCOVYBRANDENBURGFRANCEENGLANDSPAINSCOTLAND" , text) == true {
		gameReadyStart(text)
	}
}