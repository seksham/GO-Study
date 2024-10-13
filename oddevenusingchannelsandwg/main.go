package main

import (
	"fmt"
	"sync"
)

func printEven(e, o chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i += 2 {
		<-e
		fmt.Println(i, "is even")
		o <- true
	}
}

func printOdd(e, o chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 9; i += 2 {
		<-o
		fmt.Println(i, "is odd")
		e <- true
	}
	<-o // Consume the last signal
}

func main() {
	oddChannel := make(chan bool)
	evenChannel := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)

	go printEven(evenChannel, oddChannel, &wg)
	go printOdd(evenChannel, oddChannel, &wg)

	evenChannel <- true // Start the process
	wg.Wait()           // Wait for both goroutines to finish

	// These lines can be safely removed
	close(oddChannel)
	close(evenChannel)
}
