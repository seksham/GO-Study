package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func appleReader(ctx context.Context, parentWg *sync.WaitGroup, appleStream chan any) {
	defer parentWg.Done()
	var wg sync.WaitGroup
	newCtx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	doWork := func(id int){
		defer wg.Done()
		for{
			select{
			case <- newCtx.Done():
				return
			case v, ok := <- appleStream:
				if(!ok){
					fmt.Println("Apple stream closed")
					return
				}
				time.Sleep(time.Second)
				fmt.Println(v, id)
			}
		}
	}

	for i:= 0; i<3; i++{
		wg.Add(1)
		go doWork(i)
	}
	wg.Wait()

}

func genericReader(ctx context.Context, wg *sync.WaitGroup, stream chan any){
	defer wg.Done()
	for{
		select{
		case <- ctx.Done():
			return
		case v, ok := <- stream:
			if(!ok){
				fmt.Println("Stream closed")
				return
			}
			time.Sleep(time.Second)
			fmt.Println(v)
		}
	}

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	appleStream := make(chan any)
	orangeStream := make(chan any)
	peachStream := make(chan any)
	generator := func(data string, c chan any) {
		for {
			select {
			case <- ctx.Done():
				return
			case c <- data:
			}
		}
	}

	go generator("apple", appleStream)
	go generator("orange", orangeStream)
	go generator("peach", peachStream)
	var wg sync.WaitGroup 
	func ()  {
		wg.Add(1)
		go appleReader(ctx, &wg, appleStream)
		wg.Add(1)
		go genericReader(ctx, &wg, orangeStream)
		wg.Add(1)
		go genericReader(ctx, &wg, peachStream)
	}()
	wg.Wait()
}
