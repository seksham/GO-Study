package main

import (
	"fmt"
	"sync"
)

// orDone is a concurrency pattern that combines a 'done' channel with another channel
// It allows for graceful cancellation of channel operations
func orDone(done, c <-chan any) <-chan any {
	relayStream := make(chan any)
	go func() {
		defer close(relayStream)
		for {
			select {
			// If 'done' signal is received, exit immediately
			case <-done:
				return
			// Try to receive from the input channel 'c'
			case v, ok := <-c:
				// If 'c' is closed, exit
				if !ok {
					return
				}
				// Try to send the received value, but also check for 'done' signal
				select {
				case <-done:
					return
				case relayStream <- v:
					// Successfully sent the value
				}
			}
		}
	}()
	return relayStream
}

// consumeCows uses orDone to safely consume values from the 'cows' channel
func consumeCows(done <-chan any, cows <-chan any, wg *sync.WaitGroup) {
	for cow := range orDone(done, cows) {
		// some complex logic
		fmt.Println(cow)
	}
	wg.Done()
}

// consumePigs uses orDone to safely consume values from the 'pigs' channel
func consumePigs(done <-chan any, pigs <-chan any, wg *sync.WaitGroup) {
	for pig := range orDone(done, pigs) {
		// some complex logic
		fmt.Println(pig)
	}
	wg.Done()
}

func main() {
	// Create a WaitGroup to synchronize goroutines
	var wg sync.WaitGroup
	
	// Create channels for cows and pigs
	cows := make(chan any)
	pigs := make(chan any)
	
	// Create a done channel for cancellation
	done := make(chan any)
	
	// Start a goroutine to send "mow" to the cows channel
	go func() {
		for {
			select {
			case <-done:
				// Exit if done signal is received
				return
			default:
				// Send "mow" to the cows channel
				cows <- "mow"
			}
		}
	}()

	// Start a goroutine to send "squik" to the pigs channel
	go func() {
		for {
			select {
			case <-done:
				// Exit if done signal is received
				return
			default:
				// Send "squik" to the pigs channel
				pigs <- "squik"
			}
		}
	}()

	// Add 1 to the WaitGroup and start consuming cows
	wg.Add(1)
	go consumeCows(done, cows, &wg)
	
	// Add 1 to the WaitGroup and start consuming pigs
	wg.Add(1)
	go consumePigs(done, pigs, &wg)
	
	// Wait for all goroutines to complete
	wg.Wait()
}
