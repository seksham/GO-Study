package main

import "fmt"

func numberCounter(start int) func() int {
	count := start

	return func() int {
		count++
		return count
	}
}

// more examples
func debitCardFunction(balance int) func(int) int {
	return func(withdrawal int) int {
		if balance < withdrawal {
			return balance
		}
		balance -= withdrawal
		return balance
	}
}

func main() {
	counter1 := numberCounter(2)
	fmt.Println(counter1())
	fmt.Println(counter1())
	counter2 := numberCounter(3)
	fmt.Println(counter2())
	fmt.Println(counter2())
	fmt.Println(counter1())

	debitCard := debitCardFunction(100)
	fmt.Println(debitCard(20))
	fmt.Println(debitCard(100))
	fmt.Println(debitCard(40))
	debitCard2 := debitCardFunction(100)
	fmt.Println(debitCard(20))
	fmt.Println(debitCard2(90))
	fmt.Println(debitCard2(40))
}
