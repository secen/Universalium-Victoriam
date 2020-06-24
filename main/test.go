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
	var cnt = country{gdp: 100,
		debt:          20,
		income:        40,
		expenses:      50,
		money:         120,
		population:    300,
		militaryScore: 10,
		cultureScore:  100,
		interest:      0,
		code:          0}
	var str = captureOutput(func() { VIEWShowCountryFinances(cnt) })
	if strings.Contains(str, "Debt:  20") == true {
		return true
	} else {
		return false
	}
}
func testIOReadFromFile() bool {
	var str = readFromFile("nationPickerData.txt")
	if strings.Contains(str, "KOREA - The Hermit Kingdom") == true {
		return true
	} else {
		return false
	}
}

//noinspection ALL
func testECONCountryFinancialModifyTax() bool {
	var cnt = country{gdp: 100,
		debt:          20,
		income:        40,
		expenses:      50,
		money:         120,
		population:    300,
		militaryScore: 10,
		cultureScore:  100,
		interest:      0,
		code:          0}
	var economicsqueue = queue{}
	economicsqueue = appendQueue(economicsqueue, func() { cnt.income += 10 })
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
func testIOLoadCountry() bool {
	var testCountry = loadCountryFromString(readFromFile("debugCountries.txt"))
	if testCountry.code == 12 {
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

	var butter =  Good{2,4,1.5,0.000004,0.002,0.00004};
	for {
		butter=goodCalculateNextTick(butter)
		if butter.price>1.5{
			return true;
		}else {
			return false;
		}
	}
}

func testECONNationalize() bool {
	var cnt = loadCountryFromString(readFromFile(debugCountriesFilename))
	var pastMoney = cnt.money
	cnt = ECONNationalizeIndustries(cnt)
	if cnt.money/30 == pastMoney {
		return true
	}
	return false
}

func testECONDebase() bool {
	var cnt = loadCountryFromString(readFromFile(debugCountriesFilename))
	var bakgdp = cnt.gdp
	cnt = ECONDebaseCurrency(cnt)
	if bakgdp == cnt.gdp*2 {
		return true
	}
	return false
}

func testECONLowerTaxes() bool {
	var cnt = loadCountryFromString(readFromFile(debugCountriesFilename))
	var bakgdp = cnt.gdp
	cnt = ECONLowerTax(cnt)
	if cnt.gdp-5 == bakgdp {
		return true
	}
	return false
}

func testECONRaiseTaxes() bool {
	var cnt = loadCountryFromString(readFromFile(debugCountriesFilename))
	var bakgdp = cnt.gdp
	cnt = ECONRaiseTax(cnt)
	if cnt.gdp == bakgdp-5 {
		return true
	}
	return false
}
