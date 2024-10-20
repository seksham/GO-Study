package main

import "fmt"

// abcabcdbb
// O(n^2) inefficient
func findLongestSubString(s string) string {
	n := len(s)
	max := 1
	maxStart := 0
	end := 0
	for start := 0; start < n; start++ {
		freqMap := make(map[string]int)
		for end = start; end < n; end++ {
			freqMap[string(s[end])]++
			if freqMap[string(s[end])] == 2 {
				break
			}
		}
		if end-start > max {
			max = end - start
			maxStart = start
		}
	}
	return s[maxStart : maxStart+max]
}
// O(n) efficient
func longestSubstringWithoutRepeats(s string) string {
    charMap := make(map[rune]int)
    start := 0
    maxLength := 0
    maxStart := 0

    for end, char := range s {
        if lastSeen, found := charMap[char]; found && lastSeen >= start {
            start = lastSeen + 1
        }
        charMap[char] = end
        currentLength := end - start + 1
        if currentLength > maxLength {
            maxLength = currentLength
            maxStart = start
        }
    }

    return s[maxStart : maxStart+maxLength]
}

func main() {
	fmt.Println(findLongestSubString("abcdbcdefdghpqga"))
	fmt.Println(longestSubstringWithoutRepeats("abcdbcdefdghpqga"))
}
