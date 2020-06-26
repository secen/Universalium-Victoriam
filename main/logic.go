package main

import (
	"fmt"
)

func gameReadyStart(nation string) {
	_, _ = initAll()
	fmt.Println(nation)
}





func handleMilitaryInput(pickedNation Country) {

}


func executeLogic(globalEconomicsQueue queue) {
	EXECUPDatePlayerActions()
	EXECUPDateAIActions()
	runNextStepInSimulation(globalEconomicsQueue)
}

func runNextStepInSimulation(globalEconomicsQueue queue) {
	EXECUPDateEconomics(globalEconomicsQueue)
	//executeUpdateCountries();
	EXECUPDateTroops()
}

func executeUpdateCountries() {

}
func EXECUPDateEconomics(economicTaskQueue queue) {
	var f = func(){}
	if len(economicTaskQueue)>0{
		f,economicTaskQueue = popQueue(economicTaskQueue)
		f()
	}
}
func EXECUPDateTroops() {

}

func EXECUPDateAIActions() {

}


func getGameData(pickedNation Country,globalGoods []Good, globalNations []Country) GameData{
	var aux =  GameData{ pickedNation:pickedNation, globalGoods:globalGoods,globalNations:globalNations}
	return aux;
}
func EXECUPDatePlayerActions() {

}
func Tick(currentView VIEW, isFirstTime bool, pickedNation Country, globalEconomicsQueue queue, controller Controller) {
	executeLogic(globalEconomicsQueue)
}
