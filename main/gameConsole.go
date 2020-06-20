package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func showHelpGameConsoleNationPicker() {
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
func nationPicker() {
	fmt.Println("N A T I O N   P I C K E R!!!!!")
	showHelpGameConsoleNationPicker()
	for{
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		inputHandlerGameConsole(text)
	}
}
var pickedNation = country{}

func nationListPicker() {
	fmt.Println("SCOTLAND - FOREVER HIGHLANDER!!!!")
	fmt.Println("ENGLAND - The Troubled Years")
	fmt.Println("FRANCE - The Big Blue")
	fmt.Println("BRANDENBURG - German Military Might")
	fmt.Println("MUSCOVY - Aspiring Hegemon")
	fmt.Println("PLC - Polish-Lithuanian Commonwealth - The Eastern European Hegemon")
	fmt.Println("HUNGARY - Crysis in The Carpatheans")
	fmt.Println("THE OTTOMANS - Aspiring Islamic Powerhouse")
	fmt.Println("BYZANTIUM - The Last Western Bastion in Asia Minor")
	fmt.Println("AUSTRIA - The Habsburgic Monarchy")
	fmt.Println("SPAIN - Internal Strife and Wealth Of New Worlds")
	fmt.Println("TWO NAPLES - Italian Struggles")
	fmt.Println("THE PAPAL STATES - The Only Godly Kingdom")
	fmt.Println("THE TEUTONIC ORDER - PURGE THY HERESY!")
	fmt.Println("JAPAN - The 1000 Years Shogunate")
	fmt.Println("CHINA - The Ming Paper Tiger")
	fmt.Println("KOREA - The Hermit Kingdom")
	fmt.Println("Please Enter your pick: ")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	if strings.Contains("KOREACHINAJAPANTEUTONICPAPALAUSTRIABYZANTIUMNAPLESOTTOMANSHUNGARYPLCPOLANDLITHUANIAMUSCOVYBRANDENBURGFRANCEENGLANDSPAINSCOTLAND" , text) == true {
		gameReadyStart(text)
	}
}
func inputHandlerGameConsole(text string) {

	if strings.Compare("exit",text) == 0{
		os.Exit(0)
	}
	if strings.Compare("clear",text) == 0{
		CallClear()
	}
	if strings.Compare("help",text) == 0{
		showHelpGameConsoleNationPicker()
	}
	if strings.Compare("pick",text) == 0 {
		nationListPicker()
	}
	if strings.Compare("back",text)==0{
		gameState = MainMenu
		main()
	}
}

//noinspection GoBoolExpressions
func startGame(){
	var isGameRunning = true
	for{
		if !isGameRunning{
			break
		}
		Tick(DefaultView,true, pickedNation)
	}
}