package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func numberGenerator(ctx context.Context) chan any{
	stream := make(chan any)
	go func(){
		defer close(stream)
		for{
			select{
			case <- ctx.Done():
				fmt.Println("context cancelled")
				stream <- ctx.Err()
				return
			case stream <- rand.Intn(100):
			}
		}
	}()
	return stream
}

func main() {
	// ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
	ctx, cancel := context.WithTimeout(context.Background(),time.Second*10)
	defer cancel()
	numberStream := numberGenerator(ctx)
	for num := range numberStream{
		time.Sleep(time.Second/4)
		switch num.(type){
		case int:
			fmt.Println(num)
		case error:
			fmt.Println(num)
			return
		}
	}


}
