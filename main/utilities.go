package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
	"runtime"
)

var clear map[string]func() //create a map for storing clear funcs
func readFromFile(fileName string)string{
	var s,_=ioutil.ReadFile(fileName)
	return string(s)
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
}
func containsTech(arr []technology, t technology)bool{
	for i:=0;i<len(arr); i++{
		if arr[i].ID ==t.ID {
			return true
		}
	}
	return false
}
func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok { //if we defined a clear func for that platform:
		value()  //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
