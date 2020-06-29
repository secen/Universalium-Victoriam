package main

var countryCodes map[int]string

func initCountries() []Country {
	var auxvec []Country
	auxvec = make([]Country, 10)
	var auxcnt = Country{
		Code:                1,
		Money:               30,
		CultureScore:        100,
		InfrastructureScore: 20,
		Gdp:                 30,
		MilitaryScore:       20,
		Debt:                0,
		Population:          100,
	}
	countryCodes = make(map[int]string)
	countryCodes[1] = "Austria"
	auxvec[0] = auxcnt
	return auxvec
}
func initGoods() []Good {
	return []Good{}
}
func initCountryRelations(countries []Country) []RelationEntry {
	var relations = make([]RelationEntry, 10)
	relations[0] = RelationEntry{Cnt1: countries[0], Cnt2: countries[1], Rel: PEACE}
	return relations
}

//noinspection ALL
func initGameView() Controller {
	var gameView View
	var gameOutputView View
	var (
		gameController = *NewController(gameView, parseGameData, func(cntrl Controller) Controller { cntrl.output = parseGameData(gameView); return cntrl }, gameOutputView)
	)
	gameController = gameController.run(gameController)
	return gameController
}
func initAll() ([]Country, []RelationEntry) {
	print("Loading assets....")
	var countries = make([]Country, 10)
	countries = initCountries()
	var relations = initCountryRelations(countries)
	return countries, relations
}
