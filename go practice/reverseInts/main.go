package main

import (
	"fmt"
)

func reverseInt(n int) int {
	reversed := 0
	for n != 0 {
		digit := n % 10
		reversed = reversed*10 + digit
		n /= 10
	}
	return reversed
}

func main() {
	fmt.Println(reverseInt(4672))
	fmt.Println(reverseInt(-4672))
	fmt.Println(reverseInt(1000))
	fmt.Println(reverseInt(-1000))
	fmt.Println(reverseInt(1234567890))
	fmt.Println(reverseInt(-1234567890))
	fmt.Println(reverseInt(1))
	fmt.Println(reverseInt(-1))
	fmt.Println(reverseInt(0))
	fmt.Println(reverseInt(10))
	fmt.Println(reverseInt(-10))
}
