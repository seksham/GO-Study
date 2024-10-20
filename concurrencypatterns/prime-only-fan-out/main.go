package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func fanIn[T any, K any](done <- chan T, channels ...<-chan K) <- chan K{
	fannedInchannel := make(chan K)
	var wg sync.WaitGroup
	transfer := func(c <- chan K){
		defer wg.Done()
		for i := range c{
			select{
			case <- done:
				return
			case fannedInchannel <- i:
			}
		}
	}
	for _,channel := range channels{
		wg.Add(1)
		go transfer(channel)
	}
	go func(){
		wg.Wait()
		close(fannedInchannel)
	}()
	return fannedInchannel
}

func randIntFetcher() int {
	return rand.Intn(50000000)
}

func streamGenerator[T any, K any](done <-chan T, fn func() K) <-chan K {
	stream := make(chan K)
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()
	return stream
}

func take[T any, K any](done <-chan T, stream <-chan K, n int) <-chan K {
	taken := make(chan K)
	go func() {
		defer close(taken)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case taken <- <-stream:
			}
		}
	}()
	return taken
}

func primeStreamGenerator[T any](done <-chan T, stream <-chan int, finalChannel chan int){
	isPrime := func(n int) bool {
		for i := n - 1; i > 1; i-- {
			if n%i == 0 {
				return false
			}
		}
		return true
	}
	go func() {
		defer close(finalChannel)
		for {
			select {
			case <-done:
				return
			case num := <-stream:
				if isPrime(num) {
					finalChannel <- num
				}
			}
		}
	}()
}

func main() {
	t := time.Now()
	done := make(chan bool)
	defer close(done)

	// Fan-out: Create a single source of random integers
	randomIntStream := streamGenerator(done, randIntFetcher)

	availableCPUs := runtime.NumCPU()
	fmt.Println("available cpus are:", availableCPUs)

	finalChannel := make(chan int)

	// Fan-out: Distribute work across multiple goroutines
	for i := 0; i < availableCPUs; i++ {
		primeStreamGenerator(done, randomIntStream, finalChannel)
	}

	// Process the results from the combined channel
	for num := range take(done, finalChannel, 10) {
		fmt.Println(num)
	}

	fmt.Println(time.Since(t))
}
