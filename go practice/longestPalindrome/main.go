package main

import (
	"fmt"
	"strings"
)

// isPalindrome checks if a given string is a palindrome
func isPalindrome(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}

// printMatrix prints a 2D integer matrix
func printMatrix(matrix [][]bool) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

// Helper function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// manacher's algorithm
func transformForManacher(s string) string {
	transformedString := "#"
	for _, c := range s {
		transformedString += string(c) + "#"
	}
	return transformedString
}

// findLongestPalindromeManacher implements Manacher's algorithm
// Visual representation:
//
//	Original string: a b a c b c d c b c a c a b a
//	Transformed:     #   a   #   b   #   a   #   c   #   b   #   c   #   d   #   c   #   b   #   c   #   a   #   c   #   a   #   b   #   a   #
//	Index:           0   1   2   3   4   5   6   7   8   9   10  11  12  13  14  15  16  17  18  19  20  21  22  23  24  25  26  27  28  29  30
//	P array:         0   1   0   3   0   1   0   1   0   3   0   1   0   9   0   1   0   3   0   1   0   _   _   _   _   _   _   _   _   _   _
//										 |		 |_______|_______|		 C		 |_______|_______|		 |
//										 |_______________j_______________________________i_______________|
//										 L															     R
//
//	L: leftBoundary (index 5, character 'a')
//	R: rightBoundary (index 21, character 'a')
//	i: current position (index 17, character 'b')
//	j: mirror of i (index 9, character 'b')
//
// Time complexity: O(n)
// Space complexity: O(n)
func findLongestPalindromeManacher(s string) string {
	// Step 1: Transform the input string
	t := transformForManacher(s)
	// fmt.Println("Transformed string:", t)

	// Initialize variables
	leftBoundary := 0   // L: left boundary of rightmost palindrome
	rightBoundary := 0  // R: right boundary of rightmost palindrome
	n := len(t)         // length of transformed string
	p := make([]int, n) // P array to store palindrome radii
	maxLen := 0         // length of the longest palindrome found
	center := 0         // center of the longest palindrome
	p[0] = 0            // first element is always 0

	// Step 2: Iterate through the transformed string
	for i := 1; i < n; i++ {
		var k int // k represents the current radius of palindrome centered at i

		// Step 3: Determine the initial value of k
		if i > rightBoundary {
			// Case 1: i is outside the rightmost known palindrome
			k = 0 // start with a radius of 0
		} else {
			// Case 2: i is inside the rightmost known palindrome
			j := leftBoundary + (rightBoundary - i) // j is the mirror of i
			if j-p[j] > leftBoundary {
				// Subcase 2a: The mirrored palindrome is entirely within the current palindrome
				p[i] = p[j] // copy the length from the mirrored position
				continue    // no need to expand, move to next i
			} else {
				// Subcase 2b: The mirrored palindrome reaches the boundary
				k = rightBoundary - i // start with the distance to the rightBoundary
			}
		}

		// Step 4: Attempt to expand palindrome centered at i
		for i-k >= 0 && i+k < n && t[i-k] == t[i+k] {
			k++
		}
		k--      // adjust k as the loop overshoots by 1
		p[i] = k // store the palindrome radius at position i

		// Step 5: Update longest palindrome if necessary
		if p[i] > maxLen {
			maxLen = p[i]
			center = i / 2 // convert to index in original string
		}

		// Step 6: Update boundaries of right most palindrome if palindrome at i expands past rightBoundary.
		if i+k > rightBoundary {
			leftBoundary = i - k
			rightBoundary = i + k
		}
	}

	// Step 7: Extract the longest palindrome from the original string
	start := center - maxLen/2
	end := start + maxLen
	return s[start:end]
}



func findLongestPalindromeOptimisedManacher(s string) string {
	// Transform string by adding boundaries and separators
	// Example: "abc" -> "^#a#b#c#$"
	// ^ and $ serve as sentinels, eliminating bound checking
	// equivalent to transformForManacher() but with boundaries
	T := "^#" + strings.Join(strings.Split(s, ""), "#") + "#$"
	n := len(T)

	// P[i] stores the radius of palindrome centered at i
	// equivalent to p[] array in original version
	P := make([]int, n)

	// C: Center of the current palindrome being processed
	// R: Right boundary of the current palindrome
	// These replace leftBoundary and rightBoundary from original
	// Note: leftBoundary can be calculated as 2*C-R if needed
	C, R := 0, 0

	// Main loop: process each character except boundaries
	// n-1 because we don't need to process the $ boundary
	for i := 1; i < n-1; i++ {
		// If current position is within a known palindrome
		// equivalent to the else block in original version
		if R > i {
			// Mirror position is 2*C-i (equivalent to j in original)
			// min(R-i, P[2*C-i]) handles two cases from original:
			// 1. If mirror palindrome fits within boundary: use P[2*C-i]
			// 2. If it extends beyond boundary: use R-i
			P[i] = min(R-i, P[2*C-i])
		}
		// else implicitly P[i] = 0 (equivalent to k = 0 in original)

		// Attempt to expand palindrome centered at i
		// T[i+1+P[i]] is the next character to right of known palindrome
		// T[i-1-P[i]] is the next character to left of known palindrome
		// Equivalent to the expansion loop in original but without bound checking
		// because ^ and $ boundaries will never match
		for T[i+1+P[i]] == T[i-1-P[i]] {
			P[i]++
		}

		// If palindrome centered at i expands past R,
		// update C and R to this new, larger palindrome
		// equivalent to boundary update in original
		if i+P[i] > R {
			C = i
			R = i + P[i]
		}
	}

	// Find the longest palindrome
	// equivalent to tracking maxLen and center in original
	maxLen := 0
	centerIndex := 0
	for i, v := range P {
		if v > maxLen {
			maxLen = v
			centerIndex = i
		}
	}

	// Convert center and length from transformed string to original string
	// (centerIndex-maxLen)/2 gives start position in original string
	// The division by 2 is needed because we added extra characters
	return s[(centerIndex-maxLen)/2 : (centerIndex+maxLen)/2]
}

