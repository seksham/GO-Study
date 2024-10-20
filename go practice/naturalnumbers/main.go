package main

import (
	"fmt"
)

func main() {
	// var wg sync.WaitGroup
	// wg.Add(1)
	done := make(chan any)
	go func() {
		// defer wg.Done()
		for i := 1; i <= 10; i++ {
			fmt.Println(i)
		}
		done <- true
	}()
	// wg.Wait()
	<- done
}
