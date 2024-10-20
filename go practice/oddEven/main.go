package main

import "fmt"

func printEven(o, e chan bool) {
	for i := 0; i <= 10; i+=2 {
		<-e
		fmt.Println(i)
		o <- true
	}
}

func printOdd(o, e, done chan bool) {
	for i := 1; i <= 9; i+=2 {
		<-o
		fmt.Println(i)
		e <- true
	}
	<-o
	done <- true
}

func main() {
	o:= make(chan bool)
	e:= make(chan bool)
	done := make(chan bool)
	go printEven(o, e)
	go printOdd(o, e, done)
	e <- true
	<- done
}
