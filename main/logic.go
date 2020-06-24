package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func gameReadyStart(nation string) {
	_, _ = initAll()
	fmt.Println(nation)
}

func writeOutput(currentView VIEW, isFirstTime bool, pickedNation country, globalEconomicsQueue queue) {
	if isFirstTime && currentView!=NationPicker {
		VIEWShowCountryIntroductionPopup(pickedNation)
		isFirstTime = false
	}
	if isFirstTime {
		VIEWShowCountryIntroductionPopup(pickedNation)
		CONSOLEWRITEOverview(&currentView, pickedNation,globalEconomicsQueue)
		isFirstTime=false
	}else{
		CONSOLEWRITEOverview(&currentView, pickedNation,globalEconomicsQueue)
	}
}


func handleFinancesInput(pickedNation country, economicQueue queue){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	if strings.Compare("raise tax",text)==0 {
		appendQueue(economicQueue, func() {pickedNation=ECONRaiseTax(pickedNation) })
	}
	if strings.Compare("lower tax",text)==0{
		appendQueue(economicQueue, func() {
			pickedNation=ECONLowerTax(pickedNation)
		})
	}
	if (containsTech(pickedNation.techs,technology{ID:2}) == true){
		if strings.Compare("debase",text)==0{
			appendQueue(economicQueue, func(){pickedNation=ECONDebaseCurrency(pickedNation)})
		}
	}
	if (containsTech(pickedNation.techs,technology{ID:14}) == true){
		if strings.Compare("nationalize",text)==0{
			appendQueue(economicQueue,func(){
				pickedNation=ECONNationalizeIndustries(pickedNation)
			})
		}
	}
}

func handleMilitaryInput(pickedNation country) {

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

func EXECUPDatePlayerActions() {

}
func handleGameInput() {
	fmt.Println("type help")
}
func Tick(currentView VIEW, isFirstTime bool, pickedNation country, globalEconomicsQueue queue, controller Controller) {
	writeOutput(currentView, isFirstTime, pickedNation, globalEconomicsQueue)
	executeLogic(globalEconomicsQueue)
	handleGameInput()
}
