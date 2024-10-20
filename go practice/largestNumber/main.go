package main

import "fmt"

// [9 4 5 9 1 3]
func findTwoLargest(nums []int) (int, int) {
	if len(nums) < 2 {
		return 0, 0 // or handle error appropriately
	}

	largest, secondLargest := nums[0], nums[1]
	if secondLargest > largest {
		largest, secondLargest = secondLargest, largest
	}

	for _, num := range nums {
		if num > largest {
			secondLargest = largest
			largest = num
		} else if num > secondLargest && num < largest {
			secondLargest = num
		}
	}

	return largest, secondLargest
}

func main() {
	// s := []int{3, 1, 5, 55, 33, 100, 46, 32, 2, 2, 333, 99}
	s := []int{9, 4, 5, 9, 1, 3}
	l, sl := findTwoLargest(s)
	fmt.Println(l, sl)
}
