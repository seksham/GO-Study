package main

import (
	"fmt"
	"sync"
)

func printEvenNumber(wg *sync.WaitGroup, mutex *sync.Mutex, turn *int) {
	defer wg.Done()
	for i := 0; i <= 10; i += 2 {
		for {
			if *turn == 0 {
				mutex.Lock()
				fmt.Println(i, "is even")
				*turn = 1
				mutex.Unlock()
				break
			}
		}
	}
}

func printOddNumber(wg *sync.WaitGroup, mutex *sync.Mutex, turn *int) {
	defer wg.Done()
	for i := 1; i <= 9; i += 2 {
		for {
			if *turn == 1 {
				mutex.Lock()
				fmt.Println(i, "is odd")
				*turn = 0
				mutex.Unlock()
				break
			}
		}
	}
}

func main() {
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	turn := 0
	wg.Add(2)
	go printEvenNumber(&wg, &mutex, &turn)
	go printOddNumber(&wg, &mutex, &turn)
	wg.Wait()
}
