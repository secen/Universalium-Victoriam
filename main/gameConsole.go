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
	for{
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		inputHandlerGameConsole(text)
	}
}

var pickedNation = Country{}


func inputHandlerGameConsole(text string) {

	if strings.Compare("exit",text) == 0{
		os.Exit(0)
	}
	if strings.Compare("clear",text) == 0{
		CallClear()
	}
	if strings.Compare("help",text) == 0{
		VIEWShowNationPickerHelpData()
	}
	if strings.Compare("pick",text) == 0 {
		CONSOLEWRITENationListPicker()
	}
	if strings.Compare("back",text)==0{
		gameState = MainMenu
		main()
	}
}

//noinspection GoBoolExpressions
func startGame(){
	var gameViewController = initGameView()
	var globalEconomicsQueue = queue{}
	var isGameRunning = true
	var currentView = new(VIEW)
	var globalCountries = initCountries();
	var globalGoods = initGoods();
	*currentView = DefaultView
	for{
		if !isGameRunning{
			break
		}
		CONSOLEWRITEOverview(currentView,pickedNation,globalEconomicsQueue)
		CONSOLEHandlePlayerInput(currentView,&pickedNation,globalEconomicsQueue, gameViewController,globalGoods,globalCountries)
		Tick(*currentView, true, pickedNation,globalEconomicsQueue, gameViewController)

	}
}

func CONSOLEHandlePlayerInput(view *VIEW, pickedNation *Country, economicsQueue queue, controller Controller, globalGoods []Good, globalCountries []Country) {
	reader := bufio.NewReader(os.Stdin)
	switch *view {
	case DefaultView:
		var commands = map[string]func(){
			"econ":func(){*view = FinanceView},
			"economy": func(){*view = FinanceView},
			"exit": func(){os.Exit(0)},
			"save" :func(){saveGame(getGameData(*pickedNation,globalGoods,globalCountries))},
			"diplo" :func(){},//TODO: IMPLEMENT THE DIPLOMACY SCREEN
			"diplomacy": func(){},//TODO: IMPLEMENT THE DIPLOMACY SCREEN
		}
		text,_ := reader.ReadString('\n')
		f,ok:= commands[text];
		if ok {
			f();
		}
		break;
	case FinanceView:
		var commands = map[string]func(){
		"lower tax": func(){*pickedNation=ECONLowerTax(*pickedNation)},
		"raise tax":func(){*pickedNation=ECONRaiseTax(*pickedNation)},
		"back":func(){*view=DefaultView},
		}
		text,_ := reader.ReadString('\n')
		f,ok:= commands[text];
		if ok {
			f();
		}
		break;
	case LawView:
		break;
	case ProductionView:
		break;
	case MilitaryView:
		break;
	case TechView:
		break;
	case NationPicker:
		break;
	}
}