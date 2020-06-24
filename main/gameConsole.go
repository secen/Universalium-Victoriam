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

var pickedNation = country{}


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
	*currentView = DefaultView
	for{
		if !isGameRunning{
			break
		}
		CONSOLEWRITEOverview(currentView,pickedNation,globalEconomicsQueue)
		Tick(*currentView, true, pickedNation,globalEconomicsQueue, gameViewController)

	}
}