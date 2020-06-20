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
	if strings.Compare("debug show country",text) == 0 {
		var debugCountry = country{
		cultureScore:        2,
		militaryScore:       20,
		infrastructureScore: 10,
		debt:                500,
		money:               300,
		code:                12,
		gdp:                 100,
			}
		showCountryStats(debugCountry)
	}
	if strings.Compare("debug tests",text)==0{
		tests()
	}
}
func showCountryStats(cnt country) {
	fmt.Println(cnt.code)
	fmt.Println("Infrastrucutal Score: ",cnt.infrastructureScore)
	fmt.Println("Military Score: ",cnt.militaryScore)
	fmt.Println("Cultural Score:", cnt.cultureScore)
	fmt.Println("Debt: ",cnt.debt,)
	fmt.Println("Money: ",cnt.money)
	fmt.Println("GDP: ", cnt.gdp)
}
func showCommands(){
	fmt.Println("exit - exits the game")
	fmt.Println("start - starts the game")
	fmt.Println("clear - clears the screen")
	fmt.Println("options - shows the options screen")
	fmt.Println("more debug - shows debug options")
}
func showDebugCommands() {
	fmt.Println("debug show country - shows a debug country overview")
	fmt.Println("debug tests - tests")
}