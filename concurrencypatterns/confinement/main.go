package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

func process(num int) int {
	time.Sleep(time.Second * 2)
	return num * 2
}

func processData(wg *sync.WaitGroup, num int, result *int) {
	defer wg.Done()
	*result =  process(num)
}

func main() {
	start := time.Now()
	a := []int{1, 2, 3, 4, 5}
	result := make([]int, len(a))
	var wg sync.WaitGroup
	for i, num := range a {
		wg.Add(1)
		go processData(&wg, num, &result[i]) //no shared resource between goroutines because of confinement. Confinement is a principle that restricts the scope of a variable to a specific part of the program.
	}
	wg.Wait()
	fmt.Println(result)
	fmt.Println(time.Since(start))
}
