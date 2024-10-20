package main

import (
	"testing"
)

func TestFindTwoLargest(t *testing.T) {
	tests := []struct {
		name           string
		input          []int
		expectLargest  int
		expectSecond   int
		expectError    bool
	}{
		{
			name:          "Normal case",
			input:         []int{3, 7, 2, 8, 5, 1, 9, 4},
			expectLargest: 9,
			expectSecond:  8,
		},
		{
			name:          "All same numbers",
			input:         []int{5, 5, 5, 5, 5},
			expectLargest: 5,
			expectSecond:  5,
		},
		{
			name:          "Two elements",
			input:         []int{3, 1},
			expectLargest: 3,
			expectSecond:  1,
		},
		{
			name:          "Negative numbers",
			input:         []int{-5, -2, -10, -1, -8},
			expectLargest: -1,
			expectSecond:  -2,
		},
		{
			name:          "Largest number repeated",
			input:         []int{9, 4, 5, 9, 1, 3},
			expectLargest: 9,
			expectSecond:  5,
		},
		{
			name:        "One element",
			input:       []int{1},
			expectError: true,
		},
		{
			name:        "Empty slice",
			input:       []int{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			largest, second := findTwoLargest(tt.input)
			
			if tt.expectError {
				if largest != 0 || second != 0 {
					t.Errorf("Expected error for input %v, got %d and %d", tt.input, largest, second)
				}
			} else {
				if largest != tt.expectLargest {
					t.Errorf("Largest number for input %v: expected %d, got %d", tt.input, tt.expectLargest, largest)
				}
				if second != tt.expectSecond {
					t.Errorf("Second largest number for input %v: expected %d, got %d", tt.input, tt.expectSecond, second)
				}
			}
		})
	}
}