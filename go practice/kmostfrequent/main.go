package main

import (
	"fmt"
	"sort"
)

/*
Given an integer array nums and an integer k, return the k most frequent elements. Example 1:
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
*/

type kv struct {
	key int
	val int
}

func main() {
	freqMap := make(map[int]int)
	k := 2
	s := []int{1, 1, 1, 2, 2, 3, 9, 9, 9, 9, 9, 6, 6, 6, 6, 6, 6, 6, 6}
	for _, val := range s {
		freqMap[val]++
	}
	kvSlice := []kv{}
	for key, val := range freqMap {
		kvSlice = append(kvSlice, kv{key, val})
	}
	sort.Slice(kvSlice, func(i, j int) bool {
		return kvSlice[i].val > kvSlice[j].val
	})
	fmt.Println(kvSlice[:k])
}
