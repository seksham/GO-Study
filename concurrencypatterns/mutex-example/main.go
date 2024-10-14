package main

import (
	"fmt"
	"sync"
	"time"
)

// mu is a global mutex used to synchronize access to shared resources
var mu sync.Mutex

func process(num int) int {
	time.Sleep(time.Second * 2)
	return num * 2
}

func processDataWithoutConfinement(wg *sync.WaitGroup, num int, result *[]int) {
	defer wg.Done()

	// Lock the mutex before accessing the shared resource (result slice)
	mu.Lock()
	// Unlock the mutex when the function returns
	defer mu.Unlock()

	// Critical section: Append to the shared result slice
	*result = append(*result, process(num))

	// Comments explaining why mutex is required:
	// 1. Multiple goroutines are concurrently accessing and modifying the shared 'result' slice
	// 2. Without synchronization, this can lead to race conditions and data corruption
	// 3. The mutex ensures that only one goroutine can append to the slice at a time
	// 4. This prevents potential issues like:
	//    - Lost updates (one goroutine's append being overwritten by another)
	//    - Incorrect slice length
	//    - Memory corruption due to simultaneous memory allocation and copying
}

func main() {
	start := time.Now()
	a := []int{1, 2, 3, 4, 5}
	result := []int{}
	var wg sync.WaitGroup
	for _, num := range a {
		wg.Add(1)
		go processDataWithoutConfinement(&wg, num, &result)
	}
	wg.Wait()
	fmt.Println(result)
	fmt.Println(time.Since(start))
}
