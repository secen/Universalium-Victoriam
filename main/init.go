package main

var countryCodes map[int]string
func initCountries () []country{
	var auxvec []country
	auxvec = make([]country,10)
	var auxcnt = country{
		code : 1,
		money: 30,
		cultureScore:100,
		infrastructureScore:20,
		gdp:30,
		militaryScore:20,
		debt:0,
		population:100,
	}
	countryCodes = make(map[int]string)
	countryCodes[1] = "Austria"
	auxvec[0] = auxcnt
	return auxvec
}
func initCountryRelations (countries []country) []RelationEntry {
	var relations = make([]RelationEntry,10)
	relations[0] = RelationEntry{cnt1:countries[0],cnt2:countries[1],rel:PEACE}
	return relations
}

//noinspection ALL
func initGameView() Controller {
var gameView View
	var gameOutputView View
	var (
	gameController = *NewController(gameView,parseGameData,func(cntrl Controller)Controller{cntrl.output= parseGameData(gameView);return cntrl },gameOutputView))
gameController = gameController.run(gameController)
return gameController
}
func initAll() ([]country,[]RelationEntry){
	print("Loading assets....")
	var countries = make([]country,10)
	countries = initCountries()
	var relations = initCountryRelations(countries)
	return countries, relations
}
