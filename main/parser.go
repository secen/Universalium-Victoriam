package main

import (
	"encoding/json"
)
type GameResult View
type ParserResult View
func parseGameData(gameView View) View {
	return View{}
}
func parseInterfaceToJSON(goodToBeParsed interface{}) string{
	/*var goodDemand =strconv.FormatFloat(goodToBeParsed.demand,'f',-1,64)
	var goodSupply =strconv.FormatFloat(goodToBeParsed.supply,'f',-1,64)
	var goodPrice =strconv.FormatFloat(goodToBeParsed.price,'f',-1,64)
	var goodVolatility =strconv.FormatFloat(goodToBeParsed.volatility,'f',-1,64)
	var goodVolatilityOfRateOfChange = strconv.FormatFloat(goodToBeParsed.volatilityOfRateOfChange,'f',-1,64)
	var goodRateOfChange =strconv.FormatFloat(goodToBeParsed.rateOfChange,'f',-1,64)
	*//*var json = "{\n" +
		"Name: " + goodToBeParsed.Name+
		"\"demand\": " + goodDemand + ",\n"+
		"\"supply\": " + goodDemand + ",\n"+
		"\"price\": " + goodDemand + ",\n"+
		"\"volatility\": " + goodDemand + ",\n"+
		"\"volatilityOfRateOfChange\": " + goodDemand + ",\n"+
		"\"demand\": " + goodDemand + ",\n"+
		"}"
	*/
	var aux = &goodToBeParsed;
	b,_ := json.Marshal(aux)
	return string(b);
}
func parseCountryFromJSON(JSONStr string) Country{
	var data = []byte(JSONStr)
	var str Country;
	_ = json.Unmarshal(data,&str)
	return str;

}
func parseJSONToGood(JSONStr string) Good{
	var data = []byte(JSONStr)
	var str Good;
	_ = json.Unmarshal(data,&str)
	return str;
}


func parseJSONToCountryListings(JSONStr string) []CountryListing{
	var data = []byte(JSONStr)
	var str []CountryListing;
	_ = json.Unmarshal(data,&str)
	return str;
}

func parseJSONToListing(JSONStr string) CountryListing{
	var data = []byte(JSONStr)
	var str CountryListing;
	_ = json.Unmarshal(data,&str)
	return str;
}