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
func TestShowCountryFinancials() {
	var cnt = country{gdp: 100,
		debt:20,
		income:40,
		expenses:50,
		money:120,
		population:300,
		militaryScore:10,
		cultureScore:100,
		interest: 0,
		code:0}
	var str = captureOutput(func(){writeFinances(cnt)})
	if strings.Contains(str,"Debt:  20") == true {
		fmt.Println("Test Passed!")
	}else {
		fmt.Println("Test Failed")
	}
}
func testAbolishLaw() {
	var laws2 = make([]law,10);
	laws2[0] = law{name:"Testing"}
	laws2[1] = law{name:"Testing2"}
	laws2[2] = law{name:"Mandate Of Heaven"}
	abolishLaw(1,laws2)
	if strings.Compare(laws2[1].name,"Mandate Of Heaven") == 0 {
		fmt.Println("Test Passed!")
	} else{
		fmt.Println("Test failed")
	}
}
func tests() {
	fmt.Println("Testing country financial showing info:")
	TestShowCountryFinancials()
	fmt.Println("Testing abolishing laws: ");
	testAbolishLaw();
	fmt.Println("Testing Complete")
}