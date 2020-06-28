package main

import "fmt"

type VIEW int

const (
	DefaultView    VIEW = 0
	FinanceView    VIEW = 1
	LawView        VIEW = 2
	ProductionView VIEW = 3
	MilitaryView   VIEW = 4
	TechView       VIEW = 5
	NationPicker VIEW = 6
)

func VIEWShowNationPickerHelpData() {
	fmt.Println(readFromFile(nationPickerFilename))
}
func VIEWShowCountryFinances(pickedNation Country) {
	fmt.Println("Treasury: ", pickedNation.Money)
	fmt.Println("Debt: ", pickedNation.Debt)
	fmt.Println("Income: ", pickedNation.Income)
	fmt.Println("Expenses: ", pickedNation.Expenses)
	fmt.Println("Interest: ", pickedNation.interest)
	fmt.Println("Balance: ", int32(pickedNation.Income)-int32(pickedNation.Expenses))
	fmt.Println("To raise taxes please type 'raise tax'")
	fmt.Println("To lower taxes please type 'lower tax'")
	fmt.Println("To go back to the default view, please type 'back'")
}
func VIEWDiplomacyOverview(pickedNation Country){
	fmt.Println("Please select one of the following countries to have diplomatic interactions with: ")
	var countries = parseJSONToCountryListings(readFromFile("nationListings.json"))
	for i, country:= range countries {
		fmt.Println(i,".",country.CountryName);
	}
}
func VIEWDiplomaticActions(pickedNation Country, countryToInfluence Country){
	fmt.Println("type 'declare war' to declare war");
	fmt.Println("type 'claim' to make a claim");
	fmt.Println("type 'renounce claim' to renounce a claim");
	fmt.Println("type: 'offer alliance' to offer an alliance");
	fmt.Println("type: 'break alliance' to break an alliance");
	fmt.Println("type: 'show values' to show the selected country's relation values");
	fmt.Println("type 'insult' to decrease your relations with the chosen country");
	fmt.Println("type 'improve' to improve your relations with the chosen country");
	fmt.Println("type 'diplomat' to access your diplomat")
	fmt.Println("type 'spy' to access your spy")
	fmt.Println("Chosen country: ", countryToInfluence.Name);
}
func VIEWWRITECountryTechs(pickedNation Country) {
	//writes the Technology interface on the console
	var str = readFromFile("techs.json");
	var techArr []Technology = fromJSONToTechArr(str);
	for i:=0;i<len(techArr);i++{
		var tech = techArr[i];
		for j:=0;j<len(pickedNation.techs);j++{
			if tech == pickedNation.techs[j] {
				if pickedNation.techs[j].Taken {
					fmt.Print("[TAKEN]||")
				} else {
					fmt.Print("[     ]||")
				}
				fmt.Print(pickedNation.techs[j].Name)
				fmt.Println("||"+pickedNation.techs[j].Category)
			}
		}
	}
}

func VIEWShowCountryMilitary(pickedNation Country) {
	//writes the military interface on the console
	//TODO: IMPLEMENT MILITARY INTERFACE
}

func VIEWWRITECountryProduction(pickedNation Country) {
	//writes the production interface on the console
	//TODO: IMPLEMENT PRODUCTION INTERFACE
}

func VIEWWRITECountryLaws(pickedNation Country) {
	//writes the law interface on the console
	fmt.Println("Laws Enacted:")
	for i, s := range pickedNation.laws {
		fmt.Println(i, s.name)
	}
	fmt.Println("To propose new legistlation please type 'propose'")
	fmt.Println("To abolish legislation" +
		" please type 'abolish' and the number of the law")
	fmt.Println("-> ")
}

func VIEWShowCountryIntroductionPopup(pickedNation Country) {
	fmt.Println("The nation of ", countryCodes[(int)(pickedNation.Code)])
	showCountryStats(pickedNation)
	fmt.Println("May you be successful in your endevours!")
}
