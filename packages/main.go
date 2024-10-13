package main

import (
	"fmt"
	"packages/mypackage"
)

func main() {
	fmt.Println("Hello, World!")
	mypackage.ExportedFunction()
	fmt.Println(mypackage.ExportedVar)
}
