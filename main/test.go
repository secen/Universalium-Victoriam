package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

func TestingPathfinding() bool {
	var paths [][]bool = make([][]bool, 20)
	for i := range paths {
		paths[i] = make([]bool, 20)
	}
	paths = initDefaultPathVals(paths)
	paths = initPath(KOREA_SOUTH, KOREA_NORTH, paths)
	paths = initPath(KOREA_NORTH, JAPAN_OKINAWA, paths)
	visited := make([]bool, 20)
	var aux = make([]bool, 20)
	if KOREA_NORTH > KOREA_SOUTH {
		aux = searchForPath(KOREA_SOUTH, JAPAN_OKINAWA, paths, visited)
	} else {
		aux = searchForPath(KOREA_NORTH, JAPAN_OKINAWA, paths, visited)
	}
	aux2 := searchForPath(KOREA_NORTH, CHINA_JILIN, paths, visited) //aux2[CHINA_JILIN] should be false indicating that no path was found
	if aux[KOREA_SOUTH] && aux[JAPAN_OKINAWA] && aux2[CHINA_JILIN] == false {
		return true
	}
	return false
}

func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		_, _ = io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	_ = writer.Close()
	return <-out
}
func TestShowCountryFinancials() bool {
	var cnt Country = parseCountriesFromJSON(readFromFile(countriesFilename))[0]
	var str = captureOutput(func() { VIEWShowCountryFinances(cnt) })
	if strings.Contains(str, "Debt:  20") == true {
		return true
	} else {
		return false
	}
}
func testIOReadFromFile() bool {
	var str = readFromFile("nationListings.json")
	if strings.Contains(str, "KOREA") == true {
		return true
	} else {
		return false
	}
}
func testParsingOfListings() bool {
	var str = readFromFile("nationListings.json")
	var listings []CountryListing = parseJSONToCountryListings(str)
	if strings.Compare(listings[0].CountryName, "SCOTLAND") == 0 {
		return true
	}
	return false
}
func testTroopLoading() bool {
	var troops = parseJSONToTroops(readFromFile("gameData/debugTroops.json"))
	if troops[0].FirePower == 4 {
		return true
	}
	return false

}

//noinspection ALL
func testECONCountryFinancialModifyTax() bool {
	var cnt = Country{Gdp: 100,
		Debt:          20,
		Income:        40,
		Expenses:      50,
		Money:         120,
		Population:    300,
		MilitaryScore: 10,
		CultureScore:  100,
		interest:      0,
		Code:          0}
	var economicsqueue = queue{}
	economicsqueue = appendQueue(economicsqueue, func() { cnt.Income += 10 })
	EXECUPDateEconomics(economicsqueue)
	var str = captureOutput(func() { VIEWShowCountryFinances(cnt) })
	if strings.Contains(str, "Debt:  20") == true {
		return true
	} else {
		return false
	}
}
func testLAWAbolishLaw() bool {
	var laws2 = make([]Law, 10)
	laws2[0] = Law{Name: "Testing"}
	laws2[1] = Law{Name: "Testing2"}
	laws2[2] = Law{Name: "Mandate Of Heaven"}
	abolishLaw(1, laws2)
	if strings.Compare(laws2[1].Name, "Mandate Of Heaven") == 0 {
		return true
	} else {
		return false
	}
}
func testTechConsoleUI() bool {
	var cnt Country = parseCountriesFromJSON(readFromFile(countriesFilename))[0]
	var techArr []Technology = fromJSONToTechArr(readFromFile(techsFilename))
	cnt.Technologies = append(cnt.Technologies, techArr[0])
	var output = captureOutput(func() { VIEWWRITECountryTechs(cnt) })
	if strings.Contains(output, "Central Banking") {
		return true
	}
	return false
}
func testMap() bool {
	var mp string = getMap("gameData/debugArea.txt")
	if testMapHasWater(mp) &&
		testMapHasMountains(mp) &&
		testMapHasForests(mp) &&
		testEnemyExistsOnMap(mp) &&
		!testDeadExistOnMap(mp) {
		return true
	}
	return false
}

func testMapHasForests(mp string) bool {
	return mapHasForests(mp)
}

func testDeadExistOnMap(mp string) bool {
	return mapDeadExistOnMap(mp)
}

func testEnemyExistsOnMap(mp string) bool {
	return mapEnemyExists(mp)
}

func testMapHasMountains(mp string) bool {
	return mapHasMountains(mp)
}

func testMapHasWater(mp string) bool {
	return mapHasWater(mp)
}

