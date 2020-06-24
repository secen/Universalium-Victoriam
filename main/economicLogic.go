package main
func ECONDebaseCurrency (pickedNation Country) Country {
	//[TESTED]
	pickedNation.Income *=5
	pickedNation.Gdp /=2
	return pickedNation
}

func ECONNationalizeIndustries(pickedNation Country) Country{
	//[TESTED]
	pickedNation.Income /=2
	pickedNation.Gdp /=2
	pickedNation.Debt = 0
	pickedNation.InfrastructureScore *=5
	pickedNation.interest=0
	pickedNation.Money *=30
	return pickedNation
}

func ECONRaiseTax(pickedNation Country) Country{
	//[TESTED]
		pickedNation.Income += 5
	pickedNation.Gdp -= 5
	return pickedNation
}

func ECONLowerTax(pickedNation Country) Country {
	//[TESTED]
	pickedNation.Income -=5
	pickedNation.Gdp +=5
	return pickedNation
}