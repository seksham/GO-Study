package main

import "fmt"

func FindLastOccurrence(s []int, target int) int {
	low := 0
	high := len(s) - 1
	ans := -1
	var mid int
	for low <= high {
		mid = (low + high) / 2
		if s[mid] == target {
			low = mid + 1
			ans = mid
		} else if target < s[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
}
func main() {
	s := []int{1, 2, 2, 4}
	index := FindLastOccurrence(s, 4)
	fmt.Print(index)
}
