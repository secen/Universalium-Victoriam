package main

type econTags uint32

const (
	INFLATIONPRESSURE = 12
	DEFLATIONPRESSURE = 13
	NATIONALIZED      = 14
	PRIVATIZED        = 15
	BUYINGASSETS      = 16
	SELLINGASSETS     = 17
	DEBASED           = 18
	RAISINGTAX        = 19
	LOWERINGTAX       = 20
	BONDSELLING       = 21
)

func calculateEconomicStability(pickedNation Country) int32 {
	var score int32 = 0
	for _, item := range pickedNation.NationalEffects {
		if item.ID == INFLATIONPRESSURE || item.ID == PRIVATIZED || item.ID == DEBASED || item.ID == LOWERINGTAX {
			score += 1
		} else {
			score -= 1
		}
	}
	return score
}
func ECONDebaseCurrency(pickedNation Country) Country {
	//[TESTED]
	pickedNation.Income *= 5
	pickedNation.Gdp /= 2
	ECONAddTag(pickedNation, DEBASED)
	return pickedNation
}
func ECONPrivatize(pickedNation Country) Country {
	var privatization = NationalEffect{
		ID:          3,
		Name:        "Privatization",
		Description: "We have sold off some of our key assets to private enterprises",
		ConditionID: 3,
		EffectID:    3,
		effect: func(country Country) Country {
			country.Gdp *= 4
			country.Money *= 40
			if country.Money == 0 {
				country.Money += 2000
			}
			country.Income /= 20
			return country
		},
		DurationInMonths: 4 * 12,
	}
	pickedNation = addNationalEffect(privatization, pickedNation)
	pickedNation = ECONAddTag(pickedNation, PRIVATIZED)
	pickedNation = applyNationalEffect(privatization, pickedNation)
	return pickedNation
}
func ECONAssetSelling(pickedNation Country) Country {
	var boughtAssets = NationalEffect{
		ID:          4,
		Name:        "Asset Selling",
		Description: "We are selling key assets that we own to the highest bidder",
		ConditionID: 4,
		EffectID:    4,
		effect: func(country Country) Country {
			country.Income /= 10
			country.Debt /= 2
			country.Money *= 4
			return country
		},
		DurationInMonths: 4,
	}
	pickedNation = ECONAddTag(pickedNation, SELLINGASSETS)
	pickedNation.NationalEffects = append(pickedNation.NationalEffects, boughtAssets)
	pickedNation = pickedNation.NationalEffects[0].effect(pickedNation)
	return pickedNation
}
func ECONAddTag(pickedNation Country, tag uint32) Country {

	pickedNation.NationalEffects = append(pickedNation.NationalEffects, NationalEffect{
		ID:               tag,
		Name:             "",
		Description:      "",
		ConditionID:      0,
		EffectID:         0,
		effect:           nil,
		DurationInMonths: 0,
	})
	return pickedNation
}
func ECONAssetBuying(pickedNation Country) Country {
	var boughtAssets = NationalEffect{
		ID:          4,
		Name:        "Asset Buying",
		Description: "We are buying assets tha have performed well in the past in order for our income to rise",
		ConditionID: 4,
		EffectID:    4,
		effect: func(country Country) Country {
			country.Income *= 10
			country.Debt *= 2
			country.Money /= 4
			return country
		},
		DurationInMonths: 4,
	}
	pickedNation = addNationalEffect(boughtAssets, pickedNation)
	pickedNation = ECONAddTag(pickedNation, BUYINGASSETS)
	pickedNation = applyNationalEffect(boughtAssets, pickedNation)
	return pickedNation
}

