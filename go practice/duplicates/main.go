package main

import (
	"errors"
	"fmt"
	"sort"
)

func checkDuplicates(s []int) map[int]int {
	countMap := make(map[int]int)
	for _, val := range s {
		countMap[val]++
	}
	return countMap
}

type kv struct {
	k int
	v int
}

func topKduplicates(s []int, k int) ([]kv, error) {
	if k <= 0 {
		return nil, errors.New("K should be a positive integer")
	}
	if k > len(s) {
		return nil, errors.New("K should be less than or equal to the length of slice")
	}

	counts := checkDuplicates(s)
	topK := make([]kv, 0, len(counts))

	for key, value := range counts {
		if value >= 2 {
			topK = append(topK, kv{key, value})
		}
	}

	sort.Slice(topK, func(i, j int) bool {
		return topK[i].v > topK[j].v
	})

	if k > len(topK) {
		k = len(topK)
	}

	return topK[:k], nil
}

func main() {
	s := []int{1, 1, 2, 3, 4, 0, 5, 1, 5, 4, 1, 5}
	counts := checkDuplicates(s)
	fmt.Println("Counts:", counts)

	result, err := topKduplicates(s, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Top 2 duplicates: %#v\n", result)
	}
}