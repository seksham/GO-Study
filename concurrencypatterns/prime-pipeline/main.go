package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func primeStreamGenerator[T any](done <-chan T, stream <-chan int) <-chan int {
	primes := make(chan int)
	isPrime := func(n int) bool {
		for i := n - 1; i > 1; i-- {
			if n%i == 0 {
				return false
			}
		}
		return true
	}
	go func() {
		defer close(primes)
		for {
			select {
			case <-done:
				return
			case num := <-stream:
				if isPrime(num) {
					primes <- num
				}
			}
		}
	}()
	return primes
}

func main() {
	t := time.Now()
	done := make(chan bool)
	defer close(done)
	randomIntStream := streamGenerator(done, randIntFetcher)
	primeIntStream := primeStreamGenerator(done, randomIntStream)
	for num := range take(done, primeIntStream, 10) {
		fmt.Println(num)
	}
	fmt.Println(time.Since(t))
}
