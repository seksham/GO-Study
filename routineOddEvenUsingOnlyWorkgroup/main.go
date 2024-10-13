package main

import (
	"fmt"
	"sync"
)

func printEven(wg *sync.WaitGroup, mutex *sync.Mutex, turn *int) {
	defer wg.Done()
	for i := 0; i <= 10; i += 2 {
		for {
			mutex.Lock()
			if *turn == 0 {
				fmt.Println(i, "Even")
				*turn = 1
				mutex.Unlock()
				break
			}
			mutex.Unlock()
		}
	}
}

func printOdd(wg *sync.WaitGroup, mutex *sync.Mutex, turn *int) {
	defer wg.Done()
	for i := 1; i <= 9; i += 2 {
		for {
			mutex.Lock()
			if *turn == 1 {
				fmt.Println(i, "Odd")
				*turn = 0
				mutex.Unlock()
				break
			}
			mutex.Unlock()
		}
	}
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	turn := 0

	wg.Add(2)
	go printEven(&wg, &mutex, &turn)
	go printOdd(&wg, &mutex, &turn)

	wg.Wait()
}