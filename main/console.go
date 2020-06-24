package main

import (
	"fmt"
	"os"
	"strings"
)
func inputHandler(text string) {

	if strings.Compare("exit",text) == 0{
		os.Exit(0)
	}
	if strings.Compare("clear",text) == 0{
		CallClear()
	}
	if strings.Compare("help",text) == 0{
		showCommands()
	}
	if strings.Compare("more debug",text) == 0{
		showDebugCommands()
	}
	if strings.Compare("start",text) == 0 {
		startGame()
	}
	if strings.Compare("options",text)==0{
		startOptions()
	}
	if strings.Compare("debug show Country",text) == 0 {
		var debugCountry = Country{
		CultureScore:        2,
		MilitaryScore:       20,
		InfrastructureScore: 10,
		Debt:                500,
		Money:               300,
		Code:                12,
		Gdp:                 100,
			}
		showCountryStats(debugCountry)
	}
	if strings.Compare("debug tests",text)==0{
		execTests()
	}
}
func showCountryStats(cnt Country) {
	fmt.Println(cnt.Code)
	fmt.Println("Infrastrucutal Score: ",cnt.InfrastructureScore)
	fmt.Println("Military Score: ",cnt.MilitaryScore)
	fmt.Println("Cultural Score:", cnt.CultureScore)
	fmt.Println("Debt: ",cnt.Debt,)
	fmt.Println("Money: ",cnt.Money)
	fmt.Println("GDP: ", cnt.Gdp)
}
func showCommands(){
	fmt.Println("exit - exits the game")
	fmt.Println("start - starts the game")
	fmt.Println("clear - clears the screen")
	fmt.Println("options - shows the options screen")
	fmt.Println("more debug - shows debug options")
}
func showDebugCommands() {
	fmt.Println("debug show Country - shows a debug Country overview")
	fmt.Println("debug tests - tests")
}