// findLongestPalindromeExpand uses center expansion technique
// Visual representation:
//
//	s: a b c b a
//	   ^ ^ ^ ^ ^
//	   | | | | |
//	   | | | | expand around 'a'
//	   | | | expand around 'b'
//	   | | expand around 'c'
//	   | expand around 'b'
//	   expand around 'a'
//
// Time complexity: O(n^2)
// Space complexity: O(1)
func findLongestPalindromeExpand(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	expandAroundCenter := func(left, right int) int {
		for left >= 0 && right < n && s[left] == s[right] {
			left--
			right++
		}
		return right - left - 1 // Length of palindrome: (right-1) - (left+1) + 1
	}

	// Changed initial maxLen to 1 since any single character is a palindrome
	start, maxLen := 0, 1
	// Changed loop condition to optimize by skipping positions that can't yield longer palindromes
	for i := 0; i < n-maxLen/2; i++ {
		// Check for odd-length palindromes
		len1 := expandAroundCenter(i, i)
		// Check for even-length palindromes
		len2 := expandAroundCenter(i, i+1)
		lenMax := max(len1, len2)
		if lenMax > maxLen {
			start = i - (lenMax-1)/2
			maxLen = lenMax
			// fmt.Printf("Palindrome found at index %d, length %d\n", i, lenMax)
		}
	}
	// fmt.Printf("Longest palindrome length: %d\n", maxLen)
	return s[start : start+maxLen]
}

// findLongestPalindromeDP uses dynamic programming method
// Visual representation of the DP matrix for "abcba":
//
//	  a b c b a
//	a 1 0 0 0 1
//	b   1 0 1 0
//	c     1 0 0
//	b       1 0
//	a         1
//
// 1 indicates a palindrome
// Time complexity: O(n^2)
// Space complexity: O(n^2)
func findLongestPalindromeDP(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	start, maxLength := 0, 1

	// All substrings of length 1 are palindromes
	for beg := 0; beg < n; beg++ {
		dp[beg][beg] = true
	}

	// Check for substrings of length 2
	for beg := 0; beg < n-1; beg++ {
		end := beg + 1
		if s[beg] == s[end] {
			dp[beg][end] = true
			start = beg
			maxLength = 2
		}
	}

	// Check for substrings of length 3 or more
	for length := 3; length <= n; length++ {
		for beg := 0; beg < n-length+1; beg++ {
			end := beg + length - 1
			if s[beg] == s[end] && dp[beg+1][end-1] {
				dp[beg][end] = true
				if length > maxLength {
					start = beg
					maxLength = length
				}
			}
		}
	}
	// printMatrix(dp)

	return s[start : start+maxLength]
}

// findLongestPalindromeBruteForce uses brute force approach
// It checks all possible substrings
// Time complexity: O(n^3)
// Space complexity: O(1)
func findLongestPalindromeBruteForce(s string) string {
	maxLen := 0
	longestSubstring := string(s[0])
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			substring := s[i:j]
			if isPalindrome(substring) && len(substring) > maxLen {
				maxLen = len(substring)
				longestSubstring = substring
			}
		}
	}
	return longestSubstring
}

func main() {
	sub := findLongestPalindromeBruteForce("babad")
	fmt.Println(sub)
	sub = findLongestPalindromeBruteForce("abccadddcdddacqlde")
	fmt.Println(sub)
	fmt.Println(isPalindrome("abccadddcdddacqlde"))
	fmt.Println("DP:", findLongestPalindromeDP("abccadddcdddacqlde"))
	fmt.Println("Expand:", findLongestPalindromeExpand("abccadddcdddacqlde"))
	fmt.Println("Expand:", findLongestPalindromeExpand("abccadddccdddacqlde"))
	fmt.Println("Manacher:", findLongestPalindromeManacher("abccadddcdddacqlde"))
	fmt.Println("Manacher:", findLongestPalindromeManacher("abccadddccdddacqlde"))
}
