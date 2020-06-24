package main
func ECONDebaseCurrency (pickedNation country) country {
	//[TESTED]
	pickedNation.income*=5
	pickedNation.gdp/=2
	return pickedNation
}

func ECONNationalizeIndustries(pickedNation country) country{
	//[TESTED]
	pickedNation.income/=2
	pickedNation.gdp/=2
	pickedNation.debt= 0
	pickedNation.infrastructureScore*=5
	pickedNation.interest=0
	pickedNation.money*=30
	return pickedNation
}

func ECONRaiseTax(pickedNation country) country{
	//[TESTED]
		pickedNation.income += 5
	pickedNation.gdp -= 5
	return pickedNation
}

func ECONLowerTax(pickedNation country) country {
	//[TESTED]
	pickedNation.income-=5
	pickedNation.gdp+=5
	return pickedNation
}