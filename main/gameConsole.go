package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func nationPicker() {
	fmt.Println("N A T I O N   P I C K E R!!!!!")
	VIEWShowNationPickerHelpData()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		inputHandlerGameConsole(text)
	}
}

func inputHandlerGameConsole(text string) {

	var validInputs = map[string]func(){
		"exit":  func() { os.Exit(0) },
		"clear": CallClear,
		"help":  VIEWShowNationPickerHelpData,
		"pick":  CONSOLEWRITENationListPicker,
		"back": func() {
			gameState = MainMenu
			main()
		},
	}
	result, ok := validInputs[text]
	if ok {
		result()
	}
}

//noinspection GoBoolExpressions
func startGame() {
	var gameViewController = initGameView()
	var globalEconomicsQueue = queue{}
	var isGameRunning = true
	var currentView = new(VIEW)
	var globalCountries = initCountries()
	var globalGoods = initGoods()
	var pickedNation Country
	*currentView = DefaultView
	for {
		if !isGameRunning {
			break
		}
		CONSOLEWRITEOverview(currentView, pickedNation, globalEconomicsQueue)
		CONSOLEHandlePlayerInput(currentView, &pickedNation, globalEconomicsQueue, gameViewController, globalGoods, globalCountries)
		Tick(*currentView, true, pickedNation, globalEconomicsQueue, gameViewController)

	}
}

func CONSOLEHandlePlayerInput(view *VIEW, pickedNation *Country, economicsQueue queue, controller Controller, globalGoods []Good, globalCountries []Country) {
	reader := bufio.NewReader(os.Stdin)
	switch *view {
	case DefaultView:
		var commands = map[string]func(){
			"econ":      func() { *view = FinanceView },
			"economy":   func() { *view = FinanceView },
			"exit":      func() { os.Exit(0) },
			"save":      func() { saveGame(getGameData(*pickedNation, globalGoods, globalCountries)) },
			"diplo":     func() { *view = DiplomacyView }, //TODO: IMPLEMENT THE DIPLOMACY SCREEN
			"diplomacy": func() { *view = DiplomacyView }, //TODO: IMPLEMENT THE DIPLOMACY SCREEN
		}
		text, _ := reader.ReadString('\n')
		f, ok := commands[text]
		if ok {
			f()
		}
		break
	case FinanceView:
		var commands = map[string]func(){
			"lower tax": func() { *pickedNation = ECONLowerTax(*pickedNation) },
			"raise tax": func() { *pickedNation = ECONRaiseTax(*pickedNation) },
			"back":      func() { *view = DefaultView },
		}
		text, _ := reader.ReadString('\n')
		f, ok := commands[text]
		if ok {
			f()
		}
		break
	case LawView:
		break
	case ProductionView:
		break
	case MilitaryView:
		break
	case TechView:
		break
	case NationPickerView:
		break
	}
}
