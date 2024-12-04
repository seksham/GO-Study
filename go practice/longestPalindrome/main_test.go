package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

// generateTestString creates a string of given size with random lowercase letters
func generateTestString(size int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, size)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// generatePalindromicString creates a string of given size that contains some palindromes
func generatePalindromicString(size int) string {
	if size <= 2 {
		return generateTestString(size)
	}

	// Create base random string
	s := []byte(generateTestString(size))

	// Insert some palindromes of various sizes
	palindromes := []string{
		"aba",
		"abba",
		"abcba",
		"abccba",
		"racecar",
	}

	// Insert palindromes at random positions
	for _, p := range palindromes {
		if len(p) < size {
			pos := rand.Intn(size - len(p))
			copy(s[pos:], p)
		}
	}

	return string(s)
}

// generateWorstCaseString creates a string that's particularly bad for expand-around-center
// It creates a string like "aaaa..." which forces maximum expansions at each position
func generateWorstCaseString(size int) string {
	return strings.Repeat("a", size)
}

// generateOverlappingPalindromes creates a string with many overlapping palindromes
// It creates patterns like "aaabaaaabaaaabaaa..."
func generateOverlappingPalindromes(size int) string {
	if size <= 4 {
		return strings.Repeat("a", size)
	}

	// Create pattern that forces lots of overlapping checks
	pattern := "aaab"
	result := strings.Builder{}
	result.Grow(size)

	for result.Len() < size {
		result.WriteString(pattern)
	}

	return result.String()[:size]
}

func BenchmarkPalindrome(b *testing.B) {
	sizes := []int{10, 100, 500, 1000, 2000, 5000, 40000}

	for _, size := range sizes {
		// Test with random strings
		randomStr := generateTestString(size)
		b.Run(fmt.Sprintf("Random_Expand_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				findLongestPalindromeExpand(randomStr)
			}
		})
		b.Run(fmt.Sprintf("Random_Manacher_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				findLongestPalindromeOptimisedManacher(randomStr)
			}
		})

		// Test with strings containing palindromes
		palindromicStr := generatePalindromicString(size)
		b.Run(fmt.Sprintf("Palindromic_Expand_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				findLongestPalindromeExpand(palindromicStr)
			}
		})
		b.Run(fmt.Sprintf("Palindromic_Manacher_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				findLongestPalindromeOptimisedManacher(palindromicStr)
			}
		})
	}
}

func BenchmarkWorstCase(b *testing.B) {
	sizes := []int{1000, 5000, 10000, 50000, 100000}

	for _, size := range sizes {
		// Test with worst-case string (all same characters)
		worstCase := generateWorstCaseString(size)
		b.Run(fmt.Sprintf("WorstCase_Expand_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				findLongestPalindromeExpand(worstCase)
			}
		})
		b.Run(fmt.Sprintf("WorstCase_Manacher_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				findLongestPalindromeOptimisedManacher(worstCase)
			}
		})

		// Test with overlapping palindromes
		overlapping := generateOverlappingPalindromes(size)
		b.Run(fmt.Sprintf("Overlapping_Expand_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				findLongestPalindromeExpand(overlapping)
			}
		})
		b.Run(fmt.Sprintf("Overlapping_Manacher_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				findLongestPalindromeOptimisedManacher(overlapping)
			}
		})
	}
}

// TestCorrectnessWorstCase verifies both algorithms produce correct results
func TestCorrectnessWorstCase(t *testing.T) {
	testCases := []struct {
		name string
		size int
	}{
		{"Small_WorstCase", 10},
		{"Medium_WorstCase", 100},
		{"Large_WorstCase", 1000},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := generateWorstCaseString(tc.size)
			expand := findLongestPalindromeExpand(input)
			manacher := findLongestPalindromeOptimisedManacher(input)

			if len(expand) != len(manacher) {
				t.Errorf("Length mismatch: Expand=%d, Manacher=%d", len(expand), len(manacher))
			}
			if expand != manacher {
				t.Errorf("Result mismatch:\nExpand: %s\nManacher: %s", expand, manacher)
			}
			if len(expand) != tc.size {
				t.Errorf("Expected palindrome of length %d, got %d", tc.size, len(expand))
			}
		})
	}
}