func testIOLoadCountry() bool {
	var testCountry Country = loadCountryFromString(readFromFile(debugCountryFilename))
	if testCountry.Code == 12 {
		return true
	} else {
		return false
	}
}
func testDiplomaticActions() bool {
	var testCountry Country = parseCountriesFromJSON(readFromFile(countriesFilename))[0]
	var countryToDeclareOn Country = parseCountriesFromJSON(readFromFile(countriesFilename))[1]
	var tests []bool = make([]bool, 6)
	tests[0] = DIPLOgetRelations(testCountry, countryToDeclareOn) == 0
	testCountry, countryToDeclareOn, _ = DIPLOImproveRelations(testCountry, countryToDeclareOn)
	tests[1] = DIPLOgetRelations(testCountry, countryToDeclareOn) == 20
	testCountry, countryToDeclareOn, _ = DIPLOInsult(testCountry, countryToDeclareOn)
	tests[5] = DIPLOgetRelations(testCountry, countryToDeclareOn) == 0
	testCountry, countryToDeclareOn, _ = DIPLOOfferAlliance(testCountry, countryToDeclareOn)
	tests[2] = DIPLOHasAlliance(testCountry, countryToDeclareOn) == false
	testCountry, countryToDeclareOn, _ = DIPLOImproveRelations(testCountry, countryToDeclareOn)
	testCountry, countryToDeclareOn, _ = DIPLOImproveRelations(testCountry, countryToDeclareOn)
	testCountry, countryToDeclareOn, _ = DIPLOImproveRelations(testCountry, countryToDeclareOn)
	testCountry, countryToDeclareOn, _ = DIPLOImproveRelations(testCountry, countryToDeclareOn)
	testCountry, countryToDeclareOn, _ = DIPLOImproveRelations(testCountry, countryToDeclareOn)
	testCountry, countryToDeclareOn, _ = DIPLOImproveRelations(testCountry, countryToDeclareOn)
	testCountry, countryToDeclareOn, _ = DIPLOImproveRelations(testCountry, countryToDeclareOn)
	testCountry, countryToDeclareOn, _ = DIPLOImproveRelations(testCountry, countryToDeclareOn)
	testCountry, countryToDeclareOn, _ = DIPLOImproveRelations(testCountry, countryToDeclareOn)
	testCountry, countryToDeclareOn, _ = DIPLOImproveRelations(testCountry, countryToDeclareOn)
	testCountry, countryToDeclareOn, _ = DIPLOOfferAlliance(testCountry, countryToDeclareOn)
	tests[3] = DIPLOHasAlliance(testCountry, countryToDeclareOn)
	testCountry, countryToDeclareOn, _ = DIPLOBreakAlliance(testCountry, countryToDeclareOn)
	tests[4] = DIPLOHasAlliance(testCountry, countryToDeclareOn) == false
	DIPLORemoveALLRelations(testCountry, countryToDeclareOn)
	//TODO: Finish this function
	return checkIfAlltrue(tests)
}

func checkIfAlltrue(arr []bool) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == false {
			return false
		}
	}
	return true
}
func execTests() {
	fmt.Println("-------------------------------------------")
	fmt.Println("EXECUTING TESTS")
	fmt.Println("-------------------------------------------")
	var tests = []func() bool{TestShowCountryFinancials,
		testLAWAbolishLaw,
		testECONCountryFinancialModifyTax,
		testIOReadFromFile,
		testIOLoadCountry,
		testECONRaiseTaxes,
		testECONLowerTaxes,
		testECONDebase,
		testECONNationalize,
		testECONPrivatize,
		testsECONTick,
		testParsingOfListings,
		testTechConsoleUI,
		testMap,
		testTroopLoading,
		TestingPathfinding,
		testDiplomaticActions,
	}
	var hasPassed = make([]bool, len(tests))
	for i, testFunction := range tests {
		fmt.Println("Executing TEST NO.", i, " - ", GetFunctionName(testFunction))
		hasPassed[i] = testFunction()
		if hasPassed[i] {
			fmt.Println("Function", GetFunctionName(testFunction), " PASSED tests!")
		} else {
			fmt.Println("Function", GetFunctionName(testFunction), " FAILED tests!")
		}
	}
	fmt.Println("-------------------------------------------")
	fmt.Println("END RESULTS")
	fmt.Println("-------------------------------------------")
	var testsFailed = 0
	var testsSuccesfull = 0
	for i := 0; i < len(hasPassed); i++ {
		if !hasPassed[i] {
			testsFailed += 1
			fmt.Println("ERROR: TEST NO.", i, "-", GetFunctionName(tests[i]), "HAS FAILED")
		} else {
			testsSuccesfull += 1
		}
	}
	fmt.Println("FAILED TESTS: ", testsFailed)
	fmt.Println("PASSED TESTS: ", testsSuccesfull)
	fmt.Println("TOTAL TESTS: ", len(tests))

}

func testECONPrivatize() bool {
	var cnts = parseCountriesFromJSON(readFromFile(countriesFilename))
	var pickedNation = cnts[0]
	var bakGDP = pickedNation.Gdp
	pickedNation = ECONPrivatize(pickedNation)
	return bakGDP == pickedNation.Gdp/4
}

func testsECONTick() bool {

	var butter Good = parseJSONToGood(readFromFile(goodsFilename))
	for {
		butter = goodCalculateNextTick(butter)
		if butter.Price > 1.5 {
			return true
		} else {
			return false
		}
	}
}

func testECONNationalize() bool {
	var cnt = parseCountriesFromJSON(readFromFile(countriesFilename))[0]
	var pastMoney = cnt.Money
	cnt = ECONNationalizeIndustries(cnt)
	if cnt.Money/40 == pastMoney {
		return true
	}
	return false
}
func testECONDebase() bool {
	var cnt = loadCountryFromString(readFromFile(debugCountryFilename))
	var bakgdp = cnt.Gdp
	cnt = ECONDebaseCurrency(cnt)
	if bakgdp == cnt.Gdp*2 {
		return true
	}
	return false
}

func testECONLowerTaxes() bool {
	var cnt = loadCountryFromString(readFromFile(debugCountryFilename))
	var bakgdp = cnt.Gdp
	cnt = ECONLowerTax(cnt)
	if cnt.Gdp-5 == bakgdp {
		return true
	}
	return false
}

func testECONRaiseTaxes() bool {
	var cnt = loadCountryFromString(readFromFile(debugCountryFilename))
	var bakgdp = cnt.Gdp
	cnt = ECONRaiseTax(cnt)
	if cnt.Gdp == bakgdp-5 {
		return true
	}
	return false
}
