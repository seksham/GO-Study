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

func processData(wg *sync.WaitGroup, num int, result *[]int) {
	defer wg.Done()
	// mu.Lock()
	*result = append(*result, process(num))
	// mu.Unlock()
}

func main() {
	start := time.Now()
	a := []int{1, 2, 3, 4, 5}
	result := []int{}
	var wg sync.WaitGroup
	for _, num := range a {
		wg.Add(1)
		go processData(&wg, num, &result)
	}
	wg.Wait()
	fmt.Println(result)
	fmt.Println(time.Since(start))
}
