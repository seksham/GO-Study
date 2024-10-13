package main

import "fmt"

func printEven(e, o chan bool) {
	for i := 0; i <= 10; i += 2 {
		<-e
		fmt.Println(i, "is even")
		o <- true
	}
}

func printOdd(e, o, done chan bool) {
	for i := 1; i <= 9; i += 2 {
		<-o
		fmt.Println(i, "is odd")
		e <- true
	}
	<-o
	done <- true
}

func main() {
	oddChannel := make(chan bool)
	evenChannel := make(chan bool)
	done := make(chan bool)
	go printEven(evenChannel, oddChannel)
	go printOdd(evenChannel, oddChannel, done)
	evenChannel <- true
	<-done
}
