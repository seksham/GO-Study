package main

import "fmt"

func main() {
	// Create a slice with initial values
	numbers := []int{1, 2, 3}
	// Append new elements to the slice
	numbers = append(numbers, 4, 5)
	fmt.Println(numbers)

	// Create an array with a fixed size of 3
	var array [3]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	fmt.Println(array)

	// Create a slice using make function with length 3
	slice := make([]int, 3)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	fmt.Println(slice)
	fmt.Println(len(slice)) // Print the length of the slice
	fmt.Println(cap(slice)) // Print the capacity of the slice

	// Append more elements to the slice, causing it to grow
	slice = append(slice, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(slice)
	fmt.Println(len(slice)) // New length after appending
	fmt.Println(cap(slice)) // New capacity after growing

	// Append one more element
	slice = append(slice, 11)
	fmt.Println(slice)
	fmt.Println(len(slice)) // Length increases by 1
	fmt.Println(cap(slice)) // Capacity may have doubled

	// Create a slice with initial length 3 and capacity 10
	slice2 := make([]int, 3, 10)
	fmt.Println(slice2)
	fmt.Println(len(slice2)) // Initial length
	fmt.Println(cap(slice2)) // Initial capacity

	// Append elements to slice2, not exceeding its capacity
	slice2 = append(slice2, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(slice2)
	fmt.Println(len(slice2)) // New length after appending
	fmt.Println(cap(slice2)) // Capacity remains the same
	//demo of copy function
	// Explanation:
	// The copy function copies elements from a source slice to a destination slice.
	// It returns the number of elements copied, which is the minimum of the lengths of the source and destination slices.
	// In this example, slice3 is copied into slice4.
	// The elements from slice3 are copied to the beginning of slice4, and any remaining elements in slice4 are not affected.
	slice3 := []int{1, 2, 3}
	slice4 := []int{4, 5, 6, 7, 8, 9, 10}
	n := copy(slice3, slice4)
	fmt.Println(n) // Output: 3
	fmt.Println(slice3) // Output: [4 5 6] Explanation: The first 3 elements of slice4 are copied to slice3.
	fmt.Println(slice4) // Output: [4 5 6 7 8 9 10] Explanation: The original slice4 remains unchanged.
	slice5 := []int{1, 2, 3}
	slice6 := []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	n = copy(slice6, slice5)
	fmt.Println(n) // Output: 3
	fmt.Println(slice6) // Output: [1 2 3 7 8 9 10 11 12 13 14 15] Explanation: The first 3 elements of slice5 are copied to slice6.
	fmt.Println(slice5) // Output: [1 2 3] Explanation: The original slice5 remains unchanged.
}

