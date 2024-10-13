package mypackage

import "fmt"

func ExportedFunction() {
	fmt.Println("Hello from myPackage")
	unexportedFunction()
}

func unexportedFunction() {
	fmt.Println("Log inside unexported function called from exported function")
	fmt.Println(unexportedVar)
}

var ExportedVar = "I'm an exported variable"

var unexportedVar = "I'm an unexported variable usedinside unexported function"

func PrintHello() {
	fmt.Println("Hello from myPackage")
}
