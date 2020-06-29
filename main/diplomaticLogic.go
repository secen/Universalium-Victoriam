package main

import "strconv"

func execDiplomaticActions(pickedNation Country, countryToInfluence Country) {
	var (
		commands = map[string]func(Country, Country) (Country, Country, string){
			"declare war":    DIPLODeclareWarOnCountry,
			"claim":          DIPLOClaimInterface,
			"renounce claim": DIPLORenounceClaimInterface,
			"offer alliance": DIPLOOfferAlliance,
			"break alliance": DIPLOBreakAlliance,
			"show values":    DIPLOShowCountryValues,
			"insult":         DIPLOInsult,
			"improve":        DIPLOImproveRelations,
			"diplomat":       DIPLOViewDiplomat,
			"spy":            DIPLOViewSpy,
		}
	)
	commands["declare war"](pickedNation, countryToInfluence)
}

func DIPLOClaimInterface(country4 Country, country2 Country) (country Country, country3 Country, s string) {
	return country4, country2, "Claim Claimed"
}

func DIPLORenounceClaimInterface(country4 Country, country2 Country) (country Country, country3 Country, s string) {
	return country4, country2, "Claim Renounced"
}

func DIPLOViewSpy(pickedNation Country, country2 Country) (country Country, country3 Country, s string) {
	return pickedNation, country2, "Spy Viewed"
}

func DIPLOViewDiplomat(pickedNation Country, country2 Country) (country Country, country3 Country, s string) {
	return pickedNation, country2, "Diplomat Viewed"
}

func DIPLOShowCountryValues(pickedNation Country, to Country) (country Country, country2 Country, str string) {

	aux := ""
	b := []byte("")
	for _, val := range pickedNation.relations {
		b = strconv.AppendInt(b, int64(val.Opinionvalue), 10)
		aux += val.Cnt2.Name + ":" + string(b)
		switch val.Rel {
		case PEACE:
			aux += ": at peace\n"
			break
		case WAR:
			aux += ": at war\n"
			break
		case CORDIAL:
			aux += ": in cordial relations\n"
			break
		case THREATENED:
			aux += ": feels threatened by us\n"
			break
		case HISTORICAL:
			aux += ": historical relationship\n"
			break
		case ALLIED:
			aux += ": in alliance\n"
			break
		}
	}
	return pickedNation, to, aux
}

func DIPLOBreakAlliance(nation Country, to Country) (country Country, country3 Country, str string) {
	nation.relations = append(nation.relations, RelationEntry{
		Cnt1:         nation,
		Cnt2:         to,
		Rel:          THREATENED,
		Opinionvalue: 150,
	})
	to.relations = append(to.relations, RelationEntry{
		Cnt1:         to,
		Cnt2:         nation,
		Rel:          THREATENED,
		Opinionvalue: 150,
	})
	return nation, to, "Alliance Broken"
}

func DIPLOImproveRelations(pickedCountry Country, countryToDeclareOn Country) (country Country, country3 Country, str string) {
	pickedCountry.relations = append(pickedCountry.relations, RelationEntry{pickedCountry, countryToDeclareOn, CORDIAL, 20})
	countryToDeclareOn.relations = append(countryToDeclareOn.relations, RelationEntry{pickedCountry, countryToDeclareOn, CORDIAL, 20})
	return pickedCountry, countryToDeclareOn, "Relations Improved"
}

func DIPLOOfferAlliance(pickedNation Country, countryToAllyTo Country) (country Country, country3 Country, str string) {
	if isAllianceOfInterest(pickedNation, countryToAllyTo) {
		return DIPLOCreateAlliance(pickedNation, countryToAllyTo)
	} else {
		return pickedNation, countryToAllyTo, "Alliance Not Made"
	}
}

func DIPLOCreateAlliance(nation Country, to Country) (Country, Country, string) {
	nation.relations = append(nation.relations, RelationEntry{
		Cnt1:         nation,
		Cnt2:         to,
		Rel:          ALLIED,
		Opinionvalue: 150,
	})
	to.relations = append(to.relations, RelationEntry{
		Cnt1:         to,
		Cnt2:         nation,
		Rel:          ALLIED,
		Opinionvalue: 150,
	})
	return nation, to, "Alliance Created"
}
func isAllianceOfInterest(nation Country, to Country) bool {
	if DIPLOgetRelations(nation, to) > 100 {
		return true
	}
	return false
}

func DIPLOgetRelations(nation Country, to Country) int32 {
	var relationSum int32 = 0
	for _, rel := range nation.relations {
		if rel.Cnt1.Code == nation.Code && rel.Cnt2.Code == to.Code {
			relationSum += rel.Opinionvalue
		}
	}
	return relationSum
}

func DIPLOInsult(pickedCountry Country, countryToDeclareOn Country) (country Country, country3 Country, str string) {
	pickedCountry.relations = append(pickedCountry.relations, RelationEntry{pickedCountry, countryToDeclareOn, THREATENED, -20})
	countryToDeclareOn.relations = append(countryToDeclareOn.relations, RelationEntry{pickedCountry, countryToDeclareOn, THREATENED, -20})
	return pickedCountry, countryToDeclareOn, "Country Insulted"
}

func DIPLODeclareWarOnCountry(pickedCountry Country, countryToDeclareOn Country) (country Country, country3 Country, str string) {
	pickedNation.relations = append(pickedNation.relations, RelationEntry{pickedNation, countryToDeclareOn, WAR, -200})
	countryToDeclareOn.relations = append(countryToDeclareOn.relations, RelationEntry{pickedNation, countryToDeclareOn, WAR, -200})
	return pickedNation, countryToDeclareOn, "War Declared"
}
