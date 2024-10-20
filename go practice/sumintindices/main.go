package main

import "fmt"

func indecesThatMatchtheSum(s []int, target int) (int, int) {
	numMap := make(map[int]int)

	for i, num := range s {
		complement := target - num
		if j, found := numMap[complement]; found {
			return j, i
		}
		numMap[num] = i
	}

	return -1, -1
}

func main() {
	s := []int{9, 2, 3, 3, 6, 5, 4, 1}
	target := 4
	index1, index2 := indecesThatMatchtheSum(s, target)

	if index1 != -1 && index2 != -1 {
		fmt.Printf("Indices that sum to %d: %d and %d\n", target, index1, index2)
		fmt.Printf("Values: %d + %d = %d\n", s[index1], s[index2], target)
	} else {
		fmt.Println("No solution found")
	}
}
