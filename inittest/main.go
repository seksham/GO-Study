package main

import "fmt"

var a string

func init() {
	fmt.Println("init", a)
	if a == "" {
		a = "trs"
	}
	fmt.Println("init if", a)
}

func main() {
	fmt.Println("init in main", a)
	a = "test"
	fmt.Println("init in main", a)
}
