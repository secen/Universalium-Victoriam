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
	var cnt Country = parseCountryFromJSON(readFromFile(countriesFilename))
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
func testParsingOfListings() bool{
	var str =readFromFile("nationListings.json");
	var listings []CountryListing = parseJSONToCountryListings(str);
	if strings.Compare(listings[0].CountryName, "SCOTLAND") == 0{
		return true;
	}
	return false;
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
	var laws2 = make([]law, 10)
	laws2[0] = law{name: "Testing"}
	laws2[1] = law{name: "Testing2"}
	laws2[2] = law{name: "Mandate Of Heaven"}
	abolishLaw(1, laws2)
	if strings.Compare(laws2[1].name, "Mandate Of Heaven") == 0 {
		return true
	} else {
		return false
	}
}
func testTechConsoleUI() bool {
	var cnt Country = parseCountryFromJSON(readFromFile("countries.json"))
	var techArr []Technology = fromJSONToTechArr(readFromFile("techs.json"))
	cnt.techs = append(cnt.techs, techArr[0])
	var output = captureOutput(func(){VIEWWRITECountryTechs(cnt)});
	if strings.Contains(output, "[TAKEN]||Central Banking||Banking And Finance"){
		return true;
	}
	return false
}
func testIOLoadCountry() bool {
	var testCountry Country = loadCountryFromString(readFromFile(debugCountryFilename))
	if testCountry.Code == 12 {
		return true
	} else {
		return false
	}
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
		testsECONTick,
		testParsingOfListings,
		testTechConsoleUI,
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

func testsECONTick() bool {

	var butter Good = parseJSONToGood(readFromFile("goods.json"))
	for {
		butter=goodCalculateNextTick(butter)
		if butter.Price>1.5{
			return true;
		}else {
			return false;
		}
	}
}

func testECONNationalize() bool {
	var cnt = loadCountryFromString(readFromFile(debugCountryFilename))
	var pastMoney = cnt.Money
	cnt = ECONNationalizeIndustries(cnt)
	if cnt.Money/30 == pastMoney {
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
