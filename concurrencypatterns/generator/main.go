package main

import (
	"fmt"
	"math/rand"
	"time"
)

// repeaterFunc is a generic function that repeatedly calls a given function
// and sends its results to a channel until a "done" signal is received.
func repeaterFunc[T any, K any](done <-chan T, fn func() K) <-chan K {
	stream := make(chan K)
	go func() {
		defer close(stream)
		for {
			select {
			case <- done:
				return // Exit the goroutine when done signal is received
			default:
				stream <- fn() // Call the provided function and send result to stream
			}
		}
	}()
	return stream
}

func main() {
	done := make(chan bool)
	
	// Define a function that generates random numbers between 0 and 9
	rondomNumber := func() int {
		return rand.Intn(10)
	}
	
	count := 0
	// Use repeaterFunc to generate a stream of random numbers
	for num := range repeaterFunc(done, rondomNumber) {
		time.Sleep(time.Second/4) // Pause for 250ms between numbers
		fmt.Println(num)
		count++
		if count == 20 {
			close(done) // Stop the stream after 20 numbers
		}
	}
}
