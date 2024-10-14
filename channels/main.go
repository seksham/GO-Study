package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkLink(url string, c chan string){
	_, err := http.Get(url)
	if err != nil{
		fmt.Println(url, "is down")
		c <- url
		return
	}
	fmt.Println(url, "is Up!")
	c <- url
}

func main(){
	links := [] string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	c := make(chan string)
	for _, url := range(links){
		go checkLink(url, c)
	}
	for l := range c{
		go func(link string){
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}
// package main

// import "time"

// 	func main() {
// 		c := make(chan string)
// 		go func() {
// 			c <- "Hi there!"
// 		}()
// 		select {}
// 	}
// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     go func() {
//         for {
//             fmt.Println("Goroutine is running...")
//             time.Sleep(time.Second)
//         }
//     }()

//     // This will cause the program to hang indefinitely
//     select{}
// }

// package main

// func main() {
//     // This will cause the program to hang indefinitely
//     select{}
// }

// func main() {
// 	c := make(chan string)

// 	go func() {
// 		c <- "Hi there!"
// 	}()

// 	message := <-c
// 	fmt.Println("Received:", message)
// }
