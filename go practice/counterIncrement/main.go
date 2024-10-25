package main

import (
	"fmt"
	"sync"
)

/*
Write a Go program that spawns two goroutines. Each goroutine should increment a shared counter variable,
but only one goroutine should increment it by 1, and the other should increment it by 2.
The two goroutines should alternate in their operations, ensuring the increments are performed in a synchronized manner,
and the final value of the counter should be consistent.

Ex:
1  (increment by 1)
3  (increment by 2)
4  (increment by 1)
6  (increment by 2)
7  (increment by 1)
Final counter value: 7
*/
var counter int
func incrementByOne(wg *sync.WaitGroup, k int, one, two chan bool) {
	defer close(two)
	defer wg.Done()
	for {
		select {
		case _, ok := <-one:
			if !ok {
				return
			}
			if counter >= k {
				return
			}
			counter++
			fmt.Println(counter)
			two <- true
		}
	}
}

func incrementByTwo(wg *sync.WaitGroup, k int, one, two chan bool) {
	defer close(one)
	defer wg.Done()
	for {
		select {
		case _, ok := <-two:
			if !ok {
				return
			}
			if counter >= k {
				return
			}
			counter += 2
			fmt.Println(counter)
			one <- true
		}
	}
}

// 1 3 4 6 7

func main() {
	k := 7
	one := make(chan bool)
	two := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)
	go incrementByOne(&wg, k, one, two)
	go incrementByTwo(&wg, k, one, two)
	one <- true
	wg.Wait()
	// time.Sleep(time.Second)

}