func applyNationalEffect(effect NationalEffect, nation Country) Country {
	for _, item := range nation.NationalEffects {
		if item.ID == effect.ID {
			nation = item.effect(nation)
			return nation
		}
	}
	return nation
}
func addNationalEffect(effect NationalEffect, country Country) Country {
	if len(country.NationalEffects) == 0 {
		country.NationalEffects = make([]NationalEffect, 1)
		country.NationalEffects[0] = effect
	} else {
		country.NationalEffects = append(country.NationalEffects, effect)
	}
	return country
}
func ECONNationalizeIndustries(pickedNation Country) Country {
	//[TESTED]
	var nationalizeIndustry = NationalEffect{
		ID:          3,
		Name:        "Nationalization",
		Description: "The key assets of our state have gotten back into the state's hands from the hands of private enterprise",
		ConditionID: 3,
		EffectID:    3,
		effect: func(country Country) Country {
			country.Gdp /= 4
			country.Money *= 40
			if country.Money == 0 {
				country.Money += 2000
			}
			country.Income /= 20
			country.interest = 0
			country.Debt = 0
			return country
		},
		DurationInMonths: 4 * 12,
	}
	pickedNation = addNationalEffect(nationalizeIndustry, pickedNation)
	pickedNation = ECONAddTag(pickedNation, NATIONALIZED)
	pickedNation = applyNationalEffect(nationalizeIndustry, pickedNation)
	return pickedNation
}
func ECONInflate(pickedNation Country) Country {
	var inflatedSupply = NationalEffect{Name: "Inflated Money Supply", Description: "Due to the actions of our central bank, the money supply has been inflated, this has resulted in higher interest rates due to our inflationary pressure.", effect: func(country Country) Country {
		country.interest *= 70
		if country.Debt == 0 {
			country.Debt = 5
		}
		country.Debt *= 20
		country.Income /= 4
		country.Gdp *= 2
		country.Money *= 60
		return country
	}, DurationInMonths: 3}
	pickedNation = addNationalEffect(inflatedSupply, pickedNation)
	pickedNation = ECONAddTag(pickedNation, BUYINGASSETS)
	pickedNation = applyNationalEffect(inflatedSupply, pickedNation)
	return pickedNation
}
func ECONDeflate(pickedNation Country) Country {
	var inflatedSupply = NationalEffect{Name: "Deflated Money Supply", Description: "Due to the actions of our central bank, the money supply has been deflated, this has resulted in lower interest rates due to our deflationary pressure.", effect: func(country Country) Country {
		country.interest /= 70
		if country.Debt > 0 {
			country.Debt /= 20
		}
		country.Income *= 4
		country.Gdp /= 2
		country.Money /= 60
		return country
	}, DurationInMonths: 3}
	pickedNation = addNationalEffect(inflatedSupply, pickedNation)
	pickedNation = ECONAddTag(pickedNation, DEFLATIONPRESSURE)
	pickedNation = applyNationalEffect(inflatedSupply, pickedNation)
	return pickedNation
}
func ECONRaiseTax(pickedNation Country) Country {
	//[TESTED]
	pickedNation.Income += 5
	pickedNation.Gdp -= 5
	pickedNation = ECONAddTag(pickedNation, RAISINGTAX)
	return pickedNation
}
func ECONBondSelling(pickedNation Country) Country {
	var boughtAssets = NationalEffect{
		ID:          8,
		Name:        "Selling Bonds",
		Description: "We are selling bonds to private investors",
		ConditionID: 8,
		EffectID:    8,
		effect: func(country Country) Country {
			country.Gdp *= 2
			country.Debt *= 2
			country.Income *= 2
			return country
		},
		DurationInMonths: 12,
	}
	pickedNation = addNationalEffect(boughtAssets, pickedNation)
	pickedNation = ECONAddTag(pickedNation, BONDSELLING)
	pickedNation = applyNationalEffect(boughtAssets, pickedNation)
	return pickedNation
}
func ECONLowerTax(pickedNation Country) Country {
	//[TESTED]
	pickedNation.Income -= 5
	pickedNation.Gdp += 5
	pickedNation = ECONAddTag(pickedNation, LOWERINGTAX)

	return pickedNation
}
