package main

import (
	"fmt"
)

func printEven(evenChan, oddChan chan bool) {
	for i := 0; i <= 10; i += 2 {
		<-evenChan           // Wait for signal from main or odd goroutine
		fmt.Println(i, "Even") // Print even number
		oddChan <- true      // Signal odd goroutine
	}
}

func printOdd(oddChan, evenChan chan bool, doneChan chan bool) {
	for i := 1; i <= 9; i += 2 {
		<-oddChan            // Wait for signal from even goroutine
		fmt.Println(i, "Odd")  // Print odd number
		evenChan <- true     // Signal even goroutine
	}
	<-oddChan                // Wait for the last signal from even goroutine
	doneChan <- true         // Signal main function that printing is done
}

func main() {
	evenChan := make(chan bool)
	oddChan := make(chan bool)
	doneChan := make(chan bool)

	go printEven(evenChan, oddChan)
	go printOdd(oddChan, evenChan, doneChan)

	evenChan <- true         // Start the sequence by signaling even goroutine
	<-doneChan               // Wait for the completion signal
}

/*output steps
evenChan kick start the sequence by sending a signal
printEven start execution and print 0
send signal to oddChan to start printing odd numbers
printOdd start execution and print 1
send signal to evenChan to print next even number
printEven start execution and print 2
send signal to oddChan to print next odd number
printOdd start execution and print 3
send signal to evenChan to print next even number
printEven start execution and print 4
send signal to oddChan to print next odd number
printOdd start execution and print 5
send signal to evenChan to print next even number
printEven start execution and print 6
send signal to oddChan to print next odd number
printOdd start execution and print 7
send signal to evenChan to print next even number
printEven start execution and print 8
send signal to oddChan to print next odd number
printOdd start execution and print 9
send signal to evenChan to print next even number
printEven start execution and print 10
send signal to oddChan to print next odd number
Here printOdd will not print 11 because the loop condition is i <= 9
it will still empty the oddChan and evenChan will not get signal to print next even number
main function will get the doneChan signal and print "done"
*/
